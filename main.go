package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	thrustPin1 = rpio.Pin(17)
	thrustPin2 = rpio.Pin(27)

	steeringPin1 = rpio.Pin(15)
	steeringPin2 = rpio.Pin(18)
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	thrustPin1.Output()
	thrustPin2.Output()

	steeringPin1.Output()
	steeringPin2.Output()

	pins := []rpio.Pin{thrustPin1, thrustPin2}

	for i := 1; i <= 5; i++ {
		pins[i%2].High()
		pins[1-i%2].Low()
		time.Sleep(1 * time.Second)
	}

	thrustPin1.Low()
	thrustPin2.Low()
}
