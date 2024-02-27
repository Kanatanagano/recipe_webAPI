package main

import (
	"encoding/json"
	"os"
)

func removeDataFromConfig(id string) error {
	// ファイルを開きます
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// JSONをデコードします
	var config []Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	// 削除したいデータを探します
	for i, item := range config {
		if item.ID == id {
			// データを削除します
			config = append(config[:i], config[i+1:]...)
			break
		}
	}

	// データを再度エンコードします
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// データをファイルに書き込みます
	err = os.WriteFile("config/config.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
