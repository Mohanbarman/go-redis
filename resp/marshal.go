package resp

import (
	"fmt"
	"strconv"
)

type Value struct {
	Typ   string
	Str   string
	Num   int64
	Bulk  string
	Array []Value
}

func (v Value) Marshal() []byte {
	switch v.Typ {
	case "array":
		return v.marshalArray()
	case "bulk":
		return v.marshalBulk()
	case "string":
		return v.marshalString()
	case "null":
		return v.marshalNull()
	case "error":
		return v.marshalError()
  case "integer":
    return v.marshalInteger()
	default:
		return []byte{}
	}
}

func (v Value) marshalString() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, STRING)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalBulk() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, BULK)                         // adding bulk character
	bytes = append(bytes, strconv.Itoa(len(v.Bulk))...) // adding size of bulk string
	bytes = append(bytes, '\r', '\n')                   // CRLF
	bytes = append(bytes, v.Bulk...)                    // bulk string value
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalArray() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(len(v.Array))...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < len(v.Array); i++ {
		bytes = append(bytes, v.Array[i].Marshal()...)
	}
	return bytes
}

func (v Value) marshalError() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, ERROR)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalNull() []byte {
	return []byte("$-1\r\n")
}

func (v Value) marshalInteger() []byte {
  bytes := make([]byte, 0)
  bytes = append(bytes, INTEGER)
  bytes = append(bytes, fmt.Sprintf("%d", v.Num)...)
  bytes = append(bytes, '\r', '\n')
  return bytes
}
