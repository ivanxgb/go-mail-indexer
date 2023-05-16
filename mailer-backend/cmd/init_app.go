package main

import (
	"flag"
	"fmt"
	"log"
	el "mailer-backend/env_loader"
	"mailer-backend/internal/app/api"
	"net/http"
	"os/exec"
)

const (
	// port where zinc search runs
	zSearchPort = ":4080"
)

var (
	// port where the server will run, 1507 by default
	port = flag.String("port", "1507", "Port to run the server on")
)

func initApp() {
	flag.Parse()
	el.LoadEnv()
	port := fmt.Sprintf(":%s", *port)

	go startZSearcher()
	printLink(zSearchPort, "ZSearch is running on")
	router := api.GetRouter()

	printLink(port, "Server is running on")

	err := http.ListenAndServe(port, router)

	if err != nil {
		log.Fatal(err)
	}
}

// startZSearcher runs the shell script that starts the zinc search server
func startZSearcher() {
	cmd := exec.Command("sh", "./z_launcher.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// printLink prints the url where the server is running as a clickable link
func printLink(port string, cliText string) {
	url := fmt.Sprintf("http://localhost%s", port)
	displayText := fmt.Sprintf("%s %s", cliText, url)

	link := fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, displayText)
	fmt.Println(link)
}
