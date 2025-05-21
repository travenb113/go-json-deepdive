package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Species     string
	Description string
}

func main() {
	carJson := `{"species": "car", "description": "drive on the road slowly/fast"}`
	var car Car
	json.Unmarshal([]byte(carJson), &car)
	fmt.Printf("Species: %s, Description: %s", car.Species, car.Description)
	//Species: car, Description: drive on the road slowly/fast
}
