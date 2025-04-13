package main

import (
	"encoding/json"
	"net/http"
)

var repository DataRepository = &staticRepository{}

func GetAllItemsAsJSON(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(repository.GetAll())
}

func GetAllDoneItemsAsJSON(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(repository.GetDone())
}

func GetAllPendingItemsAsJSON(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(repository.GetPending())
}
