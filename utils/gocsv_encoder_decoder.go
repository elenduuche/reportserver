package utils

import (
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"
)

type gocsvEncoderDecoder struct {
}

func (*gocsvEncoderDecoder) Marshal(value interface{}) ([]byte, error) {
	return gocsv.MarshalBytes(value)
}

func (*gocsvEncoderDecoder) Encode(w io.Writer, in interface{}) error {
	return gocsv.Marshal(in, w)
}

func (*gocsvEncoderDecoder) Unmarshal(value []byte, out interface{}) error {
	return gocsv.UnmarshalBytes(value, out)
}

func (*gocsvEncoderDecoder) Decode(r io.Reader, out interface{}) error {
	return gocsv.Unmarshal(r, out)
}

func (c *gocsvEncoderDecoder) WriteToCSVFile(in []string, w *csv.Writer) error {
	if err := w.Write(in); err != nil {
		return err
	}
	return nil
}

//NewGoCSVEncoder returns a csv Encoder
func NewGoCSVEncoder() Encoder {
	return new(gocsvEncoderDecoder)
}

//NewGoCSVDecoder returns a csv Decoder
func NewGoCSVDecoder() Decoder {
	return new(gocsvEncoderDecoder)
}
