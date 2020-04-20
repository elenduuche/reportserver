package utils

import (
	"encoding/csv"
	"io"
)

type csvEncoderDecoder struct {
}

func (*csvEncoder) Marshal(value interface{}) ([]byte, error) {
	return csv.Marshal(value)
}

func (*csvEncoder) Encode(w io.Writer, in interface{}) error {
	csvWriter := 
	csvWriter.Write()
	return csv.NewEncoder(w).Encode(in)
}

func (*csvEncoder) Unmarshal(value []byte, out interface{}) error {
	return csv.Unmarshal(value, out)
}

func (*csvEncoder) Decode(r io.Reader, out interface{}) error {
	return csv.NewDecoder(r).Decode(out)
}

//NewJSONEncoder returns a Json Encoder
func NewJSONEncoder() Encoder {
	return new(jsonEncoderDecoder)
}

//NewJSONDecoder returns a Json Decoder
func NewJSONDecoder() Decoder {
	return new(jsonEncoderDecoder)
}
