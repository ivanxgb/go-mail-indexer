package handler

import (
	"indexer/internal/app/model"
	"indexer/internal/app/utils"
	"io"
	"net/mail"
	"strings"
)

const (
	ZincIndex = "en_mails"
)

// EmailConverter receives a string that represents the path of the email to be read
// and returns a model.Email struct with the email data
func EmailConverter(filePath string) (model.Email, error) {
	content, err := utils.FileReader(filePath)

	if err != nil {
		return model.Email{}, err
	}

	return mailToStruct(content)
}

// mailToStruct receives a string that represents the content of the mail fail and returns a model.Email struct with the file data
func mailToStruct(content string) (model.Email, error) {
	emailAsReader := strings.NewReader(content)
	emailParsed, err := mail.ReadMessage(emailAsReader)

	if err != nil {
		return model.Email{}, err
	}

	email := model.Email{
		MessageID: emailParsed.Header.Get("Message-ID"),
		Date:      emailParsed.Header.Get("Date"),
		From:      emailParsed.Header.Get("From"),
		To:        toList(emailParsed.Header.Get("To")),
		Subject:   emailParsed.Header.Get("Subject"),
		CC:        toList(emailParsed.Header.Get("Cc")),
		BCC:       toList(emailParsed.Header.Get("Bcc")),
		XFrom:     emailParsed.Header.Get("X-From"),
		XTo:       toList(emailParsed.Header.Get("X-To")),
		XCC:       toList(emailParsed.Header.Get("X-Cc")),
		XBCC:      toList(emailParsed.Header.Get("X-Bcc")),
		XFolder:   emailParsed.Header.Get("X-Folder"),
		XFileName: emailParsed.Header.Get("X-Filename"),
		Content:   emailContentExtractor(emailParsed),
	}

	return email, nil
}

// emailContentExtractor receives a mail.Message and returns the content of the email as a string
func emailContentExtractor(email *mail.Message) string {
	emailContent, _ := io.ReadAll(email.Body)
	return string(emailContent)
}

// toList receives a string that represents a list of email addresses separated by comma and returns a slice of strings
func toList(emails string) []string {
	if emails == "" {
		return []string{}
	}

	return strings.Split(emails, ",")
}
