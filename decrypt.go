package main

import (
	"io/ioutil"
	"fmt"
	"EncryptedDeploy/filecrypto"
	"github.com/danieljoos/wincred"
	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
)

func main() {
	b, err := ioutil.ReadFile("encrypted-credentials.txt")
	if err != nil {
		fmt.Print(err)
	}

	passBytes, err := ioutil.ReadFile("pass.txt")
	if err != nil {
		fmt.Print(err)
	}

	decipheredtext, _ := filecrypto.Open(passBytes, b)
	fmt.Println(string(decipheredtext))
}
