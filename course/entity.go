package course

import "time"

type Course struct {
	ID          int
	Name        string
	Slug        string
	Thumbnail   string
	Price       int
	Level       string
	Description string
	MentorID    string
	CategoryID  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
