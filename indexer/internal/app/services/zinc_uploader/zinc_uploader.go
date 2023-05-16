package zinc_uploader

import (
	"bytes"
	"fmt"
	el "indexer/env_loader"
	"indexer/internal/app/model"
	"net/http"
)

// UploadMails receives a slice of model.Email to be sent to the server
func UploadMails(emails *[]model.Email, index string) bool {
	jsonEmails, err := bodyBuilder(emails, index)
	if err != nil {
		fmt.Println("There was an error converting the emails to json")
		return false
	}

	return postData(jsonEmails)
}

// bodyBuilder receives a slice of model.Email and returns a slice of bytes that represents the emails in
// json format accepted by the server
func bodyBuilder(emails *[]model.Email, index string) ([]byte, error) {
	emailBulk := model.BulkV2{
		Index:   index,
		Records: *emails,
	}

	return emailBulk.ToJson()
}

// postData receives a slice of bytes that represents the data to be sent
func postData(data []byte) bool {
	api, user, pass := el.GetEnvData()

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("There was an error creating the request")
		return false
	}

	// Setting headers and auth
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)

	// Creating a new client
	client := &http.Client{}

	// Sending the request via the client
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("There was an error sending the request")
		return false
	}

	// Closing the response body
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Req. Rejected", resp.StatusCode, resp.Status)
		return false
	}

	return true
}
