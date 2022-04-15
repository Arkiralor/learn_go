package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Go's equivalent to Python Dictionaries
var example = map[string]interface{}{
	"name":  "Gopher",
	"title": "programmer",
	"contact": map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	},
	"children": []string{"Mary", "Susan"},
	"pets": map[string]interface{}{
		"dog":  "Fido",
		"cat":  "Tigger",
		"fish": []string{"angel", "clownfish", "surgeonfish"},
	},
}

// Main entry point for the program
func main() {

	json_str, err := MapToJson(example)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(json_str)

}

// Function to convert a map to a JSON string
func MapToJson(mapped_variable map[string]interface{}) (string, error) {
	json_str, err := json.Marshal(mapped_variable)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return string(json_str), err
}
