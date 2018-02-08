package main

import (
	// "fmt"
	"time"
)

func main() {
	datassets := AssetNames()
	for _, v := range datassets {
		RestoreAsset(".", v)
	}
	// fmt.Println(datassets)
	time.Sleep(10000 * time.Millisecond)
}
