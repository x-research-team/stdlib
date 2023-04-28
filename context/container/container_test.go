package container

import (
	"stdlib/context"
	"stdlib/context/modules"
	"testing"

	"github.com/goccy/go-reflect"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestModule(t *testing.T) {
	type args struct {
		ctx context.Context
		k   string
	}
	tests := []struct {
		name    string
		args    args
		want    *zerolog.Logger
		wantErr bool
	}{
		{
			name: "load logger",
			args: args{
				ctx: context.New(context.Logger(&zerolog.Logger{})),
				k:   modules.Logger,
			},
			want: &zerolog.Logger{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Module[*zerolog.Logger](tt.args.ctx, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStream(t *testing.T) {
	type args struct {
		ctx context.Context
		k   string
	}
	tests := []struct {
		name     string
		args     args
		want     TStream[string]
		wantErr  bool
		listener func(t *testing.T, ctx context.Context)
	}{
		{
			name: "stream ok",
			args: args{
				ctx: context.New(context.Stream("stream", make(chan string, 10))),
				k:   "stream",
			},
			want: make(chan string),
			listener: func(t *testing.T, ctx context.Context) {
				s, err := Stream[string](ctx, "stream")
				if err != nil {
					t.Fatal(err)
				}
				for {
					v := <-s
					if v == "" {
						continue
					}
					t.Logf("got %s", v)
					assert.IsType(t, string(""), v, "should be string")
					return
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Stream[string](tt.args.ctx, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stream() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("Stream() = %v, want %v", got, tt.want)
			}
			var g errgroup.Group
			g.Go(func() error {
				tt.listener(t, tt.args.ctx)
				return nil
			})
			got <- "Hello World"
			require.NoError(t, g.Wait())
		})
	}
}

func TestHandler(t *testing.T) {
	type args struct {
		ctx context.Context
		k   string
	}
	tests := []struct {
		name    string
		args    args
		want    func() error
		wantErr bool
	}{
		{
			name: "handler ok",
			args: args{
				ctx: context.New(context.Handler("handler", func() error {
					return errors.New("test")
				})),
				k: "handler",
			},
			want: func() error {
				return errors.New("test")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Handler[func() error](tt.args.ctx, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got().Error(), tt.want().Error()) {
				t.Errorf("Handler() = %v, want %v", got(), tt.want())
			}
		})
	}
}
