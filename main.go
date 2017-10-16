package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	pin := rpio.Pin(10)
	fmt.Println(3)
	pin.Output()
}
