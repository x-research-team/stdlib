package types

type USize uint

func (s USize) Value() uint {
	return uint(s)
}

type Size int

func (s Size) Value() int {
	return int(s)
}
