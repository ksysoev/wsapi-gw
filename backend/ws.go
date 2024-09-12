package backend

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/ksysoev/wasabi"
	"golang.org/x/sync/singleflight"
	"nhooyr.io/websocket"
)

type WSBackend struct {
	group       *singleflight.Group
	connections map[string]*websocket.Conn
	lock        *sync.RWMutex
	factory     WSRequestFactory
	URL         string
}

type WSRequestFactory func(r wasabi.Request) (websocket.MessageType, []byte, error)

// NewWSBackend creates a new instance of WSBackend with the specified URL.
func NewWSBackend(url string, factory WSRequestFactory) *WSBackend {
	return &WSBackend{
		group:       &singleflight.Group{},
		connections: make(map[string]*websocket.Conn),
		lock:        &sync.RWMutex{},
		factory:     factory,
		URL:         url,
	}
}

// Handle handles the incoming request from the WebSocket connection.
// It writes the request data to the WebSocket connection's context.
// The function returns an error if there is any issue with the connection or writing the data.
func (b *WSBackend) Handle(conn wasabi.Connection, r wasabi.Request) error {
	c, err := b.getConnection(conn)

	if err != nil {
		return err
	}

	msgType, data, err := b.factory(r)

	if err != nil {
		return err
	}

	return c.Write(r.Context(), msgType, data)
}

// getConnection returns the websocket connection associated with the given connection.
// If the connection is already established, it returns the existing connection.
// Otherwise, it establishes a new connection and returns it.
func (b *WSBackend) getConnection(conn wasabi.Connection) (*websocket.Conn, error) {
	b.lock.RLock()
	ws, ok := b.connections[conn.ID()]
	b.lock.RUnlock()

	if ok {
		return ws, nil
	}

	uws, err, _ := b.group.Do(conn.ID(), func() (interface{}, error) {
		fmt.Println("Connecting to", b.URL, "for connection", conn.ID())
		c, resp, err := websocket.Dial(conn.Context(), b.URL, nil)

		if err != nil {
			return nil, err
		}

		if resp.Body != nil {
			defer resp.Body.Close()
		}

		go b.responseHandler(c, conn)

		b.lock.Lock()
		b.connections[conn.ID()] = c
		b.lock.Unlock()

		return c, nil
	})

	if err != nil {
		return nil, err
	}

	ws, ok = uws.(*websocket.Conn)
	if !ok {
		panic("unexpected type")
	}

	return ws, nil
}

// responseHandler handles the response from the server to the client.
// It reads messages from the server, sends them to the client, and manages the connection lifecycle.
func (b *WSBackend) responseHandler(server *websocket.Conn, client wasabi.Connection) {
	var (
		err     error
		msgType websocket.MessageType
		reader  io.Reader
	)

	defer func() {
		code := websocket.StatusNormalClosure
		reason := "connection closed"

		var wsCloseErr websocket.CloseError
		if errors.As(err, &wsCloseErr) {
			code = wsCloseErr.Code
			reason = wsCloseErr.Reason
		}

		b.lock.Lock()
		delete(b.connections, client.ID())
		b.lock.Unlock()

		server.Close(code, reason)
		client.Close(code, reason)
	}()

	buffer := bytes.NewBuffer(make([]byte, 0))
	ctx := client.Context()

	for ctx.Err() == nil {
		buffer.Reset()

		msgType, reader, err = server.Reader(ctx)
		if err != nil {
			return
		}

		_, err = buffer.ReadFrom(reader)

		if err != nil {
			return
		}

		err = client.Send(msgType, buffer.Bytes())
		if err != nil {
			return
		}
	}

	err = ctx.Err()
}
