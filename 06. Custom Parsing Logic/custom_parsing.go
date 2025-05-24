package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Dimensions struct {
	EngineCapacity float32
	Weight         int
}

// UnmarshalJSON parses a JSON string in the format
// "EngineCapacityxWeight", for example: "3.5x1250",
// and stores the values into the Dimensions struct
func (d *Dimensions) UnmarshalJSON(data []byte) error {
	// "data" is expected to be a JSON string as a byte slice
	// for example: `"3.5x1250"`

	// Check for minimum length (must at least include quotes and content)
	if len(data) < 2 {
		return fmt.Errorf("dimensions string too short")
	}

	// Remove the surrounding quotes
	s := string(data)[1 : len(data)-1]

	// Split the string into two parts using 'x' as a separator
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return fmt.Errorf("dimensions string must contain two parts")
	}

	// Convert the first part to float (engine capacity)
	engineCapacity, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return fmt.Errorf("engine capacity must be a float")
	}

	// Convert the second part to int (weight)
	weight, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("weight must be an int")
	}

	// Assign the parsed values to the Dimensions struct fields
	d.EngineCapacity = float32(engineCapacity)
	d.Weight = weight
	return nil
}

type Car struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func main() {
	carJson := `{"species": "car","description": "drive on the road slowly/fast", "dimensions": "3.5x1250"}`
	var car Car
	if err := json.Unmarshal([]byte(carJson), &car); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(car)
	//{car drive on the road slowly/fast {3.5 1250}}
}
