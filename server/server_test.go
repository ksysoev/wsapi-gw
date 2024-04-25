package server

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/ksysoev/wasabi/mocks"
)

func TestNewServer(t *testing.T) {
	addr := ":8080"
	server := NewServer(addr)

	if server.addr != addr {
		t.Errorf("Expected port %s, but got %s", addr, server.addr)
	}

	if len(server.channels) != 0 {
		t.Errorf("Expected empty channels slice, but got %d channels", len(server.channels))
	}

	if server.mutex == nil {
		t.Error("Expected non-nil mutex")
	}
}
func TestServer_AddChannel(t *testing.T) {
	// Create a new Server instance
	server := NewServer(":0")

	// Create a new channel
	channel := mocks.NewMockChannel(t)
	channel.EXPECT().Path().Return("test")

	// Add the channel to the server
	server.AddChannel(channel)

	// Check if the channel was added correctly
	if len(server.channels) != 1 {
		t.Errorf("Expected 1 channel, but got %d channels", len(server.channels))
	}

	if server.channels[0].Path() != "test" {
		t.Errorf("Expected channel name 'test', but got '%s'", server.channels[0].Path())
	}
}

func TestServer_WithBaseContext(t *testing.T) {
	// Create a new Server instance with a base context
	//lint:ignore SA1029 This is a test
	ctx := context.WithValue(context.Background(), "test", "test")

	server := NewServer(":0", WithBaseContext(ctx))

	// Check if the base context was set correctly
	if server.baseCtx == nil {
		t.Error("Expected non-nil base context")
	}

	if server.baseCtx.Value("test") != "test" {
		t.Errorf("Expected context value 'test', but got '%s'", server.baseCtx.Value("test"))
	}
}

func TestServer_Run(t *testing.T) {
	// Create a new Server instance
	server := NewServer(":0")

	// Run the server
	done := make(chan struct{})
	go func() {
		err := server.Run()
		if err != http.ErrServerClosed {
			t.Errorf("Expected error %v, but got %v", http.ErrServerClosed, err)
		}
		close(done)
	}()

	<-time.After(50 * time.Millisecond)
	err := server.http.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Error("Expected server to stop")
	}
}
