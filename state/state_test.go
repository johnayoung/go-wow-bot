package state

import (
	"fmt"
	"testing"

	"github.com/go-vgo/robotgo"
)

func TestState(t *testing.T) {
	for {
		title := robotgo.GetTitle()
		fmt.Println(title)
	}
}
