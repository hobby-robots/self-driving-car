package main

import (
	"github.com/hobby-robots/self-driving-car/src/car"

)

func main() {
	//steering := car.NewSteering(17, 27, 15, 18)
	steering := car.DebugSteering()

	defer steering.Close()

	server := car.NewServer(8080, steering)
	if err := server.Start(); err != nil {
		panic(err)
	}
	server.Wait()
}

