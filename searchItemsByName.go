package main

import (
	"encoding/json"
	"net/http"
)

func searchItemsByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var result []Item

	for _, item := range items {
		if item.Name == name {
			result = append(result, item)
		}
	}
	json.NewEncoder(w).Encode(result)
}
