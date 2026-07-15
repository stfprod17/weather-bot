package models

import "time"

type User struct {
	ID        int64
	City      string
	Username  string
	CreatedAt time.Time
}
