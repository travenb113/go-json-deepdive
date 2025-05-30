package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Species     string
	Description *string
}

func main() {
	passengerCar := &Car{
		Species:     "Passenger Car",
		Description: nil,
	}

	data, _ := json.Marshal(passengerCar)
	fmt.Println(string(data))
	// {"Species":"Passenger Car","Description":null}
}
