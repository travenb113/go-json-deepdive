package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Species string `json:"carType"`
	// we can set the "omitempty" property as part of the JSON tag
	Description string `json:"what it does,omitempty"`
}

func main() {
	passengerCar := &Car{
		Species: "PassengerCar",
	}

	data, _ := json.Marshal(passengerCar)

	fmt.Println(string(data))

}
