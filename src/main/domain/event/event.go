package event

import "time"

type Event interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}
