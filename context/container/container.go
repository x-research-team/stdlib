package container

import (
	"fmt"
	"stdlib/context"
	"stdlib/ptr"

	"github.com/pkg/errors"
)

type TStream[T any] chan T

func Module[T any](ctx context.Context, k string) (T, error) {
	v, ok := ctx.Get(k).(T)
	if !ok {
		return ptr.Value(new(T)), fmt.Errorf("cannot cast %s to %T", k, v)
	}
	return v, nil
}

func Handler[T any](ctx context.Context, k string) (T, error) {
	v, ok := ctx.Handler(k).(T)
	if !ok {
		return ptr.Value(new(T)), fmt.Errorf("cannot cast %s to %T", k, v)
	}
	return v, nil
}

func Stream[T any](ctx context.Context, k string) (TStream[T], error) {
	ch, ok := ctx.Stream(k).(chan T)
	if !ok {
		return nil, fmt.Errorf("cannot cast %s to %T", k, ch)
	}
	return ch, nil
}

func (ch TStream[T]) Send(ctx context.Context, i int, data []T) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			return errors.WithMessage(ctx.Err(), "context canceled")
		case ch <- data[i]:
			return nil
		}
	}
}

func (ch TStream[T]) Recv(ctx context.Context) (T, error) {
	for {
		select {
		case <-ctx.Done():
			return ptr.Value(new(T)), errors.WithMessage(ctx.Err(), "context canceled")
		case v, ok := <-ch:
			if !ok {
				return ptr.Value(new(T)), errors.New("channel closed")
			}
			return v, nil
		}
	}
}
