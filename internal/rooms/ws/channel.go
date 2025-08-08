package ws

import "sync"

type Channel struct {
	name    string
	clients map[*Client]bool
	mu      sync.RWMutex
}

func (c *Channel) AddClient(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.clients[client] = true
}

func (c *Channel) isEmpty() bool {
	return len(c.clients) == 0
}

func (c *Channel) clientExists(client *Client) bool {
	_, exists := c.clients[client]
	return exists
}

func (c *Channel) RemoveClient(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.clientExists(client) {
		return
	}

	delete(c.clients, client)
	close(client.send)

	if c.isEmpty() {
		Hub.RemoveChannel(c.name)
	}
}

func (c *Channel) Broadcast(message []byte) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for client := range c.clients {
		client.send <- message
	}
}

func NewChannel(name string) *Channel {
	return &Channel{
		name:    name,
		clients: make(map[*Client]bool),
	}
}
