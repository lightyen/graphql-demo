package common

import (
	"encoding"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type UUID struct {
	id uuid.UUID
}

var (
	_ encoding.BinaryMarshaler   = (*UUID)(nil)
	_ encoding.BinaryUnmarshaler = (*UUID)(nil)
	_ encoding.TextMarshaler     = (*UUID)(nil)
	_ encoding.TextUnmarshaler   = (*UUID)(nil)
	_ json.Marshaler             = (*UUID)(nil)
	_ json.Unmarshaler           = (*UUID)(nil)
	_ graphql.Marshaler          = (*UUID)(nil)
	_ graphql.Unmarshaler        = (*UUID)(nil)
)

func NewUUID() UUID {
	return UUID{uuid.New()}
}

func (uuid UUID) MarshalBinary() ([]byte, error) {
	return uuid.id.MarshalBinary()
}

func (uuid *UUID) UnmarshalBinary(data []byte) error {
	return uuid.id.UnmarshalBinary(data)
}

func (uuid UUID) hexEncode() []byte {
	id, _ := uuid.MarshalBinary()
	dst := make([]byte, len(id)*2)
	hex.Encode(dst, id)
	return dst
}

func (uuid *UUID) hexDecode(data []byte) error {
	if len(data) != 32 {
		return fmt.Errorf("invalid Hex (got %d bytes)", len(data))
	}
	dst := make([]byte, 16)
	if _, err := hex.Decode(dst, data); err != nil {
		return err
	}
	return uuid.UnmarshalBinary(dst)
}

func (uuid UUID) MarshalText() ([]byte, error) {
	return uuid.hexEncode(), nil
}

func (uuid *UUID) UnmarshalText(data []byte) error {
	return uuid.hexDecode(data)
}

func (uuid *UUID) String() string {
	t, _ := uuid.MarshalText()
	return string(t)
}

func (uuid *UUID) FromString(s string) error {
	return uuid.UnmarshalText([]byte(s))
}

func (uuid UUID) MarshalJSON() ([]byte, error) {
	data, _ := uuid.MarshalText()
	b := make([]byte, 0, len(data)+2)
	b = append(b, '"')
	b = append(b, data...)
	b = append(b, '"')
	return b, nil
}

func (uuid *UUID) UnmarshalJSON(data []byte) error {
	// ignore null
	if string(data) == "null" {
		return nil
	}
	t, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	return uuid.UnmarshalText([]byte(t))
}

func (uuid UUID) MarshalGQL(w io.Writer) {
	data, _ := uuid.MarshalJSON()
	_, _ = w.Write(data)
}

func (uuid *UUID) UnmarshalGQL(v interface{}) error {
	switch t := v.(type) {
	case string:
		data := []byte(t)
		return uuid.UnmarshalText(data)
	}
	return fmt.Errorf("Unable unmarshal %T to UUID: %#v", v, v)
}

func (uuid UUID) Equal(b UUID) bool {
	for i := 0; i < 16; i++ {
		if uuid.id[i] != b.id[i] {
			return false
		}
	}
	return true
}
