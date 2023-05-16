package OpenAIModel

import "encoding/json"

type OpenAIReq struct {
	Model       string        `json:"model"`
	Messages    []OpenAIRoles `json:"messages"`
	Temperature float64       `json:"temperature"`
}

type OpenAIRoles struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (o *OpenAIReq) ToJson() ([]byte, error) {
	return json.Marshal(o)
}
