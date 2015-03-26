package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("SpeedCoder, v1.0")
	fmt.Println("----------------")

	rand.Seed(time.Now().UTC().UnixNano())

	GuiMain()
}
