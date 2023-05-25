package ptr

import "unsafe"

func Ptr[T any](v T) *T {
	return &v
}

func Value[T any](v *T) T {
	if v == nil {
		return *new(T)
	}
	return *v
}

func Address[T any](v *T) uintptr {
	return uintptr(unsafe.Pointer(v))
}
