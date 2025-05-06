package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"name"`
	Address    string   `json:"addr"`
	Age        int      `json:"-"`
	FaveColors []string `json:"favcolurs,omitempty"`
}

func encodeExample() {
	// TODO: create some people data
	people := []person{
		{"Jane Doe", "123 Anywhere Street", 35, nil},
		{"John Public", "456 Everywhere Blvd", 31, []string{"Purple", "Yellow", "Green"}},
	}

	// TODO: Marshal is used to convert a data structure to JSON format
	result, err := json.MarshalIndent(people, "", "	")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result)
	// MarshalIndent is used to format the JSON string with indentation

}

func main() {
	// Encode Go data as JSON
	encodeExample()
}
