package storage

import (
	"google.golang.org/genproto/googleapis/type/datetime"
	"time"
)

type Event struct {
	Title      string
	date       datetime.DateTime
	notice     string
	userId     int
	timeBefore time.Time
}
