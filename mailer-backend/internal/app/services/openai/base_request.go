package openai

import (
	"bytes"
	"fmt"
	"io"
	ev "mailer-backend/env_loader"
	"net/http"
)

// baseReq is a function that given a openAI request body, sends a request to the openAI API and returns the response
func baseReq(json []byte) ([]byte, error) {
	org, key := ev.GetOpenAIData()

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(json))

	if err != nil {
		fmt.Println("There was an error creating the request")
		return nil, err
	}

	// Setting headers and auth
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("OpenAI-Organization", org)

	// Creating a new client
	client := &http.Client{}

	// Sending the request via the client
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// Closing the response body
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("There was an error sending the request", resp.StatusCode, resp.Status)
		fmt.Println("Response body:", resp.Body)
		return nil, err
	}

	return io.ReadAll(resp.Body)
}
