package z_adapter

import (
	"bytes"
	"fmt"
	"io"
	"mailer-backend/models"
	"net/http"
)

func SearchInMails(search string) ([]byte, error) {
	query := fmt.Sprintf("SELECT * FROM default WHERE from LIKE '%%%s%%'", search)
	body := models.Z_Search{
		Query: models.Query{
			Sql:  query,
			Size: 100,
		},
	}

	json, err := body.ToJson()
	if err != nil {
		return nil, err
	}

	fmt.Println(string(json))

	resp, err := postRequest(json)

	if err != nil {
		return nil, err
	}

	println(string(resp))

	return resp, nil
}

func postRequest(json []byte) ([]byte, error) {
	user, pass := "me@ivan.com", "Z_Challenge"
	api := "http://18.235.128.9:5080/api/default/_search"

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(json))
	if err != nil {
		fmt.Println("There was an error creating the request")
		return nil, err
	}

	// Setting headers and auth
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)

	// Setting useragent
	req.Header.Set("User-Agent", "Z_Challenge")

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

	body, _ := io.ReadAll(resp.Body)
	return body, nil
}
