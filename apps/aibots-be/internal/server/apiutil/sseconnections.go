package apiutil

import (
	"context"
	"net/http"
	"sync"
)

type Connection struct {
	writer http.ResponseWriter
	Closed chan struct{}
	ctx    context.Context
	cancel context.CancelFunc
}

func NewConnection(ctx context.Context, w http.ResponseWriter) *Connection {
	// func NewConnection(ctx context.Context) *Connection {
	ctx, cancel := context.WithCancel(ctx)
	return &Connection{
		writer: w,
		Closed: make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (c *Connection) Close() {
	c.cancel()
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

// Add adds a new connection to the set.
func (s *SSEConnections) Add(c *Connection) {
	s.Lock()
	s.connections[c] = struct{}{}
	s.Unlock()
}

// Remove removes a connection from the set.
func (s *SSEConnections) Remove(c *Connection) {
	s.Lock()
	c.cancel()
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
