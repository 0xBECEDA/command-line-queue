package command

import (
	"errors"
	"fmt"
	"strings"
)

var InvalidCommandError = errors.New("command is invalid")

type CommandType string

const (
	Undefined CommandType = ""
	Add       CommandType = "addItem"
	Delete    CommandType = "deleteItem"
	Get       CommandType = "getItem"
	GetAll    CommandType = "getAllItems"
)

type Command struct {
	args        []string
	commandType CommandType
}

func ParseCommand(message string) (*Command, error) {
	command := &Command{
		args:        getArgs(message),
		commandType: getCommandType(message),
	}
	if !command.isValid() {
		return nil, fmt.Errorf("%w: Invalid message: %s\n", InvalidCommandError, message)
	}

	return command, nil
}

func NewAddCommand(key, value string) *Command {
	return &Command{
		commandType: Add,
		args:        []string{key, value},
	}
}

func NewDeleteCommand(key string) *Command {
	return &Command{
		commandType: Delete,
		args:        []string{key},
	}
}

func NewGetCommand(key string) *Command {
	return &Command{
		commandType: Get,
		args:        []string{key},
	}
}

func NewGetAllCommand() *Command {
	return &Command{
		commandType: GetAll,
		args:        []string{},
	}
}

func getCommandType(message string) CommandType {
	message = strings.TrimSpace(strings.Split(message, "(")[0])
	switch message {
	case "addItem":
		return Add
	case "deleteItem":
		return Delete
	case "getItem":
		return Get
	case "getAllItems":
		return GetAll
	default:
		return Undefined
	}
}

func getArgs(message string) []string {
	parts := strings.Split(message, "(")
	if len(parts) < 2 {
		return nil
	}
	parts = strings.Split(parts[1], ")")
	parts = strings.Split(parts[0], ",")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Trim(strings.TrimSpace(parts[i]), "'")
		if len(parts[i]) == 0 {
			parts = append(parts[:i], parts[i+1:]...)
			i--
		}
	}
	return parts
}

func (c *Command) isValid() bool {
	switch c.commandType {
	case Add:
		if len(c.args) == 2 {
			return true
		}
	case Delete, Get:
		if len(c.args) == 1 {
			return true
		}
	case GetAll:
		if len(c.args) == 0 {
			return true
		}
	}
	return false
}

func (c *Command) Key() string {
	return c.args[0]
}

func (c *Command) Type() CommandType {
	return c.commandType
}

func (c *Command) Value() string {
	if len(c.args) < 2 {
		return ""
	}
	return c.args[1]
}

func (c *Command) String() string {
	switch len(c.args) {
	case 0:
		return fmt.Sprintf("%s()", c.commandType)
	default:
		return fmt.Sprintf("%s('%s')", c.commandType, strings.Join(c.args, "', '"))
	}
}
