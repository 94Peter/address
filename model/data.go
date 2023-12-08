package model

import (
	"encoding/json"
	"os"
)

var stateData map[string]map[string]string

func InitData(file string) error {
	// Read the JSON file
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	// Decode the JSON data into stateData
	err = json.Unmarshal(data, &stateData)
	if err != nil {
		return err
	}

	return nil
}

func getCity(state string) (map[string]string, bool) {
	data, ok := stateData[state]
	return data, ok
}

func getCode(state, city string) (string, bool) {
	data, ok := stateData[state]
	if !ok {
		return "", false
	}
	code, ok := data[city]
	return code, ok
}
