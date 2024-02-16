package transaction

import (
	"be_online_course/course"
	"be_online_course/user"
	"time"
)

type Transaction struct {
	ID         int
	CourseID   int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Course     course.Course
	User       user.User
}
