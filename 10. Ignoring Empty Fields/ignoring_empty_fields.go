package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Species string `json:"-"`
}

func main() {
	passengerCar := &Car{
		Species: "PassengerCar",
	}

	data, _ := json.Marshal(passengerCar)

	fmt.Println(string(data))
	// {}
}
