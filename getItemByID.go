package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getItemByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 || id > len(items) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	item := items[id-1]
	json.NewEncoder(w).Encode(item)
}
