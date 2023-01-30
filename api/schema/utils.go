package schema

import (
	"encoding/json"
	"log"
	"os"
)

func GetJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true

	content, err := os.ReadFile(fileName)

	if err != nil {
		isOK = false
		log.Fatalf("Error reading file %s. Error: %v", fileName, err)
	}

	err = json.Unmarshal(content, result)

	if err != nil {
		isOK = false
		log.Fatalf("Error parsing JSON of %s. Error: %v", fileName, err)
	}

	return
}
