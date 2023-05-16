package OpenAIModel

import (
	"encoding/json"
	"fmt"
	"io"
)

type SummaryReq struct {
	Content string `json:"content"`
}

type OpenAIResp struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Choices []OChoice `json:"choices"`
	Usage   OUsage    `json:"usage"`
}

type OChoice struct {
	Index   int         `json:"index"`
	Message OpenAIRoles `json:"message"`
	Finish  string      `json:"finish_reason"`
}

type OUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func (os *SummaryReq) FromJson(data io.ReadCloser) error {
	decoder := json.NewDecoder(data)
	err := decoder.Decode(&os)
	if err != nil && err != io.EOF {
		return err
	}

	if os.Content == "" {
		return fmt.Errorf("incorrect json format")
	}

	return nil
}

func (os *SummaryReq) ToJson() ([]byte, error) {
	return json.Marshal(os)
}

func (o *OpenAIResp) FromJson(data []byte) error {
	return json.Unmarshal(data, o)
}
