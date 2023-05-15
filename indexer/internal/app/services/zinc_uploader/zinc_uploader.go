package zinc_uploader

import (
	"bytes"
	"fmt"
	el "indexer/env_loader"
	de "indexer/internal/app/dir_explorer"
	"indexer/internal/app/model"
	"net/http"
	"sync"
)

const (
	EmailBatchSize = 1000
	GoRoutines     = 4
)

// ProcessFiles receives a slice of strings that represent the path of the emails to be sent to the server
func ProcessFiles(filePaths *[]string) {
	files := *filePaths
	totalFiles := len(files)
	filesSlicesSize := totalFiles / GoRoutines
	var fileErrorCounter, errorPostCounter, indexSliceStart int
	fmt.Println("Total files to be processed: ", totalFiles)

	var wg sync.WaitGroup
	wg.Add(GoRoutines)
	filesErrorChan := make(chan int, GoRoutines)
	failMailsUpload := make(chan int, GoRoutines)

	for i := 0; i < GoRoutines; i++ {
		indexSliceEnd := indexSliceStart + filesSlicesSize

		// handle the remainder
		if i == GoRoutines-1 {
			indexSliceEnd = totalFiles
		}

		filesPathSlice := files[indexSliceStart:indexSliceEnd]
		indexSliceStart = indexSliceEnd

		go processFiles(filesPathSlice, &wg, filesErrorChan, failMailsUpload)
	}

	wg.Wait()
	close(filesErrorChan)
	close(failMailsUpload)

	for err := range filesErrorChan {
		fileErrorCounter += err
	}

	for err := range failMailsUpload {
		errorPostCounter += err
	}

	fmt.Println("Total files with errors: ", fileErrorCounter)
	fmt.Println("Total files don't uploaded: ", errorPostCounter)
	fmt.Println("Total files uploaded: ", totalFiles-fileErrorCounter-errorPostCounter)
}

func processFiles(filePaths []string, wg *sync.WaitGroup, errorChan chan int, failMailsUpload chan int) {
	defer wg.Done()
	var errorFilesCounter, errorPostCounter int
	var emails []model.Email

	for _, filePath := range filePaths {
		email, err := de.EmailConverter(filePath)
		if err != nil {
			errorFilesCounter++
			continue
		}

		// append the valid emails to the emails slice
		emails = append(emails, email)

		// if the emails slice is full, send it to the server
		if len(emails) == EmailBatchSize {
			if !processEmails(&emails) {
				errorPostCounter += EmailBatchSize
			}
			emails = nil
		}
	}

	// send the remaining emails to the server
	if len(emails) > 0 {
		if !processEmails(&emails) {
			errorPostCounter += len(emails)
		}
	}

	failMailsUpload <- errorPostCounter
	errorChan <- errorFilesCounter
}

// processEmails receives a slice of model.Email and converts it to json to be sent to the server
func processEmails(emails *[]model.Email) bool {
	jsonEmails, err := de.EmailsToBulkJson(emails)
	if err != nil {
		fmt.Println("There was an error converting the emails to json")
		return false
	}

	return postData(jsonEmails)
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
