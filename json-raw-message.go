package utils

import (
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	_strNull                                    = "null"
	_strJsonRawMessageUnmarshalJSONOnNilPointer = "utils.JsonRawMessage: UnmarshalJSON on nil pointer"
)

// JsonRawMessage is a raw encoded JSON value, copy from encoding/json RawMessage.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type JsonRawMessage []byte

var _ json.Marshaler = (*JsonRawMessage)(nil)
var _ json.Unmarshaler = (*JsonRawMessage)(nil)

// MarshalJSON returns m as the JSON encoding of m.
func (m JsonRawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte(_strNull), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *JsonRawMessage) UnmarshalJSON(data []byte) error {

	if m == nil {
		return errors.New(_strJsonRawMessageUnmarshalJSONOnNilPointer)
	}
	// fmt.Println("len(*m)", len(*m), "cap(*m)", cap(*m))
	// if m in a initial val, it will print "len(*m) 0 cap(*m) 0"
	// so we should judge the cap(*m), and alloc memory for m if necessary
	var lenData = len(data)
	if cap(*m) < lenData {
		*m = make(JsonRawMessage, lenData)
	} else {
		*m = (*m)[:lenData]
	}
	copy((*m), data)
	return nil
}
