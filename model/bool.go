package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalBool(b bool) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.FormatBool(b)))
	})
}

func UnmarshalBool(v interface{}) (bool, error) {
	switch v := v.(type) {
	case bool:
		return v, nil
	case int64:
		fmt.Println("int64")
		return v > 0, nil
	case int:
		fmt.Println("int")
		return v > 0, nil
	case string:
		if v == "1" || v == "yes" || v == "on" || v == "true" || v == "enabled" {
			return true, nil
		} else if v == "0" || v == "no" || v == "off" || v == "false" || v == "disabled" {
			return false, nil
		}
		return false, fmt.Errorf("TypeError: %T is not a bool", v)
	default:
		return false, fmt.Errorf("TypeError: %T is not a bool", v)
	}
}
