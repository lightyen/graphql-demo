package common

import (
	"net"
	"time"
)

// Reference: https://gqlgen.com/reference/scalars/

type Device struct {
	// Current IP.
	IP net.IP
	// Current time.
	Now time.Time
	// Description
	Description string
}
