package main

import (
	"encoding/json"
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	Brake_pedal    uint16  `json:"brake_pedal"`
	Steering_whell int16   `json:"steering_whell"` // -32k - +32K
	Top_speed_kmh  float64 `json:"top_speed_kmh"`
	Gas_pedal      uint16  `json:"gas_pedal"` // min 0 max 65535
}

func (c car) kmh() float64 {
	return float64(c.Gas_pedal) * (c.Top_speed_kmh / usixteenbitmax)
}

func main() {
	a_car := &car{Gas_pedal: 2234,
		Brake_pedal: 0, Steering_whell: 12561,
		Top_speed_kmh: 225.0,
	}
	// a_car = &car{2234, 0, 12561, 225.0}

	fmt.Println(a_car.kmh())

	b, err := json.Marshal(a_car)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
