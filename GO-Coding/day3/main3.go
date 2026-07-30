// Day 3:- Task: Create a struct with a field of type json.RawMessage. Unmarshal a JSON payload into it, then conditionally unmarshal that raw message into two different structs based on a "type" field.
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	contents := []byte(`
    {
        "name":"value",
        "Id":"value2"
    }
`)
	mapStructure := make(map[string]string)
	_ = json.Unmarshal(contents, &mapStructure)

	fmt.Println("From Map:")
	fmt.Println(mapStructure["name"])
	fmt.Println(mapStructure["Id"])

	type Custom struct {
		Name string // Must be exported ###Caution-struct field names should start With CAPITAL LETTER
		ID   string // Must be exported.
	}
	customStructure := &Custom{"Name:JSON", "ID:404"}
	err := json.Unmarshal(contents, customStructure) //###Caution-customStructure is already an pointer to Custom struct no need to pass it as pointer again
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nFrom Struct:")
	fmt.Println("Name:", customStructure.Name)

	fmt.Println("Raw ID:", string(customStructure.ID))
}
