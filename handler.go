package main

import (
	"encoding/json"
	"net/http"
)

func getRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// レシピのデータ
	recipe := map[string]string{
		"id":    "1",
		"name":  "チキンカレー",
		"howto": "xxxxx",
	}

	// レスポンスの生成
	response, err := json.Marshal(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスの書き込み
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
