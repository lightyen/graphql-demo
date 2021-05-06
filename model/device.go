package model

import (
	"net"
	"time"
)

type Device struct {
	// Current IP.
	IP net.IP
	// Current time.
	Now time.Time
	// Description
	Description string
}
