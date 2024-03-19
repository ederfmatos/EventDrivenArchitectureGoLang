package event

import "time"

type Event interface {
	GetId() string
	GetName() string
	GetDateTime() time.Time
}
