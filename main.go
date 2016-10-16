package main

import (
	"github.com/hobby-robots/self-driving-car/src/car"

	"os"
	"fmt"
)

func main() {
	steering := car.NewSteering(17, 27, 15, 18)
	//steering := car.DebugSteering()
	defer steering.Close()

	path := "/static"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fmt.Printf("Serving static %s\n", os.Args)

	server := car.NewServer(8080, steering, path)
	if err := server.Start(); err != nil {
		panic(err)
	}
	server.Wait()
}

