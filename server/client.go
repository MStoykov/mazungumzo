package server

import (
	"github.com/fzzy/sockjs-go/sockjs"

	"github.com/Vladimiroff/mazungumzo/workq"
)

type Client struct {
	session sockjs.Session
	queue   workq.Queue
	Name    string
}

var (
	clients   = NewClientPool()
	languages = NewLanguagePool()
)

func (c *Client) Send(message string) {
	c.session.Send([]byte(message))
}

func login(s sockjs.Session) (*Client, string) {
	name := askForName(s)
	nativeLanguage := askForNativeLanguage(s)
	client := &Client{
		session: s,
		Name:    name,
	}

	return client, nativeLanguage
}

func askForName(s sockjs.Session) string {
	s.Send([]byte("What is your name?"))
	return string(s.Receive())
}

func askForNativeLanguage(s sockjs.Session) string {
	s.Send([]byte("What is your native language?"))
	return string(s.Receive())
}
