package server

import (
	"fmt"
	"time"

	"github.com/fzzy/sockjs-go/sockjs"

	"github.com/Vladimiroff/mazungumzo/workq"
)

type Client struct {
	session        sockjs.Session
	queue          workq.Queue
	Name           string
	NativeLanguage string
}

var clients = NewClientPool()

func (c *Client) Send(sender *Client, message []byte) {
	translatable := new(workq.Item)
	*translatable = workq.Item{
		Sender:  sender.Name,
		Message: string(message),
		Src:     sender.NativeLanguage,
		Dest:    c.NativeLanguage,
		Done:    make(chan bool),
	}

	c.queue.Push(translatable)
	go c.stream()
}

func (c *Client) stream() {
	for !c.queue.IsEmpty() {
		message := c.queue.Pop()
		c.session.Send([]byte(fmt.Sprintf("[%v]%s: %s",
			time.Now().Format("15:04:05"),
			message.Sender,
			message.Translated,
		)))
	}
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
