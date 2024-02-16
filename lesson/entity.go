package lesson

import "time"

type Lesson struct {
	ID          int
	Title       string
	Content     string
	IsFree      bool
	MentorNote  string
	IsCompleted bool
	ChapterID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
