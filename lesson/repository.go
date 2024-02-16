package lesson

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(lesson Lesson) (Lesson, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(lesson Lesson) (Lesson, error) {
	err := r.db.Create(&lesson).Error
	if err != nil {
		return lesson, err
	}
	return lesson, nil
}
