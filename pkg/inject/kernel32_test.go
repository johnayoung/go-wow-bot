package inject

import (
	"fmt"
	"testing"
)

func TestKernel32(t *testing.T) {
	tID, _ := GetCurrentThreadId()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(tID)
}
