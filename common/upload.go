package common

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

type Upload struct {
	File        io.Reader
	Filename    string
	Size        int64
	ContentType string
}

func MarshalUpload(f Upload) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.Copy(w, f.File)
	})
}

func UnmarshalUpload(v interface{}) (Upload, error) {
	upload, ok := v.(Upload)
	if !ok {
		return Upload{}, fmt.Errorf("TypeError: %T is not an Upload", v)
	}
	return upload, nil
}
