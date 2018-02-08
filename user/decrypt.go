package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
	"github.com/tsmanikandan/EncryptedDeploy/crypto"
	"github.com/tsmanikandan/EncryptedDeploy/process"
	// "github.com/tsmanikandan/EncryptedDeploy/user/data"
)

// var pathOfExtractedFiles = "%temp%"
var exeToRun = "Setup.exe"
var exeToWatch = "mshta.exe"

func powershellScriptGeneration(creds []string, exeName string) string {
	cmdList := []string{
		"$usr = '" + creds[0] + "'",
		"$passwd = '" + creds[1] + "'",
		"$credentials = New-Object System.Management.Automation.PSCredential -ArgumentList @($usr,(ConvertTo-SecureString -String $passwd -AsPlainText -Force))",
		"Start-Process powershell -Credential $credentials -ArgumentList '-noprofile -command &{Start-Process .\\" + exeName + " -verb runas}'",
	}

	return strings.Join(cmdList, ";\n")
}

func main() {

	var passString string

	fmt.Println("Please enter the passphrase:")
	fmt.Scan(&passString)
	passBytes, err := base64.StdEncoding.DecodeString(passString)
	if err != nil {
		fmt.Print(err)
	}
	b, err := Asset("encrypted.txt")
	if err != nil {
		fmt.Println(err)
	}


	fmt.Println("\nDecrypting...")
	decipheredtext, err := crypto.Open(passBytes, b)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println("\nSuccessfully Decrypted")

	dataAssets := AssetNames()
	// dataAssets := data.AssetNames()

	// fmt.Println(dataAssets)

	fmt.Println("\nExtracting data...")
	for _, v := range dataAssets {
		if v != "encrypted.txt" {
			RestoreAssets(".", v)
		}
	}

	credPair := strings.Split(string(decipheredtext), " ")

	psArgsgenerated := powershellScriptGeneration(credPair, exeToRun)
	// fmt.Println("Decrypted Credentials - ", credPair[0], credPair[1])
	chDirErr := os.Chdir(".\\cisco")
	if chDirErr != nil {
		fmt.Println(chDirErr.Error())
		os.Exit(2)
	}
	// fmt.Println(os.Getwd())

	back := &backend.Local{}

	// start a local powershell process

	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	// defer shell.Exit()

	// time.Sleep(2000 * time.Millisecond)

	fmt.Println("\nRunning installation wizard...")
	// ... and interact with it
	stdout, stderr, err := shell.Execute(psArgsgenerated)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(stdout, stderr)

	time.Sleep(5000 * time.Millisecond)

	c := make(chan bool)
	go findprocess.WaitForProcToExit(exeToWatch, c)
	for i := range c {
		_ = i
	}
	fmt.Println("\nProgram Ended")
	shell.Exit()

	time.Sleep(5000 * time.Millisecond)

	chDirErr1 := os.Chdir("..")
	if chDirErr1 != nil {
		fmt.Println(chDirErr.Error())
		os.Exit(2)
	}
	// fmt.Println(os.Getwd())

	fmt.Println("\nRemoving Extracted data...")

	rmDirErr := os.RemoveAll(".\\cisco")
	if rmDirErr != nil {
		fmt.Println(rmDirErr.Error())
	}

	time.Sleep(5000 * time.Millisecond)
	
}
