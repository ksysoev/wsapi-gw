package channel

import (
	"context"
	"net/http"

	"github.com/ksysoev/wasabi"
	"nhooyr.io/websocket"
)

// DefaultChannel is default implementation of Channel
type DefaultChannel struct {
	path         string
	disptacher   wasabi.Dispatcher
	connRegistry *DefaultConnectionRegistry
	ctx          context.Context
	middlewares  []Middlewere
	config       channelConfig
}

type channelConfig struct {
	originPatterns []string
}

type Option func(*channelConfig)

// NewDefaultChannel creates new instance of DefaultChannel
// path - channel path
// dispatcher - dispatcher to use
// connRegistry - connection registry to use
// reqParser - request parser to use
// returns new instance of DefaultChannel
func NewDefaultChannel(
	path string,
	dispatcher wasabi.Dispatcher,
	opts ...Option,
) *DefaultChannel {
	config := channelConfig{
		originPatterns: []string{"*"},
	}

	for _, opt := range opts {
		opt(&config)
	}

	return &DefaultChannel{
		path:         path,
		disptacher:   dispatcher,
		connRegistry: NewDefaultConnectionRegistry(),
		middlewares:  make([]Middlewere, 0),
		config:       config,
	}
}

// Path returns url path for channel
func (c *DefaultChannel) Path() string {
	return c.path
}

// Handler returns http.Handler for channel
func (c *DefaultChannel) Handler() http.Handler {
	return c.setContext(c.wrapMiddleware(c.wsConnectionHandler()))
}

// wsConnectionHandler handles the WebSocket connection and sets up the necessary components for communication.
func (c *DefaultChannel) wsConnectionHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ws, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: c.config.originPatterns,
		})

		if err != nil {
			return
		}

		conn := c.connRegistry.AddConnection(ctx, ws, c.disptacher.Dispatch)
		conn.HandleRequests()
	})
}

// SetContext sets context for channel
func (c *DefaultChannel) SetContext(ctx context.Context) {
	c.ctx = ctx
}

// Use adds middlewere to channel
func (c *DefaultChannel) Use(middlewere Middlewere) {
	c.middlewares = append(c.middlewares, middlewere)
}

// useMiddleware applies middlewares to handler
func (c *DefaultChannel) wrapMiddleware(handler http.Handler) http.Handler {
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		handler = c.middlewares[i](handler)
	}

	return handler
}

// setContext sets context for handler
func (c *DefaultChannel) setContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(c.ctx))
	})
}

// WithOriginPatterns sets the origin patterns for the channel.
// The origin patterns are used to validate the Origin header of the WebSocket handshake request.
// If the Origin header does not match any of the patterns, the connection is rejected.
func WithOriginPatterns(patterns ...string) Option {
	return func(c *channelConfig) {
		c.originPatterns = patterns
	}
}
