package main

import (
	"encoding/json"
	"fmt"
)

// The same json tags will be used to encode data into JSON
type EBird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func main() {
	pigeon := &EBird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// we can use the json.Marshal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))

	// The keys need to be strings, the values can be
	// any serializable value
	birdData := map[string]interface{}{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	// JSON encoding is done the same way as before
	data, _ = json.Marshal(birdData)
	fmt.Println(string(data))
}
