package model

import (
	"errors"
	"net"
	"time"
)

// ErrNoRecords occurring when there is no record retrieved from query
var ErrNoRecords = errors.New("models: no matching record found")

// Activity represent activity collection
type Activity struct {
	Action string
	Time   time.Time
	IP     net.IP
}
