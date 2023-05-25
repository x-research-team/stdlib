package types

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Slice[T constraints.Ordered] struct {
	items []T
	len   USize
	cap   USize
}

type sliceOption[T constraints.Ordered] func(s *Slice[T])

func Len[T constraints.Ordered](v USize) sliceOption[T] {
	return func(s *Slice[T]) {
		s.len = v
	}
}

func Cap[T constraints.Ordered](v USize) sliceOption[T] {
	return func(s *Slice[T]) {
		s.cap = v
	}
}

func Items[T constraints.Ordered](v ...T) sliceOption[T] {
	return func(s *Slice[T]) {
		s.items = v
	}
}

func Vector[T constraints.Ordered](opts ...sliceOption[T]) *Slice[T] {
	s := &Slice[T]{
		items: make([]T, 0),
		len:   0,
		cap:   0,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s Slice[T]) Len() USize {
	return USize(len(s.items))
}

func (s *Slice[T]) Add(v ...T) *Slice[T] {
	s.items = append(s.items, v...)
	return s
}

func (s Slice[T]) At(i USize) T {
	return s.items[i]
}

func (s *Slice[T]) Set(i USize, v T) *Slice[T] {
	s.items[i] = v
	return s
}

func (s Slice[T]) Copy() *Slice[T] {
	return &Slice[T]{
		items: s.items,
		len:   s.len,
		cap:   s.cap,
	}
}

func (s *Slice[T]) Remove(i USize) *Slice[T] {
	s.items = append(s.items[:i], s.items[i+1:]...)
	return s
}

func (s *Slice[T]) Clear() *Slice[T] {
	s.items = make([]T, 0)
	s.len = 0
	s.cap = 0
	return s
}

func (s *Slice[T]) Slice() []T {
	return s.items
}

func (s *Slice[T]) Swap(i, j USize) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s *Slice[T]) Sort() *Slice[T] {
	sort.Slice(s.items, func(i, j int) bool {
		return s.items[i] < s.items[j]
	})
	return s
}
