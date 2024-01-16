package util

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"
)

func JsonToString(input interface{}) string {
	b, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return ""
	}

	return string(b)
}

func ValidateStringContainInStringArray(listString []string, key string) bool {
	for i := 0; i < len(listString); i++ {
		if listString[i] == key {
			return true
		}
	}
	return false
}
