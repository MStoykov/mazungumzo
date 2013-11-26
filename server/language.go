package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/Vladimiroff/mazungumzo/workq"
)

type Language struct {
	mutex   sync.RWMutex
	clients map[string]*Client
	queue   *workq.Queue
	Name    string
}

func NewLanguage(name string) *Language {
	l := new(Language)
	l.clients = make(map[string]*Client)
	l.queue = workq.NewQueue()
	l.Name = name
	go l.Stream()
	return l
}

func (l *Language) Stream() {
	for item := range l.queue.Pop() {
		for _, client := range l.clients {
			message := fmt.Sprintf("[%v] %s: %s",
				item.Time.Format("15:04:05"),
				item.Sender,
				item.Translated,
			)
			client.Send(message)
		}
	}
}

func (l *Language) AddClient(c *Client) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.clients[c.Name] = c
}

func (l *Language) RemoveClient(c *Client) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	delete(l.clients, c.Name)
}

func (l *Language) Send(sender *Client, timeSent time.Time, message []byte) {
	translatable := new(workq.Item)
	*translatable = workq.Item{
		Time:    timeSent,
		Sender:  sender.Name,
		Message: string(message),
		Src:     sender.Language,
		Dest:    l.Name,
		Done:    make(chan bool),
	}

	l.queue.Push(translatable)
}
