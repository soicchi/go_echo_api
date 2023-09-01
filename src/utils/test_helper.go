package utils

import (
	"encoding/json"
)

func GenerateExpectedJSON(expectedString string) string {
	expectedJSON, _ := json.Marshal(expectedString)
	return string(expectedJSON) + "\n"
}
