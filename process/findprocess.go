package findprocess

import (
	"errors"
	"fmt"
	"time"

	"github.com/mitchellh/go-ps"
)

func findProcessByName(procName string) (bool, error) {
	procList, err := ps.Processes()
	if err != nil {
		return false, err
	}
	for _, p := range procList {
		// fmt.Println(p.Executable())
		if p.Executable() == procName {
			fmt.Println("Process found", p.Executable())
			return true, nil
		}
	}
	return false, errors.New("Process not found or exited")
}

//WaitForProcToExit is a function that runs a loop to determine if the process is running or not
func WaitForProcToExit(procName string, c chan bool) {
	for {
		time.Sleep(100 * time.Millisecond)
		procCheck, err := findProcessByName(procName)
		if err != nil {
			fmt.Println(err)
			close(c)
			return
		}
		fmt.Println("Process is still running")
		c <- procCheck
	}
}
