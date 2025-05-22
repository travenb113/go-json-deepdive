package main

import (
	"encoding/json"
	"fmt"
)

type Dimensions struct {
	EngineCapacity float32
	Weight         int
}

type Car struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func main() {
	carJson := `{"species": "car","description": "drive on the road slowly/fast", "dimensions":{"EngineCapacity": 3.5, "Weight": 1250}}`
	var car Car
	json.Unmarshal([]byte(carJson), &car)
	fmt.Println(car)
	//{car drive on the road slowly/fast {3.5 1250}}
}
