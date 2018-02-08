package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
)

var (
	decryptFileName = "decrypt.go"
	dataFileName    = "data.go"
	foundFileCount  = 0
)

func decryptAndDataExists() bool {
	curDir, _ := os.Getwd()
	err := filepath.Walk(curDir, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, decryptFileName) || strings.Contains(path, dataFileName) {
			foundFileCount++
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Walk error [%v]\n", err)
		return false
	}
	// fmt.Println(foundFileCount)
	if foundFileCount == 2 {
		return true
	}
	return false
}

func main() {
	if decryptAndDataExists() {

		time.Sleep(2000 * time.Millisecond)

		back := &backend.Local{}

		// start a local powershell process
		shell, err := ps.New(back)
		if err != nil {
			panic(err)
		}
		defer shell.Exit()

		time.Sleep(2000 * time.Millisecond)
		fmt.Println("\nBuilding the End User Executable")
		// ... and interact with it
		stdout, stderr, err := shell.Execute("go.exe build -o EndUser.exe")
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(stdout, stderr)
		time.Sleep(2000 * time.Millisecond)
	} else {
		fmt.Println("The files required to build the End User executables aren't present")
		time.Sleep(2000 * time.Millisecond)
	}
}
