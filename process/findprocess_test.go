package findprocess

import (
	"fmt"
	"testing"
)

func TestWaitProcforExit(t *testing.T) {

	c := make(chan bool)
	go WaitForProcToExit("wintest.exe", c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Ended")
}
