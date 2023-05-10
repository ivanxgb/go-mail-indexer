package zinc_uploader

import (
	"bytes"
	"fmt"
	de "indexer/dir_explorer"
	el "indexer/env_loader"
	"indexer/model"
	"net/http"
)

const (
	EmailBatchSize = 1000
)

// SendFilesToServer receives a slice of strings that represent the path of the emails to be sent to the server
func SendFilesToServer(filesPath []string) {
	fmt.Println("Total files found: ", len(filesPath))
	var emailJsonArray []model.Email

	var mailConverted = 0
	var mailsWithError = 0

	for _, filePath := range filesPath {
		email, err := de.EmailConverter(filePath)
		if err != nil {
			mailsWithError++
			continue
		}

		mailConverted++
		emailJsonArray = append(emailJsonArray, email)

		if len(emailJsonArray) == EmailBatchSize {
			PostEmails(emailJsonArray)
			emailJsonArray = nil
		}
	}

	fmt.Println("Total mails sent: ", mailConverted)
	fmt.Println("Total mails with error or incorrect format: ", mailsWithError)

}

// PostEmails receives a slice of model.Email and converts it to json to be sent to the server
func PostEmails(emails []model.Email) {
	emailAsJson, err := de.EmailsToJSON(emails)
	if err != nil {
		fmt.Println("There was an error converting the emails to json")
		return
	}

	PostFile(emailAsJson)
}

// PostFile receives a slice of bytes that represents the emails in json format and sends it to the server
func PostFile(emails []byte) bool {
	api, user, pass := el.GetEnvData()

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(emails))
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
		return false
	}

	return true
}
