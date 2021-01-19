package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	//array JSON object which we will parse
	//array of json
	coronaVirusJSON := `[{
        "name" : "covid-11",
        "country" : "China",
        "city" : "Wuhan",
        "reason" : "Non veg Food",
        "medicines":{
            "name":"CovidSheild",
            "country":"India"
          }
    },
	{
        "name" : "covid-17",
        "country" : "UK",
        "city" : "London",
        "reason" : "Non veg Food",
 		"medicines":{
            "name":"Covaxin",
            "country":"US"
          }
    }

]`

	// Declared an empty map interface
	var result []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(coronaVirusJSON), &result)

	// Print the data type of result variable
	fmt.Println(reflect.TypeOf(result))

	// Reading each value by its key
	for key,value:= range result {
		fmt.Printf("%d=>%s\n",key,value)
	}
	result[0]["status"]=true
	fmt.Println(result[0]["status"])
}
