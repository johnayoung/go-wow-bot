package input

import (
	"fmt"
	"testing"

	hook "github.com/robotn/gohook"
)

func TestInput(t *testing.T) {
	lshift := uint16(160)

	hook.Register(hook.KeyHold, []string{}, func(e hook.Event) {
		isLeftShift := e.Rawcode == lshift

		if e.Kind == hook.KeyHold && isLeftShift {
			fmt.Println("Status: Paused")
		}
	})

	hook.Register(hook.KeyUp, []string{}, func(e hook.Event) {
		isLeftShift := e.Rawcode == lshift

		if e.Kind == hook.KeyUp && isLeftShift {
			fmt.Println("Status: Running")
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
