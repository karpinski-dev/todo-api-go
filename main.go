package main

import (
	"net/http"
)

func main() {
	router := http.NewServeMux()
	DefineEndpoints(router)

	apiServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	apiServer.ListenAndServe()
}
