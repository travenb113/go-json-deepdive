package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Car struct {
	Species     string
	Description string
	ReleaseTime time.Time
}

func main() {
	carJson := `{"species": "car", "description": "drive on the road slowly/fast", "releaseTime": "2025-05-24T12:45:23.234Z"}`
	var car Car
	json.Unmarshal([]byte(carJson), &car)
	fmt.Println(car)
}
