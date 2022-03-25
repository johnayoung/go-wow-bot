package nav

import (
	"fmt"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/johnayoung/go-wow-bot/utils"
	"golang.org/x/sys/windows"
)

const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

var (
	mapsPath = filepath.Join(utils.RootDir(), "Bot", "Navigation.dll")
)

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	modversion  = windows.NewLazySystemDLL("version.dll")
	modpsapi    = windows.NewLazySystemDLL("psapi.dll")
	modntdll    = windows.NewLazySystemDLL("ntdll.dll")
	modNav      = windows.NewLazyDLL(mapsPath)

	procGetNativeSystemInfo       = modkernel32.NewProc("GetNativeSystemInfo")
	procGetTickCount64            = modkernel32.NewProc("GetTickCount64")
	procGetSystemTimes            = modkernel32.NewProc("GetSystemTimes")
	procGlobalMemoryStatusEx      = modkernel32.NewProc("GlobalMemoryStatusEx")
	procReadProcessMemory         = modkernel32.NewProc("ReadProcessMemory")
	procGetProcessHandleCount     = modkernel32.NewProc("GetProcessHandleCount")
	procGetFileVersionInfoW       = modversion.NewProc("GetFileVersionInfoW")
	procGetFileVersionInfoSizeW   = modversion.NewProc("GetFileVersionInfoSizeW")
	procVerQueryValueW            = modversion.NewProc("VerQueryValueW")
	procGetProcessMemoryInfo      = modpsapi.NewProc("GetProcessMemoryInfo")
	procGetProcessImageFileNameA  = modpsapi.NewProc("GetProcessImageFileNameA")
	procEnumProcesses             = modpsapi.NewProc("EnumProcesses")
	procNtQueryInformationProcess = modntdll.NewProc("NtQueryInformationProcess")

	procCalculatePath = modNav.NewProc("CalculatePath")
	procFreePathArr   = modNav.NewProc("FreePathArr")
)

func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

func FreePathArr() {

	_, _, e1 := procFreePathArr.Call()
	if e1 != nil {
		fmt.Print(e1)
	}
}

func CalculatePath() {
	// var _s0 uint32
	// if straightPath {
	// 	_s0 = 1
	// }

	// startXYZ := start.toXYZ()
	// endXYZ := end.toXYZ()

	_, _, err := procCalculatePath.Call(1)
	if err != nil {
		fmt.Println(err)
	}

	// ret, _, e1 := syscall.SyscallN(
	// 	procCalculatePath.Addr(),
	// 	// uintptr(unsafe.Pointer(&startXYZ)),
	// 	// uintptr(unsafe.Pointer(&endXYZ)),
	// 	// uintptr(_s0),
	// )

	// if e1 != 0 {
	// 	err := errnoErr(e1)
	// 	fmt.Print(err)
	// } else {
	// 	err := syscall.EINVAL
	// 	fmt.Print(err)
	// }

	// fmt.Println(ret)
}

func GetTickCount64() (millis uint64, err error) {
	r0, _, e1 := syscall.Syscall(procGetTickCount64.Addr(), 0, 0, 0, 0)
	millis = uint64(r0)
	if millis == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// ReadProcessMemory reads from another process memory. The Handle needs to have
// the PROCESS_VM_READ right.
// A zero-byte read is a no-op, no error is returned.
func ReadProcessMemory(handle syscall.Handle, baseAddress uintptr, dest []byte) (numRead uintptr, err error) {
	n := len(dest)
	if n == 0 {
		return 0, nil
	}
	if err = _ReadProcessMemory(handle, baseAddress, uintptr(unsafe.Pointer(&dest[0])), uintptr(n), &numRead); err != nil {
		return 0, err
	}
	return numRead, nil
}

func _ReadProcessMemory(handle syscall.Handle, baseAddress uintptr, buffer uintptr, size uintptr, numRead *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procReadProcessMemory.Addr(), 5, uintptr(handle), uintptr(baseAddress), uintptr(buffer), uintptr(size), uintptr(unsafe.Pointer(numRead)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func abort(funcname string, err int) {
	panic(funcname + " failed: ")
}
