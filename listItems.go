package main

import (
	"encoding/json"
	"net/http"
)

func listItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}
