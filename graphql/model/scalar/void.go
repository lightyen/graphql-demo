package scalar

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalVoid(v interface{}) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// do nothing
	})
}

func UnmarshalVoid(v interface{}) (interface{}, error) {
	// do nothing
	return nil, nil
}
