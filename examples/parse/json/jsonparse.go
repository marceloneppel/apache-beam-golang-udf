package json

import (
	"encoding/json"
	"log"
)

type Person struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

func Parse(data string) string {
	var mappedData []Person
	err := json.Unmarshal([]byte(data), &mappedData)
	if err != nil {
		log.Fatal(err)
	}
	if len(mappedData) > 0 {
		return "Firstname: " + mappedData[0].Firstname + " - Lastname: " + mappedData[0].Lastname
	}
	return ""
}
