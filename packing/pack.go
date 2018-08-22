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

		fmt.Println("\nInstalling dependencies")

		back := &backend.Local{}

		// start a local powershell process
		shell, err := ps.New(back)
		if err != nil {
			panic(err)
		}
		defer shell.Exit()

		// time.Sleep(2000 * time.Millisecond)
		// fmt.Println("\nInstalling dependencies")
		// out, err := exec.Command("go", "get ./...").Output()
		// if err != nil {
		// 	fmt.Printf("The error is - %v", err)
		// 	os.Exit(3)
		// }
		// fmt.Printf("The output is - %s", out)

		// time.Sleep(2000 * time.Millisecond)
		// fmt.Println("\nBuilding the End User Executable")
		// out, err := exec.Command("go", "build -o EndUser.exe -i").Output()
		// if err != nil {
		// 	fmt.Printf("The error is - %v", err)
		// }
		// fmt.Printf("The output is - %s", out)
		//... and interact with it
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("\nBuilding the End User Executable")
		stdout, stderr, err := shell.Execute("go get -d ./...; go build -o EndUser.exe -i decrypt.go data.go")
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(stdout, stderr)
		fmt.Println("End User Executable has been generated")
		time.Sleep(5000 * time.Millisecond)
	} else {
		fmt.Println("The files required to build the End User executables aren't present")
		time.Sleep(2000 * time.Millisecond)
	}
}
