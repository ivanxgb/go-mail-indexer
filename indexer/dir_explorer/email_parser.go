package dir_explorer

import (
	"encoding/json"
	"fmt"
	"indexer/model"
	"indexer/utils"
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

	return parseMail(content)
}

func EmailsToBulkJson(emails *[]model.Email) ([]byte, error) {
	emailBulk := model.BulkV2{
		Index:   ZincIndex,
		Records: *emails,
	}

	return emailBulk.ToJson()
}

// EmailsToJSON receives a slice of model.Email and returns a slice of bytes that represents the emails in json format
func EmailsToJSON(emails []model.Email) ([]byte, error) {
	emailsAsJson, err := json.Marshal(emails)
	if err != nil {
		fmt.Println("There was an error converting the emails to json")
		return nil, err
	}

	return emailsAsJson, nil
}

// parseMail receives a string that represents the content of the email and returns a model.Email struct with the email data
func parseMail(content string) (model.Email, error) {
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

func emailContentExtractor(email *mail.Message) string {
	emailContent, _ := io.ReadAll(email.Body)
	return string(emailContent)
}

func toList(emails string) []string {
	if emails == "" {
		return []string{}
	}

	return strings.Split(emails, ",")
}

//Converting email: /Users/ivanxgb/Desktop/mails/maildir/baughman-d/calendar/19.
