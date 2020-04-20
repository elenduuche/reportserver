package utils

import (
	"io"
)

type Encoder interface {
	Marshal(value interface{}) ([]byte, error)
	Encode(w io.Writer, in interface{}) error
}

type Decoder interface {
	Unmarshal(value []byte, out interface{}) error
	Decode(w io.Reader, out interface{}) error
}
