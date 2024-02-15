package migrations

import (
	"time"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
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
	}

	type User struct {
		ID        int
		Email     string `gorm:"unique"`
		Name      string `gorm:"nullable"`
		Image     string `gorm:"nullable"`
		Password  string
		Role      string `gorm:"default:'user'"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
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

	type Category struct {
		ID        string
		Title     string `gorm:"unique"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	type Mentor struct {
		ID        string
		Name      string
		Thumbnail string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	type Chapter struct {
		ID        string
		Title     string
		CourseID  string `gorm:"column:courseId"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	type Lesson struct {
		ID          string
		Title       string
		Content     string
		IsFree      bool   `gorm:"default:false"`
		MentorNote  string `gorm:"nullable"`
		IsCompleted bool   `gorm:"default:false"`
		ChapterID   string `gorm:"column:chapterId"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	return db.AutoMigrate(&Transaction{}, &User{}, &Course{}, &Category{}, &Mentor{}, &Chapter{}, &Lesson{})
}
