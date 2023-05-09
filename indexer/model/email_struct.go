package model

type Email struct {
	MessageID string   `json:"messageID"`
	Date      string   `json:"date"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Subject   string   `json:"subject"`
	CC        []string `json:"cc"`
	BCC       []string `json:"bcc"`
	XFrom     string   `json:"xFrom"`
	XTo       []string `json:"xTo"`
	XCC       []string `json:"xCC"`
	XBCC      []string `json:"xBCC"`
	XFolder   string   `json:"xFolder"`
	XFileName string   `json:"xFileName"`
	Content   string   `json:"content"`
}
