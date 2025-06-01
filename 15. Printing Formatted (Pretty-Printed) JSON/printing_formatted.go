package main

import (
	"encoding/json"
	"fmt"
)

// Car struct with nested Dimensions
type Car struct {
	Species     string     `json:"Species"`
	Description string     `json:"Description"`
	Dimensions  Dimensions `json:"Dimensions"`
}

// Dimensions struct with custom JSON marshaling
type Dimensions struct {
	EngineCapacity float32
	Weight         int
}

// MarshalJSON implements custom JSON formatting: "3.5x1250"
func (d Dimensions) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%.1fx%d"`, d.EngineCapacity, d.Weight)), nil
}

func main() {
	car := Car{
		Species:     "passengerCar",
		Description: "drive on the road slowly/fast",
		Dimensions: Dimensions{
			EngineCapacity: 3.5,
			Weight:         1250,
		},
	}

	// Use json.MarshalIndent to pretty-print the JSON
	data, err := json.MarshalIndent(car, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
	/* {
	  		"Species": "passengerCar",
	  		"Description": "drive on the road slowly/fast",
	  		"Dimensions": "3.5x1250"
		} */
}
