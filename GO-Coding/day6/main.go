//Day6:-Task-Custom MarshalJSON (Decimal Precision): Create a struct with a float64 representing a price.
// Write a custom MarshalJSON that formats it as a string with exactly 2 decimal places.
// Notice: Floats can lose precision; financial APIs often require strict string formatting.

package main

import (
	"encoding/json"
	"fmt"
)

type price struct {
	Value float64
}

func (p price) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("%.2f", p.Value)
	return json.Marshal(formatted)
}

func main() {
	p := price{Value: 123.45678}
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error Marshing into JSON %v", err)
	}
	fmt.Printf("Value before Marshing %v,Value after Marshing %v \n", p, string(data)) //without string output:-[34 49 50 51 46 52 54 34],with string output:- "123.46" because data here is in form of []byte array using string to convert it into string
}
