package types

import "encoding/json"

type String string

func (s String) Size() USize {
	return USize(len(s))
}

func (s String) Value() string {
	return string(s)
}

func (s String) Bytes() []byte {
	return []byte(s)
}

func (s String) JSON() ([]byte, error) {
	return json.Marshal(s)
}
