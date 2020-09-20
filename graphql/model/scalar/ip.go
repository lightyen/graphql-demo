package scalar

import (
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalIP(ip net.IP) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(ip.String())))
	})
}

func UnmarshalIP(v interface{}) (net.IP, error) {
	switch v := v.(type) {
	case string:
		return net.ParseIP(v), nil
	default:
		return nil, fmt.Errorf("TypeError: %T is not an IP", v)
	}
}
