package main

import (
	"flag"
	"fmt"
	"log"
	el "mailer-backend/env_loader"
	"mailer-backend/internal/app/api"
	"net/http"
)

var (
	port = flag.String("port", "1507", "Port to run the server on")
)

func main() {
	el.LoadEnv()
	flag.Parse()
	port := fmt.Sprintf(":%s", *port)

	router := api.GetRouter()

	fmt.Println("Server running on port", port)
	err := http.ListenAndServe(port, router)

	if err != nil {
		log.Fatal(err)
	}
}
