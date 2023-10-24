package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		// stop reading when 2 bytes are readed and last second byte is \r
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	// returning line without last 2 bytes which are '\r\n'
	return line[:len(line)-2], n, nil
}

func (r *Resp) readInteger() (value int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}
	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}
	return int(i64), n, nil
}

func (r *Resp) readArray() (v Value, err error) {
	v.Typ = "array"

	arraySize, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	v.Array = make([]Value, 0)
	for i := 0; i < arraySize; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}
		v.Array = append(v.Array, val)
	}

	return v, nil
}

func (r *Resp) readBulk() (v Value, err error) {
	v.Typ = "bulk"

	strSize, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	buf := make([]byte, strSize)
	r.reader.Read(buf)

	// read '\r\n'
	r.readLine()

	v.Bulk = string(buf)
	return v, nil
}

func (r *Resp) Read() (Value, error) {
	_type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch _type {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		fmt.Printf("Unknown type: %v", string(_type))
		return Value{}, nil
	}
}
