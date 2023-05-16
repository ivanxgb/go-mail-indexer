package handler

import (
	"fmt"
	oai "mailer-backend/internal/app/models/openai"
	"mailer-backend/internal/app/services/openai"
	"net/http"
)

// SummaryHandler handles requests to the /api/search endpoint, which is used to give a summary of the mail.
func SummaryHandler(w http.ResponseWriter, r *http.Request) {
	// If the request body is empty
	if r.Body == nil {
		http.Error(w, "invalid format", http.StatusBadRequest)
		return
	}

	// If the request body is not empty but the json is not valid, it will return an error
	var summaryReq oai.SummaryReq

	err := summaryReq.FromJson(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send request to openAI
	summaryResp, err := openai.SendOpenAIReq(summaryReq.Content)

	if err != nil {
		fmt.Println("There was an error making the openAI request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the response header as json
	w.Header().Set("Content-Type", "application/json")

	w.Write(summaryResp)
}
