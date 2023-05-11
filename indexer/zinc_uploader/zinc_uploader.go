package zinc_uploader

import (
	"bytes"
	"fmt"
	de "indexer/dir_explorer"
	el "indexer/env_loader"
	"indexer/model"
	"net/http"
	"sync"
)

const (
	EmailBatchSize = 1000
	GoRoutines     = 5
)

// SendFilesToServer receives a slice of strings that represent the path of the emails to be sent to the server
func SendFilesToServer(filesPath []string) {
	totalFiles := len(filesPath)
	var errorCounter int

	var wg sync.WaitGroup
	filesErrorChan := make(chan int, GoRoutines)
	wg.Add(GoRoutines)

	fmt.Println("Total files to be processed: ", totalFiles)

	// divide the files in slices to be processed by the goroutines
	filesSlicesSize := totalFiles / GoRoutines
	indexSliceStart := 0
	for i := 0; i < GoRoutines; i++ {
		indexSliceEnd := indexSliceStart + filesSlicesSize

		// handle the remainder
		if i == GoRoutines-1 {
			indexSliceEnd = totalFiles
		}
		filesPathSlice := filesPath[indexSliceStart:indexSliceEnd]
		indexSliceStart = indexSliceEnd

		go ProcessFiles(filesPathSlice, &wg, filesErrorChan)
	}

	wg.Wait()
	close(filesErrorChan)
	for filesError := range filesErrorChan {
		errorCounter += filesError
	}

	fmt.Println("Total files with errors: ", errorCounter)
	fmt.Println("Total files sent to the server: ", totalFiles-errorCounter)
}

func ProcessFiles(filesPath []string, wg *sync.WaitGroup, errorChan chan int) {
	defer wg.Done()
	var errorCounter int
	var emails []model.Email

	for _, filePath := range filesPath {
		email, err := de.EmailConverter(filePath)
		if err != nil {
			errorCounter++
			continue
		}

		// append the valid emails to the emails slice
		emails = append(emails, email)

		// if the emails slice is full, send it to the server
		if len(emails) == EmailBatchSize {
			PostEmails(emails)
			emails = nil
		}
	}

	// send the remaining emails to the server
	if len(emails) > 0 {
		PostEmails(emails)
	}

	errorChan <- errorCounter
}

// PostEmails receives a slice of model.Email and converts it to json to be sent to the server
func PostEmails(emails []model.Email) {
	emailAsJson, err := de.EmailsToBulkJson(emails)
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
		fmt.Println("Req. Rejected", resp.StatusCode, resp.Status)
		return false
	}

	return true
}
