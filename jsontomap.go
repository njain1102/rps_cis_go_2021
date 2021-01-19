package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	//Simple JSON object which we will parse
	coronaVirusJSON := `{
        "name" : "covid-11",
        "country" : "China",
        "city" : "Wuhan",
        "reason" : "Non veg Food"
    }`

	// Declared an empty map interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(coronaVirusJSON), &result)

	// Print the data type of result variable
	fmt.Println(reflect.TypeOf(result))

	// Reading each value by its key
	fmt.Println("Name :", result["name"],
		"\nCountry :", result["country"],
		"\nCity :", result["city"],
		"\nReason :", result["reason"])

	result["status"]=true
	fmt.Println(result["status"])
}