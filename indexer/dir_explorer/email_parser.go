package dir_explorer

import (
	"encoding/json"
	"fmt"
	"indexer/model"
	"indexer/utils"
	"regexp"
	"strings"
)

// EmailConverter receives a string that represents the path of the email to be read
// and returns a model.Email struct with the email data
func EmailConverter(filePath string) (model.Email, error) {
	content, err := utils.FileReader(filePath)

	if err != nil {
		fmt.Println("There was an error reading the file: " + filePath)
		return model.Email{}, err
	}

	mail := parseMail(string(content))

	return mail, nil
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
func parseMail(content string) model.Email {
	mail := model.Email{}

	// re is a regular expression that matches the header fields of the email
	// it matches the key and the value of the field to be able to extract the headers
	re := regexp.MustCompile(`^(.*?):\s*(.*)$`)

	// headerLines is a slice of strings that represents the header of the email
	headerLines := strings.Split(content, "\n\n")[0]

	headerFields := strings.Split(headerLines, "\n")
	for _, field := range headerFields {
		match := re.FindStringSubmatch(field)
		if len(match) == 3 {
			key := match[1]
			value := match[2]

			switch strings.ToLower(key) {
			case "message-id":
				mail.MessageID = value
			case "date":
				mail.Date = value
			case "from":
				mail.From = value
			case "to":
				mail.To = strings.Split(value, ",")
			case "subject":
				mail.Subject = value
			case "cc":
				mail.CC = strings.Split(value, ",")
			case "bcc":
				mail.BCC = strings.Split(value, ",")
			case "x-from":
				mail.XFrom = value
			case "x-to":
				mail.XTo = strings.Split(value, ",")
			case "x-cc":
				mail.XCC = strings.Split(value, ",")
			case "x-bcc":
				mail.XBCC = strings.Split(value, ",")
			case "x-folder":
				mail.XFolder = value
			case "x-filename":
				mail.XFileName = value
			}
		}
	}

	// Extracting content
	contentLines := strings.Split(content, "\n")[1:]
	mail.Content = strings.Join(contentLines, "\n")

	return mail
}
