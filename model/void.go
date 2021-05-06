package model

import (
	"encoding/json"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

type Void struct{}

func MarshalVoid(v interface{}) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		data, _ := json.Marshal(v)
		_, _ = w.Write(data)
	})
}

func UnmarshalVoid(v interface{}) (Void, error) {
	// do nothing
	return Void{}, nil
}
