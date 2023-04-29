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
