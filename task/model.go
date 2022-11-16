package task

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	IsDoing     bool
	CreatedAt   time.Time
	UpdateAt    time.Time
}
