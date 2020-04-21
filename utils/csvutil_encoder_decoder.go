package utils

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
)

type csvutilEncoderDecoder struct {
}

func (*csvutilEncoderDecoder) Marshal(value interface{}) ([]byte, error) {
	return csvutil.Marshal(value)
}

func (*csvutilEncoderDecoder) Encode(w io.Writer, in interface{}) error {
	return csvutil.NewEncoder(csv.NewWriter(w)).Encode(in)
}

func (*csvutilEncoderDecoder) Unmarshal(value []byte, out interface{}) error {
	return csvutil.Unmarshal(value, out)
}

func (*csvutilEncoderDecoder) Decode(r io.Reader, out interface{}) error {
	//return csvutil.NewDecoder(csv.NewReader(r), ).Decode(out)
	return nil
}

//NewCSVUtilEncoder returns a csv Encoder
func NewCSVUtilEncoder() Encoder {
	return new(csvutilEncoderDecoder)
}

//NewCSVUtilDecoder returns a csv Decoder
func NewCSVUtilDecoder() Decoder {
	return new(csvutilEncoderDecoder)
}
