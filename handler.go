package main

import (
	"add" // addDataToConfig関数が定義されているパッケージ名に置き換えてください
	"fmt"
	"net/http"
)

func RecipeHandler(w http.ResponseWriter, r *http.Request) {
	// 新しいデータを作成します
	newData := add.Config{
		ID:    "2",
		Name:  "New Recipe",
		HowTo: "New HowTo",
	}

	// addDataToConfig関数を呼び出します
	err := add.addDataToConfig(newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスを返します
	fmt.Fprint(w, "Hello, recipes!")
}
