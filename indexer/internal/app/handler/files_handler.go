package handler

import (
	"fmt"
	"indexer/internal/app/model"
	zu "indexer/internal/app/services/zinc_uploader"
	"sync"
)

const (
	EmailBatchSize = 1000
	GoRoutines     = 4
)

// ProcessFiles receives a slice of strings that represent the path of the emails to be sent to the server
// and the index where the emails will be uploaded. Then it processes the files and sends the emails to the server in batches.
func ProcessFiles(filePaths *[]string, index string) {
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

		go processFiles(filesPathSlice, index, &wg, filesErrorChan, failMailsUpload)
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

// processFiles receives a slice of strings that represent the path of the emails to be sent to the server
// it reads the files, converts them to emails and sends them to the server in batches.
func processFiles(filePaths []string, index string, wg *sync.WaitGroup, errorChan chan int, failMailsUpload chan int) {
	defer wg.Done()
	var errorFilesCounter, errorPostCounter int
	var emails []model.Email

	for _, filePath := range filePaths {
		email, err := EmailConverter(filePath)
		if err != nil {
			errorFilesCounter++
			continue
		}

		// append the valid emails to the emails slice
		emails = append(emails, email)

		// if the emails slice is full, send it to the server
		if len(emails) == EmailBatchSize {
			if !zu.UploadMails(&emails, index) {
				errorPostCounter += EmailBatchSize
			}
			emails = nil
		}
	}

	// send the remaining emails to the server
	if len(emails) > 0 {
		if !zu.UploadMails(&emails, index) {
			errorPostCounter += len(emails)
		}
	}

	failMailsUpload <- errorPostCounter
	errorChan <- errorFilesCounter
}
