package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	fmt.Println("builder module")

	i, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("failed to read the build info")
	}
	v := i.Main.Version

	fmt.Printf("Version: %v\n", v)
}
