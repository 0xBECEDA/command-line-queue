package command

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewCommand(t *testing.T) {
	tests := []struct {
		name            string
		message         string
		expectedCommand *Command
		expectedError   bool
	}{
		{
			name:    "Valid addItem command",
			message: "addItem('key', 'val')",
			expectedCommand: &Command{
				args:        []string{"key", "val"},
				commandType: Add,
			},
			expectedError: false,
		},
		{
			name:    "Valid deleteItem command",
			message: "deleteItem('key')",
			expectedCommand: &Command{
				args:        []string{"key"},
				commandType: Delete,
			},
			expectedError: false,
		},
		{
			name:    "Valid getItem command",
			message: "getItem('key')",
			expectedCommand: &Command{
				args:        []string{"key"},
				commandType: Get,
			},
			expectedError: false,
		},
		{
			name:    "Valid getAllItems command",
			message: "getAllItems()",
			expectedCommand: &Command{
				args:        []string{},
				commandType: GetAll,
			},
			expectedError: false,
		},
		{
			name:          "Invalid message",
			message:       "Invalid message",
			expectedError: true,
		},
		{
			name:          "Invalid addItem command",
			message:       "addItem('key')",
			expectedError: true,
		},
		{
			name:          "Invalid getItem command",
			message:       "getItem()",
			expectedError: true,
		},
		{
			name:          "Invalid getAllItems command",
			message:       "getAllItems('key')",
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			command, err := ParseCommand(test.message)
			if !test.expectedError {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

			assert.True(t, reflect.DeepEqual(test.expectedCommand, command))
		})
	}
}

func TestCommandIsValid(t *testing.T) {
	tests := []struct {
		name           string
		command        Command
		expectedResult bool
	}{
		{
			name:           "Valid addItem command",
			command:        Command{commandType: "addItem", args: []string{"key", "val"}},
			expectedResult: true,
		},
		{
			name:           "Valid deleteItem command",
			command:        Command{commandType: "deleteItem", args: []string{"key"}},
			expectedResult: true,
		},
		{
			name:           "Valid getItem command",
			command:        Command{commandType: "getItem", args: []string{"key"}},
			expectedResult: true,
		},
		{
			name:           "Valid getAllItems command",
			command:        Command{commandType: "getAllItems", args: []string{}},
			expectedResult: true,
		},
		{
			name:           "Invalid addItem command",
			command:        Command{commandType: "addItem", args: []string{"key"}},
			expectedResult: false,
		},
		{
			name:           "Invalid getItem command",
			command:        Command{commandType: "getItem", args: []string{}},
			expectedResult: false,
		},
		{
			name:           "Invalid command type",
			command:        Command{commandType: "InvalidType", args: []string{}},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.command.isValid()
			assert.Equal(t, test.expectedResult, result)
		})
	}
}
