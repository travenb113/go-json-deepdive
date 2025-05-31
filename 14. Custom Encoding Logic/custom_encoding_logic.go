package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Species    string
	Dimensions Dimensions
}

type Dimensions struct {
	EngineCapacity float32
	Weight         int
}

// marshals a Dimensions struct into a JSON string
// with format "engineCapacityxweight"
func (d Dimensions) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%.1fx%d"`, d.EngineCapacity, d.Weight)), nil
}

func main() {
	car := Car{
		Species: "passengerCar",
		Dimensions: Dimensions{
			EngineCapacity: 3.5,
			Weight:         1250,
		},
	}
	CarJson, _ := json.Marshal(car)
	fmt.Println(string(CarJson))
	// {"Species":"passengerCar","Dimensions":"3.5x1250"}
}
