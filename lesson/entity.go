package lesson

import "time"

type Lesson struct {
	ID          uint
	Title       string
	Content     string
	IsFree      bool
	MentorNote  string
	IsCompleted bool
	ChapterID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
