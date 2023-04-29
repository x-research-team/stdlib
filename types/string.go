package types

type String string

func (s String) Size() Size {
	return Size(len(s))
}

func (s String) String() string {
	return string(s)
}

func (s String) ToBytes() []byte {
	return []byte(s)
}
