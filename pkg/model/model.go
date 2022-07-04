package model

type Event struct{}

type ClipboardUpdated struct {
	Event
	Content     []byte `json:"content"`
	ContentType string `json:"contentType"`
}
