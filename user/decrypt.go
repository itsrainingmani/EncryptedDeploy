package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
	"github.com/tsmanikandan/EncryptedDeploy/crypto"
	"github.com/tsmanikandan/EncryptedDeploy/process"
)

func powershellScriptGeneration(creds []string, exeName string) string {
	cmdList := []string{
		"$usr = '" + creds[0] + "'",
		"Write-Host $usr",
		"$passwd = '" + creds[1] + "'",
		"Write-Host $passwd",
		"$credentials = New-Object System.Management.Automation.PSCredential -ArgumentList @($usr,(ConvertTo-SecureString -String $passwd -AsPlainText -Force))",
		"Start-Process powershell -Credential $credentials -ArgumentList '-noprofile -command &{Start-Process .\\" + exeName + " -verb runas}'",
	}

	return strings.Join(cmdList, ";\n")
}

func main() {

	// dataAssets := AssetNames()
	dataAssets := AssetNames()

	fmt.Println(dataAssets)

	// RestoreAsset(".", "sample.ps1")
	RestoreAsset(".", "wintest.exe")
	// defer os.Remove(".\\wintest.exe")
	// b, err := ioutil.ReadFile("encrypted.txt")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	b, err := Asset("encrypted.txt")
	if err != nil {
		fmt.Println(err)
	}

	var passString string

	fmt.Println("Please enter the passphrase:")
	fmt.Scan(&passString)

	passBytes, err := base64.StdEncoding.DecodeString(passString)
	if err != nil {
		fmt.Print(err)
	}

	decipheredtext, _ := crypto.Open(passBytes, b)
	credPair := strings.Split(string(decipheredtext), " ")

	psArgsgenerated := powershellScriptGeneration(credPair, "wintest.exe")
	// fmt.Println("Decrypted Credentials - ", credPair[0], credPair[1])

	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	// ... and interact with it
	stdout, stderr, err := shell.Execute(psArgsgenerated)
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout, stderr)

	time.Sleep(5000 * time.Millisecond)

	c := make(chan bool)
	go findprocess.WaitForProcToExit("wintest.exe", c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Ended")
	os.Remove(".\\wintest.exe")
	// psArgs := ".\\sample.ps1 " + string(decipheredtext)

	// cmd := exec.Command("powershell", psArgs)
	// fmt.Println("Running command and waiting for it to finish")
	// err = cmd.Run()
	// fmt.Println("Command finished with error: ", err)

}
