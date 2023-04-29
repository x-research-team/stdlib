package types

type Slice[T any] []T

func Vector[T any](o ...Size) Slice[T] {
	if len(o) == 0 {
		return make(Slice[T], 0)
	}
	if len(o) == 1 {
		return make(Slice[T], o[0])
	}
	if len(o) > 2 {
		return make(Slice[T], o[0], o[1])
	}
	return make(Slice[T], o[0])
}

func (s Slice[T]) Len() Size {
	return Size(len(s))
}

func (s Slice[T]) Add(v ...T) Slice[T] {
	return append(s, v...)
}

func (s Slice[T]) Get(i Size) T {
	return s[i]
}

func (s Slice[T]) Set(i Size, v T) Slice[T] {
	s[i] = v
	return s
}

func (s Slice[T]) Copy() Slice[T] {
	return append(Slice[T]{}, s...)
}

func (s Slice[T]) Remove(i Size) Slice[T] {
	return append(s[:i], s[i+1:]...)
}
