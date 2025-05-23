package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	dateJson := `"2025-05-23T12:25:45.567Z"`

	var date time.Time

	json.Unmarshal([]byte(dateJson), &date)

	fmt.Println(date)
}
