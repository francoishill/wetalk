package mailer

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/astaxie/beego"
)

var (
	AuthUser, AuthPass, MailHost string
)

type Message struct {
	To      []string
	From    string
	Subject string
	Body    string
	User    string
	Type    string
	Massive bool
	Info    string
}

// create mail content
func (m Message) Content() string {
	// set mail type
	contentType := "text/plain; charset=UTF-8"
	if m.Type == "html" {
		contentType = "text/html; charset=UTF-8"
	}

	// create mail content
	content := "From: " + m.User + "<" + m.From +
		">\r\nSubject: " + m.Subject + "\r\nContent-Type: " + contentType + "\r\n\r\n" + m.Body
	return content
}

// Direct Send mail message
func Send(msg Message) (int, error) {
	host := strings.Split(MailHost, ":")

	// get message body
	content := msg.Content()

	auth := smtp.PlainAuth("", AuthUser, AuthPass, host[0])

	if len(msg.To) == 0 {
		return 0, fmt.Errorf("empty receive emails")
	}

	if msg.Massive {
		// send mail to multiple emails one by one
		num := 0
		for _, to := range msg.To {
			body := []byte("To: " + to + "\r\n" + content)
			err := smtp.SendMail(MailHost, auth, msg.From, []string{to}, body)
			if err != nil {
				return num, err
			}
			num++
		}
		return num, nil
	} else {
		body := []byte("To: " + strings.Join(msg.To, ";") + "\r\n" + content)

		// send to multiple emails in one message
		err := smtp.SendMail(MailHost, auth, msg.From, msg.To, body)
		if err != nil {
			return 0, err
		} else {
			return 1, nil
		}
	}
}

// Async Send mail message
func SendAsync(msg Message) {
	// TODO may be need pools limit concurrent nums
	go func() {
		if num, err := Send(msg); err != nil {
			tos := strings.Join(msg.To, "; ")
			info := ""
			if len(msg.Info) > 0 {
				info = ", info: " + msg.Info
			}
			// log failed
			beego.Error(fmt.Sprintf("Async send email %d succeed, not send emails: %s%s", num, tos, info))
		}
	}()
}

// Create html mail message
func NewHtmlMessage(To []string, From, Subject, Body string) Message {
	return Message{
		To:      To,
		From:    From,
		Subject: Subject,
		Body:    Body,
		Type:    "html",
	}
}
