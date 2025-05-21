package command_bus

import (
	"context"
	"errors"
	"user/package/logger"

	"go.uber.org/zap"
)

type Command interface {
	CommandName() string
	Validate() error
}

type HandlerFunc func(context.Context, Command) (interface{}, error)

type CommandBus struct {
	handlers map[string]HandlerFunc
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]HandlerFunc),
	}
}

func (bus *CommandBus) RegisterHandler(commandName string, handler HandlerFunc) {
	bus.handlers[commandName] = handler
}

func (bus *CommandBus) Dispatch(ctx context.Context, cmd Command) (interface{}, error) {
	log := logger.FromContext(ctx)
	commandName := cmd.CommandName()

	handler, exists := bus.handlers[commandName]
	if !exists {
		err := errors.New("no handler registered for command: " + commandName)
		log.Error("no handler registered for command", zap.Error(err))
		return nil, err
	}

	result, err := handler(ctx, cmd)
	if err != nil {
		log.Error("handle command get failed", zap.Error(err))
		return nil, err
	}
	return result, nil
}
