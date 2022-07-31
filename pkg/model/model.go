package model

type Event struct{}

type ClipboardUpdated struct {
	Event
	Source      string `json:"source"`
	Content     []byte `json:"content"`
	ContentType string `json:"contentType"`
}
