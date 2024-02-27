package main

import (
	"fmt"
	"net/http"
)

func RecipeHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンスを返します
	fmt.Fprint(w, "Hello, recipes!")
}
