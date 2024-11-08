package utils

import (
	"encoding/json"
	"log"
)

func PrettifyObject(o any) string {
	jsonData, err := json.MarshalIndent(o, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}
