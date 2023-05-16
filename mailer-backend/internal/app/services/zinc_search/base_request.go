package zinc_search

import (
	"bytes"
	"fmt"
	"io"
	"mailer-backend/env_loader"
	"net/http"
)

func baseReq(json []byte) ([]byte, error) {
	api, user, pass := env_loader.GetEnvData()

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(json))

	if err != nil {
		fmt.Println("There was an error creating the request")
		return nil, err
	}

	// Setting headers and auth
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)

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