package model

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

type IP struct {
	net.IP
}

func (ip IP) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(ip.String())))
}

func (ip *IP) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		ip.IP = net.ParseIP(v)
		return nil
	default:
		return fmt.Errorf("TypeError: %T is not an IP", v)
	}
}
