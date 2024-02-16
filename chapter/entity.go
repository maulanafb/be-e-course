package chapter

import "time"

type Chapter struct {
	ID        int
	Title     string
	CourseID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
