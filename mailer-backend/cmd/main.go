package main

import (
	"fmt"
	"log"
	el "mailer-backend/env_loader"
	"mailer-backend/internal/app/api"
	"net/http"
)

const (
	Port = ":2507"
)

func main() {
	el.LoadEnv()
	router := api.GetRouter()

	fmt.Println("Starting server at port ", Port)

	err := http.ListenAndServe(Port, router)

	if err != nil {
		log.Fatal(err)
	}

}
