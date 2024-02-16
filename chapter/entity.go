package chapter

import (
	"be_online_course/lesson"
	"time"
)

type Chapter struct {
	ID        int
	Title     string
	CourseID  uint
	Lessons   []lesson.Lesson `gorm:"foreignKey:ChapterID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
