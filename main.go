package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

var (
	pin1 = rpio.Pin(17)
	pin2 = rpio.Pin(27)

	pin3 = rpio.Pin(15)
	pin4 = rpio.Pin(18)
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin1.Output()
	pin2.Output()

	pins := []rpio.Pin{pin1, pin2}

	for i := 1; i <= 5; i++ {
		pins[i%2].High()
		pins[1-i%2].Low()
		time.Sleep(1 * time.Second)
	}

	pin1.Low()
	pin2.Low()
}
