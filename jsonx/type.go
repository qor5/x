package jsonx

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

type (
	RawMessage            = json.RawMessage
	Marshaler             = json.Marshaler
	Unmarshaler           = json.Unmarshaler
	UnmarshalTypeError    = json.UnmarshalTypeError
	InvalidUnmarshalError = json.InvalidUnmarshalError
	SyntaxError           = json.SyntaxError
	UnsupportedTypeError  = json.UnsupportedTypeError
	UnsupportedValueError = json.UnsupportedValueError
	MarshalerError        = json.MarshalerError
	Number                = json.Number
	Token                 = json.Token
	Delim                 = json.Delim
	Decoder               = json.Decoder
	Encoder               = json.Encoder
)

func NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}

func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func Valid(data []byte) bool {
	return json.Valid(data)
}

func HTMLEscape(dst *bytes.Buffer, src []byte) {
	json.HTMLEscape(dst, src)
}

func Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error {
	if err := json.Indent(dst, src, prefix, indent); err != nil {
		return errors.Wrap(err, "failed to indent json")
	}
	return nil
}

func Compact(dst *bytes.Buffer, src []byte) error {
	if err := json.Compact(dst, src); err != nil {
		return errors.Wrap(err, "failed to compact json")
	}
	return nil
}
