package scalar

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.FormatInt(t.Unix(), 10)))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case int64:
		return time.Unix(v, 0), nil
	case int:
		return time.Unix(int64(v), 0), nil
	default:
		return time.Time{}, fmt.Errorf("TypeError: %T is not a Time", v)
	}
}
