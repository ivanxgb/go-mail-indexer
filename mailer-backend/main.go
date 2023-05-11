package main

import (
	"fmt"
	"mailer-backend/models"
	"mailer-backend/z_adapter"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Enforcing the Content-Type header to be application/json or else the request will fail
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// endpoint to search for emails
	// will receive a json with the following structure:
	// {
	// 	"search": string
	// }

	r.Post("/search", func(w http.ResponseWriter, r *http.Request) {

		// If the request body is empty
		if r.Body == nil {
			http.Error(w, "Please send a request body", http.StatusBadRequest)
			return
		}

		// If the request body is not empty but the json is not valid, it will return an error

		var search models.Search

		err := search.FromJson(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Searching")

		// search in the database
		resp, err := z_adapter.SearchInMails(search.Search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Found")

		http.Header.Add(w.Header(), "content-type", "application/json")
		w.Write(resp)
	})

	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello World!"))
	//})

	http.ListenAndServe(":2507", r)

}
