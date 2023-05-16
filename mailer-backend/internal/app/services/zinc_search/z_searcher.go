package zinc_search

import (
	"fmt"
	"mailer-backend/internal/app/models"
)

// SearchInMails receives a search string and returns the results from the zinc search server.
func SearchInMails(search string) ([]byte, error) {
	body := bodyBuilder(search)
	bodyAsJson, err := body.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := baseReq(bodyAsJson)

	if err != nil {
		fmt.Println("There was an error sending the request")
		return nil, err
	}

	var zResp models.ZSearchResponse
	err = zResp.FromJson(resp)

	if err != nil {
		fmt.Println("There was an error parsing the response")
		return nil, err
	}

	return zResp.Hits.ToSearchResponseJson()
}

// bodyBuilder receives a search string and returns a models.ZSearchReq struct with the search data
// to be sent to the zinc search server.
func bodyBuilder(search string) models.ZSearchReq {
	return models.ZSearchReq{
		SearchType: match,
		Query: models.ZQuery{
			Term:  search,
			Field: allFields,
		},
	}
}
