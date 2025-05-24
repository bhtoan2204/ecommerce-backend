package command_bus

import (
	"context"
	"fmt"
)

type Command interface {
	CommandName() string
}

type rawHandler func(ctx context.Context, c any) (any, error)

type CommandBus struct {
	handlers map[string]rawHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]rawHandler),
	}
}

func RegisterHandler[C Command, R any](
	b *CommandBus,
	handler func(context.Context, C) (R, error),
) {
	var zeroC C
	name := zeroC.CommandName()

	b.handlers[name] = func(ctx context.Context, raw any) (any, error) {
		cmd, ok := raw.(C)
		if !ok {
			return *new(R), fmt.Errorf("invalid command type for %s", name)
		}
		return handler(ctx, cmd)
	}
}

func Dispatch[C Command, R any](b *CommandBus, ctx context.Context, cmd C) (R, error) {
	var zeroR R

	handler, ok := b.handlers[cmd.CommandName()]
	if !ok {
		return zeroR, fmt.Errorf("no handler for command %s", cmd.CommandName())
	}

	result, err := handler(ctx, cmd)
	if err != nil {
		return zeroR, err
	}

	res, ok := result.(R)
	if !ok {
		return zeroR, fmt.Errorf("invalid result type for command %s", cmd.CommandName())
	}
	return res, nil
}
