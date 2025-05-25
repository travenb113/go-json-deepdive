package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Species     string `json:"carType"`
	Description string `json:"what it does"`
}

func main() {
	carJson := `{"carType": "passenger car", "what it does": "drive on the road slowly/fast"}`
	var car Car
	json.Unmarshal([]byte(carJson), &car)
	fmt.Println(car)
	//{passenger car drive on the road slowly/fast}
}
