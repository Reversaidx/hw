package storage

import (
	"google.golang.org/genproto/googleapis/type/datetime"
	"time"
)

type Event struct {
	ID          int
	Title       string
	date        datetime.DateTime
	notice      string
	user_id     int
	time_before time.Time
}
