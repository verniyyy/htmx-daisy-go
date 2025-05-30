package todo

import "time"

// TODO represents a task in the to-do list application.
type TODO struct {
	ID        int32
	Title     string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
