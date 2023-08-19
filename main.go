package main

import (
	"fmt"
	"grooter/internal/configuration"
	"grooter/internal/router"
	"net/http"
)

func main() {
	config, err := configuration.ParseConfigs("./default-conf.yaml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config)
}

func server() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", router.ForwardHandler)

	http.ListenAndServe(":8080", mux)
}
