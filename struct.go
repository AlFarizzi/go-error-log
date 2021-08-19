package errjson

import "time"

type Error struct {
	Time  time.Time
	Error string
}
