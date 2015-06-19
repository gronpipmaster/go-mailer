package mailer

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const eol string = "\r\n"

var (
	//Set custom vars
	BufferSize int = 100
	Debug      bool

	mailQueue chan *Message
)

// Mailer represents mail service.
type Mailer struct {
	host string
	user string
	pass string
}

func New(host, user, pass string) *Mailer {
	return &Mailer{host, user, pass}
}

func (m *Mailer) Listen() {
	mailQueue = make(chan *Message, BufferSize)
	go m.processMailQueue()
}

// Direct Send mail message
func (m *Mailer) Send(msg *Message) (int, error) {
	if Debug {
		log.Println("mailer: Sending mails to:", msg.To)
	}
	host := strings.Split(m.host, ":")

	// get message body
	content := msg.Encode()

	auth := smtp.PlainAuth("", m.user, m.pass, host[0])

	if len(msg.To) == 0 {
		return 0, fmt.Errorf("empty receive emails")
	}

	if len(msg.Body) == 0 {
		return 0, fmt.Errorf("empty email body")
	}

	var num int
	for _, to := range msg.To {
		body := []byte("To: " + to + eol + content)
		err := smtp.SendMail(m.host, auth, msg.From, []string{to}, body)
		if err != nil {
			return num, err
		}
		num++
	}
	return num, nil

}

// Async Send mail message
func (m *Mailer) SendAsync(msg *Message) {
	go func() {
		mailQueue <- msg
	}()
}

func (m *Mailer) processMailQueue() {
	for {
		select {
		case msg := <-mailQueue:
			tos := strings.Join(msg.To, "; ")
			var info string
			index, err := m.Send(msg)
			if err != nil {
				if len(msg.Info) > 0 {
					info = ", info: " + msg.Info
				}
				if Debug {
					log.Println(fmt.Sprintf("mailer: Async sent email %d succeed, not send emails: %s%s err: %s", index, tos, info, err))
				}
			} else if Debug {
				log.Println(fmt.Sprintf("mailer: Async sent email %d succeed, sent emails: %s%s", index, tos, info))
			}
		}
	}
}
