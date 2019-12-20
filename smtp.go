package main

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/emersion/go-smtp"
)

// The Backend implements SMTP server methods.
type SmtpBackend struct {
	inbox *Inbox
}

// Login handles a login command with username and password.
func (b *SmtpBackend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	return &smtpSession{b.inbox, nil}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (b *SmtpBackend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

type smtpSession struct {
	inbox   *Inbox
	current *Message
}

func (s *smtpSession) currentMessage() *Message {
	if s.current == nil {
		msg := Message{}
		s.current = &msg
		s.inbox.Messages = append(s.inbox.Messages, s.current)
	}
	return s.current
}

func (s *smtpSession) Mail(from string, opts smtp.MailOptions) error {
	log.Println("Mail from:", from)
	s.currentMessage().From = from
	return nil
}

func (s *smtpSession) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	s.currentMessage().Rcpt = to
	return nil
}

func (s *smtpSession) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		data := string(b)
		s.currentMessage().Data = data
		log.Println("Data:", data)
		log.Println("Inbox length:", len(s.inbox.Messages))
	}
	return nil
}

func (s *smtpSession) Reset() {}

func (s *smtpSession) Logout() error {
	return nil
}
