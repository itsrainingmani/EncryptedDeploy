package main

import (
	"io/ioutil"
	"fmt"
	"EncryptedDeploy/filecrypto"
	"strings"
	"os/exec"
)

func main() {
	b, err := ioutil.ReadFile("encrypted.txt")
	if err != nil {
		fmt.Print(err)
	}

	passBytes, err := ioutil.ReadFile("pass.txt")
	if err != nil {
		fmt.Print(err)
	}

	decipheredtext, _ := filecrypto.Open(passBytes, b)
	credsForPowershell := string(decipheredtext)
	creds := strings.Split(string(decipheredtext), " ")
	uname, pwd := creds[0], creds[1]
	fmt.Println(uname, pwd)

	out, err := exec.Command("./sample.ps1", credsForPowershell).Output()
	fmt.Println(out)
}
