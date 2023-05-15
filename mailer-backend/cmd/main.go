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
	zSearchPort = ":4080"
)

var (
	port = flag.String("port", "1507", "Port to run the server on")
)

func main() {
	el.LoadEnv()
	flag.Parse()
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

func startZSearcher() {
	cmd := exec.Command("sh", "./z_launcher.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func printLink(port string, cliText string) {
	url := fmt.Sprintf("http://localhost%s", port)
	displayText := fmt.Sprintf("%s %s", cliText, url)

	link := fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, displayText)
	fmt.Println(link)
}
