package context

import (
	"context"
	"os"
	"os/signal"

	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
)

// New creates a new Context
type Context interface {
	context.Context
	Set(k string, v Pluggable)
	Get(k string) Pluggable
	Delete(k string)
	Ctx() context.Context
	Cancel()
	Stop()
	Options(opts ...Option) Context
	Stream(k string) Asyncable
	Streams() map[string]Asyncable
	AddHandler(k string, handler Callable)
	RemoveHandler(k string)
	Handler(k string) Callable
	Handlers() map[string]Callable
}

// type Option func(ctx *container)
type (
	Callable  any
	Asyncable any
	Pluggable any

	components map[string]Pluggable
	streams    map[string]Asyncable
	callers    map[string]Callable
)

// New creates a new Context
type container struct {
	context.Context
	container components
	cancel    context.CancelFunc
	streams   streams
	stop      context.CancelFunc
	handlers  callers
	logger    *zerolog.Logger
}

func Init[T any](capacity int) map[string]T {
	return make(map[string]T, capacity)
}

func logger() *zerolog.Logger {
	l := zerolog.New(colorable.NewColorableStdout()).With().Caller().Timestamp().Logger()
	return &l
}

func New(opts ...Option) Context {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)

	c := &container{
		Context:   ctx,
		container: Init[Pluggable](100),
		streams:   Init[Asyncable](100),
		handlers:  Init[Callable](100),
		stop:      stop,
		logger:    logger(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Delete deletes a key from the context
//
//	ctx.Delete("foo")
func (c *container) Delete(k string) {
	defer c.logger.Info().Msgf("[MODULE] Delete: %s", k)
	delete(c.container, k)
}

// Options adds options to the context
func (c *container) Options(opts ...Option) Context {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *container) Stream(k string) Asyncable {
	return c.Streams()[k]
}

func (c *container) Streams() map[string]Asyncable {
	return c.streams
}

func (c *container) Set(k string, v Pluggable) {
	defer c.logger.Info().Msgf("[MODULE] Packed: %s", k)
	c.container[k] = v
}

func (c *container) Get(k string) Pluggable {
	defer c.logger.Info().Msgf("[Module] Unpaked: %s", k)
	return c.container[k]
}

func (c *container) Ctx() context.Context {
	return c.Context
}

func (c *container) Cancel() {
	if c.cancel != nil {
		defer c.logger.Info().Msgf("The Iteration Cancelled")
		c.cancel()
	}
}

func (c *container) Stop() {
	if c.stop != nil {
		defer c.logger.Info().Msgf("The Iteration Stopped")
		c.stop()
	}
}

func (c *container) AddHandler(k string, handler Callable) {
	c.handlers[k] = handler
}

func (c *container) RemoveHandler(k string) {
	delete(c.handlers, k)
}

func (c *container) Handler(k string) Callable {
	return c.Handlers()[k]
}

func (c *container) Handlers() map[string]Callable {
	return c.handlers
}
