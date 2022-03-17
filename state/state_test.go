package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	num1 := 7249
	num2 := 6930
	diff := float64(num2) / float64(num1)

	fmt.Println(diff)

	perc := diff * 100

	fmt.Println(perc)
}
