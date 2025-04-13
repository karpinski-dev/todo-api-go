package main

import "net/http"

func DefineEndpoints(handler *http.ServeMux) {
	handler.HandleFunc("GET /all", GetAllItemsAsJSON)
	handler.HandleFunc("GET /done", GetAllDoneItemsAsJSON)
	handler.HandleFunc("GET /pending", GetAllPendingItemsAsJSON)
}
