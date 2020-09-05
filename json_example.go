package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

func main() {
	// ' ' equivalen to """ """ python
	// simulating the content of json file
	json_file := `
{"species": "pigeon",
"description": "likes to perch on rocks"}
`
	birdJson := json_file
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)

	fmt.Printf("%+v\n", bird)

	fmt.Println()
	json_file = `
	[
		{
		  "species": "pigeon",
		  "decription": "likes to perch on rocks"
		},
		{
		  "species":"eagle",
		  "description":"bird of prey"
		}
	]
	`
	birdJson = json_file
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("%+v\n", birds)

}
