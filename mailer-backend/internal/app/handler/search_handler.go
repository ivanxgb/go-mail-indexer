package handler

import (
	"fmt"
	"mailer-backend/internal/app/models"
	zs "mailer-backend/internal/app/services/zinc_search"
	"net/http"
)

// SearchHandler handles api requests to the search /api/route and returns the
// results from the zinc search server.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// If the request body is empty
	if r.Body == nil {
		http.Error(w, "invalid format", http.StatusBadRequest)
		return
	}

	// If the request body is not empty but the json is not valid, it will return an error
	var search models.Search

	err := search.FromJson(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// search in zinc
	resp, err := zs.SearchInMails(search.Search)

	if err != nil {
		fmt.Println("Error searching in zinc")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the response header as json
	w.Header().Set("Content-Type", "application/json")

	w.Write(resp)
}
