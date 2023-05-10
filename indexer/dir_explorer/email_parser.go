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

// EmailConverter receives a string that represents the path of the email to be read
// and returns a model.Email struct with the email data
func EmailConverter(filePath string) (model.Email, error) {
	content, err := utils.FileReader(filePath)

	if err != nil {
		return model.Email{}, err
	}

	emailParsed, err := parseMail(string(content))

	if err != nil {
		return model.Email{}, err
	}

	return emailParsed, nil
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

	newEmail := model.Email{
		MessageID: emailParsed.Header.Get("Message-Id"),
		Date:      emailParsed.Header.Get("Date"),
		From:      emailParsed.Header.Get("From"),
		To:        parseHeaderList(emailParsed.Header.Get("To")),
		Subject:   emailParsed.Header.Get("Subject"),
		CC:        parseHeaderList(emailParsed.Header.Get("Cc")),
		BCC:       parseHeaderList(emailParsed.Header.Get("Bcc")),
		XFrom:     emailParsed.Header.Get("X-From"),
		XTo:       parseHeaderList(emailParsed.Header.Get("X-To")),
		XCC:       parseHeaderList(emailParsed.Header.Get("X-Cc")),
		XBCC:      parseHeaderList(emailParsed.Header.Get("X-Bcc")),
		XFolder:   emailParsed.Header.Get("X-Folder"),
		XFileName: emailParsed.Header.Get("X-Filename"),
	}

	// Body from Reader to String
	emailContent, err := io.ReadAll(emailParsed.Body)
	if err == nil {
		newEmail.Content = string(emailContent)
	}
	return newEmail, nil
}

func parseHeaderList(header string) []string {
	// check if header is empty
	if header == "" {
		return nil
	}

	return strings.Split(header, ", ")
}

//Converting email: /Users/ivanxgb/Desktop/mails/maildir/baughman-d/calendar/19.
