package main

import (
	"encoding/base64"
	"fmt"
	"strings"
	// "os"
	// "os/exec"

	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"

	"github.com/tsmanikandan/EncryptedDeploy/crypto"
)

func powershellScriptGeneration(creds []string, exeName string) []string {
	cmdList := []string{
		"$usr = " + creds[0],
		"Write-Host $usr",
		"$passwd = " + creds[1],
		"Write-Host $passwd",
		"$credentials = New-Object System.Management.Automation.PSCredential -ArgumentList @($usr,(ConvertTo-SecureString -String $passwd -AsPlainText -Force))",
		"Start-Process powershell -Credential $credentials -ArgumentList '-noprofile -command &{Start-Process .\\" + exeName + " -verb runas}'",
	}
	return cmdList
}

func main() {

	dataAssets := AssetNames()
	// RestoreAsset(".", "sample.ps1")
	RestoreAsset(".", dataAssets[0])
	// b, err := ioutil.ReadFile("encrypted.txt")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	b, err := Asset(dataAssets[1])
	if err != nil {
		fmt.Println(err)
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
	credPair := strings.Split(string(decipheredtext), " ")

	fmt.Println("Decrypted Credentials - ", credPair)

	scrpt := powershellScriptGeneration(credPair, dataAssets[0])

	back := &backend.Local{}

	shell, err := ps.New(back)
	if err != nil {
		fmt.Println(err)
	}
	defer shell.Exit()

	stdout, stderr, err := shell.Execute(strings.Join(scrpt, "\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdout, stderr)

	// fmt.Println(stdout, stderr)

	// credsForPowershell := ".\\sample.ps1 " + string(decipheredtext)
	// fmt.Println(credsForPowershell)

	// cmd := exec.Command("powershell", credsForPowershell)
	// fmt.Println("Running command and waiting for it to finish")
	// err = cmd.Run()
	// fmt.Println("Command finished with error: ", err)
	// os.Remove("sample.ps1")
	// os.Remove("wintest.exe")
}
