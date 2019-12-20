package main

import (
	"log"
	"net/http"
	"time"

	"github.com/emersion/go-smtp"
)

func main() {
	inbox := Inbox{}

	be := &SmtpBackend{&inbox}

	s := smtp.NewServer(be)

	s.Addr = ":1025"
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	go func(inbox *Inbox) {
		addr := ":18080"
		log.Println("Starting http server at", addr)

		h := httpHandlers{inbox}
		http.HandleFunc("/messages", h.listMessages)
		http.ListenAndServe(addr, nil)
	}(&inbox)

	log.Println("Starting mail server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
