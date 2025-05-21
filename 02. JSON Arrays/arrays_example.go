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
	carJson := `[{"species": "car", "description": "drive on the road slowly/fast"},{"species": "lexus", "description": "brand of car"}]`
	var cars []Car
	json.Unmarshal([]byte(carJson), &cars)
	fmt.Printf("Cars : %+v", cars)
	//Cars : [{Species:car Description:drive on the road slowly/fast} {Species:lexus Description:brand of car}]
}
