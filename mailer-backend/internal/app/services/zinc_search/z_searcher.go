package zinc_search

import (
	"fmt"
	"mailer-backend/internal/app/models"
)

func bodyBuilder(search string) models.ZSearchReq {
	return models.ZSearchReq{
		SearchType: matchAll,
		Query: models.ZQuery{
			Term:  search,
			Field: allFields,
		},
	}
}

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
