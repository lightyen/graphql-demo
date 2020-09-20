package scalar

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalUint(i uint) graphql.Marshaler {
	return MarshalUint64(uint64(i))
}

func UnmarshalUint(v interface{}) (uint, error) {
	u, err := UnmarshalUint64(v)
	return uint(u), err
}

func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.FormatInt(i, 10)))
	})
}

func MarshalInt32(i int32) graphql.Marshaler {
	return MarshalInt64(int64(i))
}

func MarshalInt16(i int16) graphql.Marshaler {
	return MarshalInt64(int64(i))
}

func MarshalInt8(i int8) graphql.Marshaler {
	return MarshalInt64(int64(i))
}

func MarshalUint64(i uint64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.FormatUint(i, 10)))
	})
}

func MarshalUint32(i uint32) graphql.Marshaler {
	return MarshalUint64(uint64(i))
}

func MarshalUint16(i uint16) graphql.Marshaler {
	return MarshalUint64(uint64(i))
}

func MarshalUint8(i uint8) graphql.Marshaler {
	return MarshalUint64(uint64(i))
}

func UnmarshalInt64(v interface{}) (int64, error) {
	switch t := v.(type) {
	case string:
		return strconv.ParseInt(t, 10, 64)
	case int64:
		return int64(t), nil
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Int64: %#v", v, v)
}

func UnmarshalInt32(v interface{}) (int32, error) {
	switch t := v.(type) {
	case string:
		i, err := strconv.ParseInt(t, 10, 32)
		return int32(i), err
	case int64:
		return int32(t), nil
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Int32: %#v", v, v)
}

func UnmarshalInt16(v interface{}) (int16, error) {
	switch t := v.(type) {
	case string:
		i, err := strconv.ParseInt(t, 10, 16)
		return int16(i), err
	case int64:
		return int16(t), nil
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Int16: %#v", v, v)
}

func UnmarshalInt8(v interface{}) (int8, error) {
	switch t := v.(type) {
	case string:
		i, err := strconv.ParseInt(t, 10, 8)
		return int8(i), err
	case int64:
		return int8(t), nil
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Int8: %#v", v, v)
}

func UnmarshalUint64(v interface{}) (uint64, error) {
	switch t := v.(type) {
	case string:
		return strconv.ParseUint(t, 10, 64)
	case int64:
		return uint64(t), nil
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Uint64: %#v", v, v)
}

func UnmarshalUint32(v interface{}) (uint32, error) {
	switch t := v.(type) {
	case string:
		u, err := strconv.ParseUint(t, 10, 32)
		return uint32(u), err
	case int64:
		return uint32(t), nil
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Uint32: %#v", v, v)
}

func UnmarshalUint16(v interface{}) (uint16, error) {
	switch t := v.(type) {
	case string:
		u, err := strconv.ParseUint(t, 10, 16)
		return uint16(u), err
	case int64:
		return uint16(t), nil
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Uint16: %#v", v, v)
}

func UnmarshalUint8(v interface{}) (uint8, error) {
	switch t := v.(type) {
	case string:
		u, err := strconv.ParseUint(t, 10, 8)
		return uint8(u), err
	case int64:
		return uint8(t), nil
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Uint8: %#v", v, v)
}
