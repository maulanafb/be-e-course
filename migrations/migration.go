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
		CreatedAt  time.Time `gorm:"default:current_timestamp"`
		UpdatedAt  time.Time `gorm:"default:current_timestamp"`
	}

	type User struct {
		ID        int
		Email     string `gorm:"unique"`
		Name      string `gorm:"nullable"`
		Image     string `gorm:"nullable"`
		Password  string
		Role      string    `gorm:"default:'user'"`
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}
	type Course struct {
		ID          int
		Name        string
		Slug        string `gorm:"unique"`
		Thumbnail   string
		Price       int
		Level       string `gorm:"default:'BEGINNER'"`
		Description string
		MentorID    string    `gorm:"column:mentorId"`
		CategoryID  string    `gorm:"column:categoryId"`
		CreatedAt   time.Time `gorm:"default:current_timestamp"`
		UpdatedAt   time.Time `gorm:"default:current_timestamp"`
	}

	type Category struct {
		ID        int
		Title     string    `gorm:"unique"`
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}

	type Mentor struct {
		ID        int
		Name      string
		Thumbnail string
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}

	type Chapter struct {
		ID        int
		Title     string
		CourseID  string    `gorm:"column:courseId"`
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}

	type Lesson struct {
		ID          int
		Title       string
		Content     string
		IsFree      bool      `gorm:"default:false"`
		MentorNote  string    `gorm:"nullable"`
		IsCompleted bool      `gorm:"default:false"`
		ChapterID   string    `gorm:"column:chapterId"`
		CreatedAt   time.Time `gorm:"default:current_timestamp"`
		UpdatedAt   time.Time `gorm:"default:current_timestamp"`
	}

	return db.AutoMigrate(&Transaction{}, &User{}, &Course{}, &Category{}, &Mentor{}, &Chapter{}, &Lesson{})
}
