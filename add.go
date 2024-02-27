package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func addDataToConfig(newData Config) error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	// 新たなデータを追加します
	config.ID = newData.ID
	config.Name = newData.Name
	config.HowTo = newData.HowTo

	// 構造体をJSONにエンコードします
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// エンコードしたJSONデータをファイルに書き込みます
	err = ioutil.WriteFile("config.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
