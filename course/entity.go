package course

import (
	"be_online_course/category"
	"be_online_course/chapter"
	"time"
)

type Course struct {
	ID          uint
	Name        string
	Slug        string
	Thumbnail   string
	Price       int
	Level       string
	Description string
	MentorID    uint
	CategoryID  uint
	Category    category.Category
	Chapter     chapter.Chapter
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
