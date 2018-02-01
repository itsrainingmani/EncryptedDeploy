package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/tsmanikandan/EncryptedDeploy/crypto"
)

func main() {
	b, err := ioutil.ReadFile("encrypted.txt")
	if err != nil {
		fmt.Print(err)
	}

	var passString string

	fmt.Println("Please enter the passphrase:")
	fmt.Scan(&passString)

	// passBytes, err := ioutil.ReadFile("pass.txt")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	passBytes, err := base64.StdEncoding.DecodeString(passString)
	if err != nil {
		fmt.Print(err)
	}

	decipheredtext, _ := crypto.Open(passBytes, b)
	credsForPowershell := string(decipheredtext)
	creds := strings.Split(string(decipheredtext), " ")
	uname, pwd := creds[0], creds[1]
	fmt.Println(uname, pwd)

	credsForPowershell = ".\\sample.ps1 " + credsForPowershell

	out, _ := exec.Command("powershell", credsForPowershell).Output()
	fmt.Println(out)
}
