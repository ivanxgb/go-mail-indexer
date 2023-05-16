package openai

import models "mailer-backend/internal/app/models/openai"

const (
	api                   = "https://api.openai.com/v1/chat/completions"
	model                 = "gpt-3.5-turbo"
	temperature           = 0.3
	systemRoleDescription = ` You are an assistant for a mail company.
Your job is to people get important and relevant information in mails they receive.
You are a very helpful assistant and you are very good at your job.
You are very good at understanding and extracting useful information from mails.
If the mail contains multiple relevant information, you will extract all of them and show them to the user.

Restrictions:
- You can only use the information in the mail to answer the question.
- You will not respond with any information that is not in the mail.
- You will not respond with any information that is not relevant to the mail context.
`
)

func openAIBodyBuilder(mailContent string) ([]byte, error) {
	roles := []models.OpenAIRoles{
		roleBuilder("system", systemRoleDescription),
		roleBuilder("user", mailContent),
	}

	var body models.OpenAIReq
	body.Model = model
	body.Temperature = temperature
	body.Messages = roles

	return body.ToJson()
}

func roleBuilder(role string, content string) models.OpenAIRoles {
	return models.OpenAIRoles{
		Role:    role,
		Content: content,
	}
}
