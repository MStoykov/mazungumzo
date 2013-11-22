package server

import (
	"errors"
	"sync"
)

type ClientPool struct {
	mutex   sync.RWMutex
	clients map[string]*Client
}

func NewClientPool() *ClientPool {
	p := new(ClientPool)
	p.clients = make(map[string]*Client)
	return p
}

func (cp *ClientPool) Add(c *Client) error {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	if _, ok := cp.clients[c.Name]; ok {
		return errors.New("Client with this name already exists")
	}

	cp.clients[c.Name] = c
	return nil
}

func (cp *ClientPool) Remove(c *Client) {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	delete(cp.clients, c.Name)
}

func (cp *ClientPool) Broadcast(sender *Client, m []byte) {
	cp.mutex.RLock()
	defer cp.mutex.RUnlock()

	for _, client := range cp.clients {
		client.Send(sender, m)
	}
}
