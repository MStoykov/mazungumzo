package server

import (
	"github.com/fzzy/sockjs-go/sockjs"
)

type Client struct {
	session        sockjs.Session
	Name           string
	NativeLanguage string
}

var clients = NewClientPool()

func (c *Client) Send(message []byte) {
	c.session.Send(message)
}

func login(s sockjs.Session) *Client {
	name := askForName(s)
	nativeLanguage := askForNativeLanguage(s)
	client := &Client{
		session:        s,
		Name:           name,
		NativeLanguage: nativeLanguage,
	}
	return client
}

func askForName(s sockjs.Session) string {
	s.Send([]byte("What is your name?"))
	return string(s.Receive())
}

func askForNativeLanguage(s sockjs.Session) string {
	s.Send([]byte("What is your native language?"))
	return string(s.Receive())
}
