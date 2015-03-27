package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	fmt.Println("SpeedCoder, v0.1")
	fmt.Println("----------------")

	if len(os.Args) != 3 {
		fmt.Printf("Usage must be: %s <programming language> <keyword>\n", os.Args[0])
		os.Exit(1)
	}

	// seed the RNG so we get a different code snippet every time.
	rand.Seed(time.Now().UTC().UnixNano())

	GuiMain(os.Args[1], os.Args[2])
}
