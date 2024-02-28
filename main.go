package main

import (
	"fmt"
	"net/http"
)

func main() {
	// サーバーを起動する
	fmt.Println("Server is running on port 8080")

	// ルーティングとハンドラーの登録
	http.HandleFunc("/items", listItems)
	http.HandleFunc("/item/", getItemByID)
	http.HandleFunc("/search", searchItemsByName)

	http.ListenAndServe(":8080", nil)
}
