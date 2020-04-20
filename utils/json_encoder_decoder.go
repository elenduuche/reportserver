package utils

import (
	"encoding/json"
	"io"
)

type jsonEncoderDecoder struct {
}

func (*jsonEncoderDecoder) Marshal(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (*jsonEncoderDecoder) Encode(w io.Writer, in interface{}) error {
	return json.NewEncoder(w).Encode(in)
}

func (*jsonEncoderDecoder) Unmarshal(value []byte, out interface{}) error {
	return json.Unmarshal(value, out)
}

func (*jsonEncoderDecoder) Decode(r io.Reader, out interface{}) error {
	return json.NewDecoder(r).Decode(out)
}

//NewJSONEncoder returns a Json Encoder
func NewJSONEncoder() Encoder {
	return new(jsonEncoderDecoder)
}

//NewJSONDecoder returns a Json Decoder
func NewJSONDecoder() Decoder {
	return new(jsonEncoderDecoder)
}
