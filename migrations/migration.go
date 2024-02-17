package migrations

import (
	"time"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	type Transaction struct {
		ID         uint
		CourseID   uint `gorm:"column:course_id"`
		UserID     uint `gorm:"column:user_id"`
		Amount     int
		Status     string
		Code       string
		PaymentURL string
		CreatedAt  time.Time `gorm:"default:current_timestamp"`
		UpdatedAt  time.Time `gorm:"default:current_timestamp"`
	}

	type User struct {
		ID        uint
		Email     string `gorm:"unique"`
		Name      string `gorm:"nullable"`
		Image     string `gorm:"nullable"`
		Password  string
		Role      string    `gorm:"default:'user'"`
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}
	type Course struct {
		ID          uint
		Name        string
		Slug        string `gorm:"unique"`
		Thumbnail   string
		Price       int
		Level       string `gorm:"default:'BEGINNER'"`
		Description string
		MentorID    uint      `gorm:"column:mentor_id"`
		CategoryID  uint      `gorm:"column:category_id"`
		CreatedAt   time.Time `gorm:"default:current_timestamp"`
		UpdatedAt   time.Time `gorm:"default:current_timestamp"`
	}

	type CourseImage struct {
		ID        uint
		CourseID  uint `gorm:"column:course_id"`
		FileName  string
		IsPrimary int
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}

	type Category struct {
		ID        uint
		Title     string    `gorm:"unique"`
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}

	type Mentor struct {
		ID        uint
		Name      string
		Thumbnail string
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}

	type Chapter struct {
		ID        uint
		Title     string
		CourseID  uint      `gorm:"column:course_id"`
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}

	type Lesson struct {
		ID          uint
		Title       string
		Content     string
		IsFree      bool      `gorm:"default:false"`
		MentorNote  string    `gorm:"nullable"`
		IsCompleted bool      `gorm:"default:false"`
		ChapterID   uint      `gorm:"column:chapter_id"`
		CreatedAt   time.Time `gorm:"default:current_timestamp"`
		UpdatedAt   time.Time `gorm:"default:current_timestamp"`
	}

	// return db.AutoMigrate(&Transaction{}, &User{}, &Course{}, &Category{}, &Mentor{}, &Chapter{}, &Lesson{}, &CourseImage{})
	return db.AutoMigrate(&CourseImage{})
}
