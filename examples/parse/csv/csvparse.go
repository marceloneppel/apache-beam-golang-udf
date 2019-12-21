package csv

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
)

func Parse(data string) string {
	reader := csv.NewReader(strings.NewReader(data))
	line, error := reader.Read()
	if error == io.EOF {
		return ""
	} else if error != nil {
		log.Fatal(error)
	}
	return "Firstname: " + line[0] + " - Lastname: " + line[1]
}
