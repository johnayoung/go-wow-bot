package state

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"strconv"
)

type Cell struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type SquareReader struct {
	img *image.RGBA
}

func NewSquareReader(img *image.RGBA) SquareReader {
	return SquareReader{img}
	// return SquareReader{pixels: pixels}
}

func RGBtoHex(color color.RGBA) string {
	return fmt.Sprintf("%02x%02x%02x", color.R, color.G, color.B)
}

// function to find the hexidecimal color of a pixel based on a given x,y coordinate
func (s *SquareReader) GetColorAtCell(cell Cell) string {
	return RGBtoHex(s.img.RGBAAt(cell.X, cell.Y))
}

// Converts a cell's RGBA color code to decimal data
func (s *SquareReader) GetIntAtCell(cell Cell) int64 {
	color := s.GetColorAtCell(cell)

	num, err := strconv.ParseInt(color, 16, 64)

	if err != nil {
		log.Printf("Error at SquareReader [GetIntAtCell]: %v", err)
		return 0
	}

	return num
}

// Converts a cell's hexidecimal color to a 6 point decimal
func (s *SquareReader) GetFixedPointAtCell(cell Cell) int64 {
	return s.GetIntAtCell(cell) / 100000
}

// Converts a cell's hexidecimal color to a 3 character string
func (s *SquareReader) GetStringAtCell(cell Cell) string {
	color := s.GetIntAtCell(cell)

	if color != 0 {
		color := strconv.FormatInt(color, 6)

		var word string

		// Iterates through each ASCII code and sets it equal to relevant character
		for i := 0; i < 3; i++ {
			first := i * 2
			second := (i + 1) * 2

			log.Printf("First: %d, Second: %d", first, second)

			char := color[first:second]

			log.Printf("Char: %v", char)

			charInt, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal(err)
			}
			word += strconv.QuoteRune(rune(charInt))

			fmt.Printf("Character corresponding to Ascii Code %d = %v\n", charInt, word)

		}
		return word
	}

	return ""
}
