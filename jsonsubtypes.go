package main
import (
"encoding/json"
"fmt"
)

type Dimensions struct {
Height int
Width  int
}

type BBBird struct {
Species     string
Description string
Dimensions  []Dimensions
}

func main() {
birdJson := `{"species":"pigeon","description":"likes to perch on rocks", 
"dimensions":[{"height":24,"width":10},{"height":24,"width":10}]}`
var bird BBBird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Println(bird)
// {pigeon likes to perch on rocks {24 10}}
}
