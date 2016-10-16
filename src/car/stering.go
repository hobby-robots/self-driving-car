package car

import (
	"github.com/stianeikeland/go-rpio"
	"fmt"
)

type Steering interface {
	Left()
	Right()
	Forward()
	Backward()
	ResetSteering()
	ResetThrusting()
	Close()
}

type steering struct {
	thrustPin1   rpio.Pin
	thrustPin2   rpio.Pin
	steeringPin1 rpio.Pin
	steeringPin2 rpio.Pin
}

type logging struct {
}

func DebugSteering() Steering {
	return &logging{}
}

func NewSteering(tp1, tp2, sp1, sp2 uint8) Steering {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	steering := &steering{
		thrustPin1: rpio.Pin(tp1),
		thrustPin2: rpio.Pin(tp2),
		steeringPin1: rpio.Pin(sp1),
		steeringPin2: rpio.Pin(sp2),
	}

	steering.steeringPin1.Output()
	steering.steeringPin2.Output()
	steering.thrustPin1.Output()
	steering.thrustPin2.Output()

	return steering
}


func (s *steering) Left() {
	s.steeringPin1.Low()
	s.steeringPin2.High()
}

func (s *steering) Right() {
	s.steeringPin1.High()
	s.steeringPin2.Low()
}

func (s *steering) Forward() {
	s.thrustPin1.High()
	s.thrustPin2.Low()
}

func (s *steering) Backward() {
	s.thrustPin1.Low()
	s.thrustPin2.High()
}

func (s *steering) ResetSteering() {
	s.steeringPin1.Low()
	s.steeringPin2.Low()
}

func (s *steering) ResetThrusting() {
	s.thrustPin1.Low()
	s.thrustPin2.Low()
}

func (s *steering) Close(){
	s.steeringPin1.Low()
	s.steeringPin2.Low()
	s.thrustPin1.Low()
	s.thrustPin2.Low()
}


func (s *logging) Left() {
	fmt.Println("Turning Left...")
}

func (s *logging) Right() {
	fmt.Println("Turning Right...")
}

func (s *logging) Forward() {
	fmt.Println("Moving Forward...")
}

func (s *logging) Backward() {
	fmt.Println("Moving Backward...")
}

func (s *logging) ResetSteering() {
	fmt.Println("ResetSteering")
}

func (s *logging) ResetThrusting() {
	fmt.Println("ResetThrusting")
}

func (s *logging) Close(){
	fmt.Println("Close...")
}
