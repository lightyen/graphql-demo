package common

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// issue: gqlgen doesn't support unmarshal uint64

func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.FormatInt(i, 10)))
	})
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

func MarshalInt32(i int32) graphql.Marshaler {
	return MarshalInt64(int64(i))
}

func UnmarshalInt32(v interface{}) (int32, error) {
	switch t := v.(type) {
	case string:
		i, err := strconv.ParseInt(t, 10, 32)
		return int32(i), err
	case int64:
		i, err := strconv.ParseInt(strconv.FormatInt(t, 10), 10, 32)
		return int32(i), err
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Int32: %#v", v, v)
}

func MarshalInt16(i int16) graphql.Marshaler {
	return MarshalInt64(int64(i))
}

func UnmarshalInt16(v interface{}) (int16, error) {
	switch t := v.(type) {
	case string:
		i, err := strconv.ParseInt(t, 10, 16)
		return int16(i), err
	case int64:
		i, err := strconv.ParseInt(strconv.FormatInt(t, 10), 10, 16)
		return int16(i), err
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Int16: %#v", v, v)
}

func MarshalInt8(i int8) graphql.Marshaler {
	return MarshalInt64(int64(i))
}

func UnmarshalInt8(v interface{}) (int8, error) {
	switch t := v.(type) {
	case string:
		i, err := strconv.ParseInt(t, 10, 8)
		return int8(i), err
	case int64:
		i, err := strconv.ParseInt(strconv.FormatInt(t, 10), 10, 8)
		return int8(i), err
	}
	return 0, fmt.Errorf("Unable unmarshal %T to Int8: %#v", v, v)
}
