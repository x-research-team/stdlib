package mapper

import (
	"github.com/goccy/go-reflect"
	"golang.org/x/exp/constraints"
)

type container[T, V constraints.Ordered] struct {
	from T
	to   V
}

type containerOption[T, V constraints.Ordered] func(c *container[T, V])

func From[T, V constraints.Ordered](v T) containerOption[T, V] {
	return func(c *container[T, V]) {
		c.from = v
	}
}

func To[T, V constraints.Ordered](v V) containerOption[T, V] {
	return func(c *container[T, V]) {
		c.to = v
	}
}

func New[T, V constraints.Ordered](opts ...containerOption[T, V]) *container[T, V] {
	c := &container[T, V]{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *container[T, V]) Map() V {
	from := c.from
	to := c.to
	rto := reflect.ValueOf(to)
	if rto.Kind() == reflect.Ptr {
		rto = rto.Elem()
	}
	rfrom := reflect.ValueOf(from).Elem()
	if rfrom.Kind() == reflect.Ptr {
		rfrom = rfrom.Elem()
	}
	c.mapper(rfrom, rto)
	return rto.Interface().(V)
}

func (c *container[T, V]) mapper(from reflect.Value, to reflect.Value) {
	if from.Kind() == reflect.Ptr {
		from = from.Elem()
	}
	if to.Kind() == reflect.Ptr {
		to = to.Elem()
	}

	if to.Kind() == reflect.Slice && from.Kind() == reflect.Slice {
		for i := 0; i < to.Len(); i++ {
			c.mapper(from.Index(i), to.Index(i))
		}
	} else if to.Kind() == reflect.Map && from.Kind() == reflect.Map {
		for _, k := range from.MapKeys() {
			c.mapper(from.MapIndex(k), to.MapIndex(k))
		}
	} else if to.Kind() == reflect.Struct && from.Kind() == reflect.Struct {
		for i := 0; i < to.NumField(); i++ {
			c.mapper(from.Field(i), to.Field(i))
		}
	} else {
		if to.Type().Kind() == from.Type().Kind() {
			if to.CanSet() {
				to.Set(from)
			}
		}
	}
}
