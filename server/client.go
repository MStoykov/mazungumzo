package server

import (
	"fmt"

	"github.com/fzzy/sockjs-go/sockjs"
)

type Client struct {
	session        sockjs.Session
	Name           string
	NativeLanguage string
}

var clients = make(map[string]*Client)

func (c *Client) Send(message []byte) {
	c.session.Send(message)
}

func login(s sockjs.Session) *Client {
	name := askForName(s)
	if _, ok := clients[name]; ok {
		s.Send([]byte("Try again..."))
		return login(s)
	}

	nativeLanguage := askForNativeLanguage(s)
	client := &Client{
		session:        s,
		Name:           name,
		NativeLanguage: nativeLanguage,
	}
	clients[name] = client
	client.Send([]byte(fmt.Sprintf("Welcome %s", name)))

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
