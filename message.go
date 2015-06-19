package mailer

import "strings"

type Message struct {
	To      []string
	From    string
	Subject string
	Body    string
	User    string
	Type    string
	Info    string
}

// Create html mail message
func NewHtmlMessage(to []string, from, userFrom, subject, body string) *Message {
	msg := NewMessage(to, from, userFrom, subject, body)
	msg.Type = "html"
	return msg
}

func NewMessage(to []string, from, userFrom, subject, body string) *Message {
	return &Message{
		To:      to,
		From:    from,
		Subject: subject,
		Body:    body,
		User:    userFrom,
	}
}

// Encode mail content body
func (m Message) Encode() string {
	// create mail content
	contents := make([]string, 0)
	contents = append(contents, "MIME-Version: 1.0")
	contents = append(contents, "From: "+m.User+"<"+m.From+">")
	contents = append(contents, "Subject: "+m.Subject)
	// set mail type
	contentType := "text/plain; charset=UTF-8"
	if m.Type == "html" {
		contentType = "text/html; charset=UTF-8"
	}
	contents = append(contents, "Content-Type: "+contentType+eol)
	contents = append(contents, m.Body)

	return strings.Join(contents, eol)
}
