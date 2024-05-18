package fastjson_builder

import (
	"fmt"
	"log"
	"runtime/debug"
)

func printModuleVersion() {
	if bi, exists := debug.ReadBuildInfo(); exists {
		fmt.Println(bi.Main.Version)
	} else {
		log.Printf("No version information found. Make sure to use " +
			"GO111MODULE=on when running 'go get' in order to use specific " +
			"version of the binary.")
	}
}
