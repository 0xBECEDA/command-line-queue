package client

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"command-queue/internal/util/queue"
)

func TestClientStart(t *testing.T) {
	memQ := queue.NewMemQueue(10)
	inputCommands := []string{"add('key1,'value1')", "delete('key2')", "get_all()"}

	inputBuffer := bytes.NewBufferString("")
	for _, cmd := range inputCommands {
		inputBuffer.WriteString(cmd + "\n")
	}

	ctx := context.Background()

	c := NewClient(inputBuffer, memQ)

	err := c.Start(ctx)
	assert.Nilf(t, err, "Start returned an error: %v", err)
	err = c.Stop()
	assert.Nilf(t, err, "Stop returned an error: %v", err)

	// Check if messages were sent to the queue
	receivedMessages, _ := memQ.ReceiveMessage()
	for _, expectedMsg := range inputCommands {
		receivedMsg := <-receivedMessages
		if receivedMsg != expectedMsg {
			t.Errorf("Expected message %s, but got %s", expectedMsg, receivedMsg)
		}
	}

	// Check if there are no more messages in the queue
	select {
	case receivedMsg := <-receivedMessages:
		t.Errorf("Unexpected message in the queue: %s", receivedMsg)
	default:
		// No message in the channel, as expected
	}
}
