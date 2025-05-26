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
	carJson := `{"cars":{"passenger cars":"drive on the road slowly/fast","lexus": "brand of car"`
	var result map[string]any
	json.Unmarshal([]byte(carJson), &result)

	// The object stored in the "cars" key is also stored as
	// a map[string]any type, and its type is asserted from
	// the `any` type
	cars := result["cars"].(map[string]any)

	for key, value := range cars {
		// Each value is an `any` type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}

}
