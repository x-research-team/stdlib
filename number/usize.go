package number

type Size uint

type String string

func (s String) Size() Size {
	return Size(len(s))
}

func (s String) Str() string {
	return string(s)
}
