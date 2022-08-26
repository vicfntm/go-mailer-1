package models

type FeedbackForm struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	MessageText string `json:"message"`
	MessageType string `json:"message_type"`
}
