package server

import (
	"sync"

	"github.com/Vladimiroff/mazungumzo/workq"
)

type Language struct {
	mutex   sync.RWMutex
	clients map[string]*Client
	queue   workq.Queue
	Name    string
}

func NewLanguage(name string) *Language {
	l := new(Language)
	l.clients = make(map[string]*Client)
	l.Name = name
	return l
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
