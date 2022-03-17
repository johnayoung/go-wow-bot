package state

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/johnayoung/go-wow-bot/utils"
)

const startX = 2
const startY = 32

func GetBitmapCoords() {
	// reader := NewSquareReader()
	// metaData := reader.GetIntAtCell(Cell{X: startX, Y: startY})
	// metaString := fmt.Sprintf("%d", metaData)

	numberOfFrames := 130
	frameRows := 1
	cellSize := 6
	// fRows := metaString[len(metaString)-4 : len(metaString)-3]
	// cSize := metaString[len(metaString)-6 : len(metaString)-5]

	frameCols := math.Ceil(float64(numberOfFrames) / float64(frameRows))
	bitMapX1 := startX
	bitMapY1 := startY
	bitMapX2 := frameCols * float64(cellSize)
	bitMapY2 := frameRows * cellSize
	dataPointCoordinateArray := []Cell{
		{X: bitMapX1, Y: bitMapY1},
	}

	dataBitMap := robotgo.CaptureScreen(bitMapX1, bitMapY1, int(frameCols), bitMapY2)

	for frame := 1; frame < numberOfFrames+1; frame++ {
		for pixelX := 0; pixelX < int(bitMapX2); pixelX++ {
			for pixelY := 0; pixelY < bitMapY2; pixelY++ {
				colorHex := robotgo.GetColor(dataBitMap, pixelX, pixelY)
				color, _ := strconv.ParseInt(fmt.Sprintf("%d", colorHex), 10, 64)

				if int(color) == frame {
					dataPointCoordinateArray = append(dataPointCoordinateArray, Cell{X: pixelX, Y: pixelY + startY})

					pixelX = 8000000000000000
					pixelY = 8000000000000000
				}

				if bitMapY2 > 1000 || bitMapX2 > 1000 {
					panic("ridiculously large")
				}
			}
		}
	}

	fmt.Println(len(dataPointCoordinateArray))

	if len(dataPointCoordinateArray) > 10 {
		file, _ := json.MarshalIndent(dataPointCoordinateArray, "", " ")

		pathToJson := filepath.Join(utils.RootDir(), "state", "coordinates_test.json")
		err := ioutil.WriteFile(pathToJson, file, 0644)
		if err != nil {
			panic(err)
		}

	}
}

func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}
