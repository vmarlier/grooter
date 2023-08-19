package main

import (
	"net/http"
    "grooter/internal/router"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", router.ForwardHandler)

	http.ListenAndServe(":8080", mux)

}

