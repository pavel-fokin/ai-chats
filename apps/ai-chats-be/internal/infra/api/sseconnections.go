package api

import (
	"sync"
)

type Connection struct {
	Closed chan struct{}
}

func (c *Connection) Close() {
	close(c.Closed)
}

// SSEConnections is a set of connections for Server-Sent Events.
type SSEConnections struct {
	sync.RWMutex
	connections map[*Connection]struct{}
}

// NewSSEConnections creates a new set of connections.
func NewSSEConnections() *SSEConnections {
	return &SSEConnections{
		connections: make(map[*Connection]struct{}),
	}
}

// NewConnection creates a new connection and adds it to the set.
func (s *SSEConnections) AddConnection() *Connection {
	c := &Connection{
		Closed: make(chan struct{}),
	}
	s.Lock()
	s.connections[c] = struct{}{}
	s.Unlock()
	return c
}

// Add adds a new connection to the set.
func (s *SSEConnections) Add(c *Connection) {
	s.Lock()
	s.connections[c] = struct{}{}
	s.Unlock()
}

// Remove removes a connection from the set.
func (s *SSEConnections) Remove(c *Connection) {
	s.Lock()
	delete(s.connections, c)
	s.Unlock()
}

// CloseAll closes all connections in the set.
func (s *SSEConnections) CloseAll() {
	s.Lock()
	for c := range s.connections {
		c.Close()
	}
	s.Unlock()
}
