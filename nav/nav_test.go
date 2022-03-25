package nav

import (
	"testing"
)

// var objectManagerPtr
const OBJECT_MANAGER_BASE = 0x00E18DB4
const OBJECT_MANAGER_MAP_ID_OFFSET = 0xCC
const WALL_CLIMB_ANGLE = 0x00A37F0C
const BUILD_NUMBER = 0x00B3203C
const AUTO_LOOT_POINTER = 0x00BD0914
const AUTO_LOOT_OFFSET = 0x30
const AUTO_LOOT = 0x2b89fa90

func TestNav(t *testing.T) {
	// start := navigation.NewXYZ(-2341.46, -378.072, -8.16762)
	// end := navigation.NewXYZ(-2397.07, -384.831, -2.27768)

	// length := 1
	// ret := navigation.NavigationGetInstance().CalculatePath(1, start, end, true, &length)
	// kernel32 := syscall.NewLazyDLL("kernel32.dll")

	// mapsPath := filepath.Join(utils.RootDir(), "Bot", "Navigation.dll")

	// _, err := syscall.LoadLibrary(mapsPath)
	// if err != nil {
	// 	panic(err)
	// }
	// defer FreePathArr()

	CalculatePath()

	// var _s0 uint32 = 1
	// startXYZ := NewXYZ(-2341.46, -378.072, -8.16762)
	// endXYZ := NewXYZ(-2397.07, -384.831, -2.27768)
	// length := 1

	// dm := goxymemmory.DataManager("Ascension.exe") //Get the DataManager with the process passed.
	// if !dm.IsOpen {                                //Check if process was opened.
	// 	fmt.Printf("Failed opening process.\n")
	// 	return
	// }

	// data, err := dm.Read(AUTO_LOOT_POINTER, goxymemmory.INT) //Reads the int.
	// if err != nil {                                          //Check if not failed.
	// 	fmt.Printf("Failed reading memory. %s", err)
	// }
	// fmt.Println("INT: ", data.Value)

	// // Gives me an INT value, I need to convert to Hexadecimal...this is my address
	// h := fmt.Sprintf("%x", data.Value.(int)+48)
	// value, err := strconv.ParseInt(h, 16, 64)

	// autoloot, err := dm.Read(uint(value), goxymemmory.INT) //Reads the int.
	// if err != nil {                                        //Check if not failed.
	// 	fmt.Printf("Failed reading memory. %s", err)
	// }
	// fmt.Println("INT: ", autoloot.Value)
}
