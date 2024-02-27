package main

import (
	"encoding/json"
	"os"
)

func searchInConfig(key string) (string, error) {
	file, err := os.Open("config/config.json")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return "", err
	}

	switch key {
	case "id":
		return config.ID, nil
	case "name":
		return config.Name, nil
	case "howto":
		return config.HowTo, nil
	default:
		return "", nil
	}
}
