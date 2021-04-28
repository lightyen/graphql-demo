package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Time struct{ time.Time }

var (
	_ json.Marshaler      = (*Time)(nil)
	_ json.Unmarshaler    = (*Time)(nil)
	_ graphql.Marshaler   = (*Time)(nil)
	_ graphql.Unmarshaler = (*Time)(nil)
)

func Now() Time {
	return Time{time.Now()}
}

func (t Time) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", t.UnixNano()/1e6)
	return []byte(stamp), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// ignore null
	s := string(data)
	if s == "null" {
		return nil
	}
	s, err := strconv.Unquote(s)
	if err != nil {
		return err
	}
	ts, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	ms := ts % 1e3
	*t = Time{time.Unix(ts/1e3, ms*1e6)}
	return nil
}

func (t Time) MarshalGQL(w io.Writer) {
	data, _ := t.MarshalJSON()
	_, _ = w.Write(data)
}

func (t *Time) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case int64:
		if v < 0 {
			return errors.New("Parse time: value should not be negtive")
		}
		ms := v % 1e3
		*t = Time{time.Unix(v/1e3, ms*1e6)}
		return nil
	default:
		return fmt.Errorf("TypeError: %T is not a Time", v)
	}
}
