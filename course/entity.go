package course

import (
	"be_online_course/category"
	"be_online_course/chapter"
	"be_online_course/user"
	"time"
)

type Course struct {
	ID          uint
	Name        string
	Slug        string
	Price       int
	Level       string
	Description string
	MentorID    uint
	CategoryID  uint
	Category    category.Category
	Chapter     []chapter.Chapter

	Mentor      user.User `gorm:"foreignKey:MentorID"`
	CourseImage []CourseImage
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CourseImage struct {
	ID        uint
	CourseID  int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
