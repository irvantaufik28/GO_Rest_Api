package task

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Doing       bool
	CreatedAt   time.Time
	UpdateAt    time.Time
}
