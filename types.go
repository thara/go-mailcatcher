package main

type Inbox struct {
	Messages []*Message
}

type Message struct {
	From string
	Rcpt string
	Data string
}
