package xml

import (
	"encoding/xml"
	"log"
)

type Persons struct {
	XMLName xml.Name `xml:"persons"`
	Persons []Person `xml:"person"`
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	Firstname string   `xml:"first_name"`
	Lastname  string   `xml:"last_name"`
}

func Parse(data string) string {
	var mappedData Persons
	err := xml.Unmarshal([]byte(data), &mappedData)
	if err != nil {
		log.Fatal(err)
	}
	if len(mappedData.Persons) > 0 {
		return "Firstname: " + mappedData.Persons[0].Firstname + " - Lastname: " + mappedData.Persons[0].Lastname
	}
	return ""
}
