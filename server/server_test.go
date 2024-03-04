package server

import (
	"bytes"
	"command-queue/internal/command"
	"context"
	"os"
	"testing"

	"command-queue/internal/util/queue"
	"github.com/stretchr/testify/assert"
)

func TestServerStart(t *testing.T) {
	// Create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize memQueue
	memQ := queue.NewMemQueue(5)

	s := NewServer(memQ, 1)

	inputCommands := []string{"add('key1,'value1')", "delete('key2')", "get_all()"}

	// create a buffer with input commands
	inputBuffer := bytes.NewBufferString("")
	for _, cmd := range inputCommands {
		inputBuffer.WriteString(cmd + "\n")
	}

	// start the server in a separate goroutine
	go func() {
		err := s.Start(ctx)
		assert.Nilf(t, err, "Start returned an error: %v", err)
	}()

	// Send input commands to the memQueue
	for _, cmd := range inputCommands {
		err := memQ.SendMessage(cmd)
		assert.Nilf(t, err, "SendMessage returned an error: %v", err)
	}
	err := s.Stop()
	assert.Nilf(t, err, "Stop returned an error: %v", err)

	// Simulate cancellation of the context to stop the server
	cancel()
}

func TestProcessCommand(t *testing.T) {
	server := NewServer(nil, 1)

	// delete test files if they exist
	os.Remove("key2_1")
	os.Remove("allItems_2")

	tests := []struct {
		name    string
		message string
	}{
		{name: "AddItem1", message: command.NewAddCommand("key1", "value1").String()},
		{name: "AddItem2", message: command.NewAddCommand("key2", "value2").String()},
		{name: "DeleteItem1", message: command.NewDeleteCommand("key1").String()},
		{name: "GetItem1", message: command.NewGetCommand("key2").String()},
		{name: "GetAllItems1", message: command.NewGetAllCommand().String()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.processCommand(tt.message)
		})
	}
	keys, values := server.orderedMap.GetAll()
	assert.Equal(t, []string{"key2"}, keys)
	assert.Equal(t, []interface{}{"value2"}, values)

	bt, err := os.ReadFile("key2_1")
	assert.Nilf(t, err, "Error reading file: %v", err)
	assert.Equal(t, "key2 : value2\n", string(bt))
	os.Remove("key2_1")

	bt, err = os.ReadFile("allItems_2")
	assert.Nilf(t, err, "Error reading file: %v", err)
	assert.Equal(t, "key2 : value2\n", string(bt))
	os.Remove("allItems_2")
}
