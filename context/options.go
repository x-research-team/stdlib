package context

import (
	"context"
	"stdlib/context/modules"
	"time"

	"github.com/rs/zerolog"
	"github.com/x-research-team/sqlx"
)

type Option func(ctx *container)

func Logger(l *zerolog.Logger) Option {
	return func(ctx *container) {
		defer l.Info().Msgf("[MODULE] Logger Initiated")
		ctx.Set(modules.Logger, l)
		ctx.logger = l
	}
}

func Timeout(timeout time.Duration) Option {
	return func(ctx *container) {
		defer ctx.logger.Info().Msgf("[MODULE] Timeout Initiated")
		c, cancel := context.WithTimeout(ctx.Context, timeout)
		ctx.Context = c
		ctx.cancel = cancel
	}
}

func Deadline(deadline time.Time) Option {
	return func(ctx *container) {
		defer ctx.logger.Info().Msgf("[MODULE] Deadline Initiated")
		c, cancel := context.WithDeadline(ctx.Context, deadline)
		ctx.Context = c
		ctx.cancel = cancel
	}
}

func Cancel() Option {
	return func(ctx *container) {
		defer ctx.logger.Info().Msgf("[MODULE] Cancel Initiated")
		c, cancel := context.WithCancel(ctx.Context)
		ctx.Context = c
		ctx.cancel = cancel
	}
}

func Stream(k string, stream any) Option {
	return func(ctx *container) {
		defer ctx.logger.Info().Msgf("[MODULE] Stream Initiated")
		ctx.streams[k] = stream
	}
}

func Database(db sqlx.ExtContext) Option {
	return func(ctx *container) {
		defer ctx.logger.Info().Msgf("[MODULE] Database Initiated")
		ctx.Set(modules.Database, db)
	}
}

func Transaction(tx sqlx.ExtContext) Option {
	return func(ctx *container) {
		defer ctx.logger.Info().Msgf("[MODULE] Transaction Initiated")
		ctx.Set(modules.Transaction, tx)
	}
}

func Handler(k string, handler Callable) Option {
	return func(ctx *container) {
		defer ctx.logger.Info().Msgf("[MODULE] Handler Initiated")
		ctx.AddHandler(k, handler)
	}
}
