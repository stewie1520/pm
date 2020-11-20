package model

import (
	"net"
	"time"
)

// Activity represent activity collection
type Activity struct {
	Action string
	Time   time.Time
	IP     net.IPAddr
}
