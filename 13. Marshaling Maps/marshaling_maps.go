package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// The keys need to be strings, the values can be
	// any serializable value
	carData := map[string]any{
		"carSounds": map[string]string{
			"car engine":        "vroom",
			"motorcycle engine": "rev",
		},
		"total count engine": 2,
	}

	// JSON encoding is done the same way as before
	data, _ := json.Marshal(carData)
	fmt.Println(string(data))
	//{"carSounds":{"car engine":"vroom","motorcycle engine":"rev"},"total count engine":2}
}
