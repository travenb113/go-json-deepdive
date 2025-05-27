package main

import (
	"encoding/json"
	"fmt"
)

// The same json tags will be used to encode data into JSON
type Car struct {
	Species     string `json:"carType"`
	Description string `json:"what it does"`
}

func main() {
	passengerCar := &Car{
		Species:     "Passenger Car",
		Description: "Drive on the road slowly/fast",
	}

	// we can use the json.Marshal function to
	// encode the passengerCar variable to a JSON string
	data, _ := json.Marshal(passengerCar)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))
	//{"carType":"Passenger Car","what it does":"Drive on the road slowly/fast"}
}
