package openai

import (
	"fmt"
	oai "mailer-backend/internal/app/models/openai"
)

// SendOpenAIReq sends a request to the openAI API and returns the response.
func SendOpenAIReq(mailContent string) ([]byte, error) {
	body, err := openAIBodyBuilder(mailContent)
	if err != nil {
		fmt.Println("There was an error building the body")
		return nil, err
	}

	resp, err := baseReq(body)
	if err != nil {
		fmt.Println("There was an error making the openAI request")
		return nil, err
	}

	var openAIResp oai.OpenAIResp
	err = openAIResp.FromJson(resp)

	if err != nil {
		fmt.Println("There was an error parsing the response")
		return nil, err
	}

	var summaryResp oai.SummaryReq
	summaryResp.Content = openAIResp.Choices[0].Message.Content

	return summaryResp.ToJson()
}
