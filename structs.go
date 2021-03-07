package main

import (
	"encoding/json"
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

// Car of struct is an example
type Car struct {
	BrakePedal    uint16  `json:"brake_pedal"`
	SteeringWhell int16   `json:"steering_whell"` // -32k - +32K
	TopSpeedKmh   float64 `json:"top_speed_kmh"`
	GasPedal      uint16  `json:"gas_pedal"` // min 0 max 65535
}

// Value Receivers
func (c Car) kmh() float64 {
	return float64(c.GasPedal) * (c.TopSpeedKmh / usixteenbitmax)
}

// Pointer Receivers
func (c *Car) newTopSpeed(newspeed float64) {
	c.TopSpeedKmh = newspeed
}

// This should be not changed
func (c Car) newTopSpeedNotPointer(newspeed float64) {
	c.TopSpeedKmh = newspeed
}

func (c *Car) toJSONString() string {
	b, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func TestStructs() {
	aCar := &Car{GasPedal: 2234,
		BrakePedal: 0, SteeringWhell: 12561,
		TopSpeedKmh: 225.0,
	}
	// a_car = &car{2234, 0, 12561, 225.0}

	fmt.Println(aCar.kmh())
	fmt.Println(aCar.toJSONString())
	aCar.newTopSpeed(100)
	fmt.Println(aCar.toJSONString())
	// not changed
	aCar.newTopSpeedNotPointer(150)
	fmt.Println(aCar.toJSONString())
}
