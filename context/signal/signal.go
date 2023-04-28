package signal

import (
	"stdlib/context"
	"stdlib/context/container"
	"stdlib/ptr"

	"golang.org/x/sync/errgroup"
)

func Send[T any](ctx context.Context, k string, data ...T) error {
	ch, err := container.Stream[T](ctx, k)
	if err != nil {
		return err
	}
	group, _ := errgroup.WithContext(ctx)
	group.SetLimit(len(data))
	for i := range data {
		group.Go(ch.Send(ctx, i, data))
	}
	return group.Wait()
}

func Recv[T any](ctx context.Context, k string) (T, error) {
	ch, err := container.Stream[T](ctx, k)
	if err != nil {
		return ptr.Value(new(T)), err
	}
	return ch.Recv(ctx)
}
