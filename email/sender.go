package email

import (
	"bytes"
	"context"
	htmltemplate "html/template"
	texttemplate "text/template"
)

// Sender is a service which can deliver a message.
type Sender interface {
	Send(ctx context.Context, message *Message) error
}

// Message represents a email.
type Message struct {
	From     string            `json:"from,omitempty"`
	To       []string          `json:"to,omitempty"`
	CC       []string          `json:"cc,omitempty"`
	BCC      []string          `json:"bcc,omitempty"`
	Reply    []string          `json:"reply,omitempty"`
	Subject  string            `json:"subject,omitempty"`
	HTMLBody string            `json:"htmlBody,omitempty"`
	TextBody string            `json:"textBody,omitempty"`
	Data     map[string]string `json:"data,omitempty"`
}

// HTML produces the html output from the HTMLBody template
func (m *Message) HTML() (string, error) {
	body, err := htmltemplate.New("body").Parse(m.HTMLBody)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	err = body.Execute(buf, m.Data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Text produces the text output from TextBody template
func (m *Message) Text() (string, error) {
	body, err := texttemplate.New("body").Parse(m.TextBody)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	err = body.Execute(buf, m.Data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
