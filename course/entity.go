package course

import "time"

type Course struct {
	ID          string
	Name        string
	Slug        string `gorm:"unique"`
	Thumbnail   string
	Price       int
	Level       string `gorm:"default:'BEGINNER'"`
	Description string
	MentorID    string `gorm:"column:mentorId"`
	CategoryID  string `gorm:"column:categoryId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
