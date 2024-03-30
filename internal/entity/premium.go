package entity

import "time"

type Premium struct {
	UserID     string
	ActiveTime time.Time
}
