package course

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(course Course) (Course, error)
	FindAll() ([]Course, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(course Course) (Course, error) {
	err := r.db.Create(&course).Error
	if err != nil {
		return course, err
	}
	return course, nil
}

func (r *repository) FindAll() ([]Course, error) {
	var courses []Course
	err := r.db.Preload("Category").Preload("Chapter").Preload("Mentor").Preload("Chapter.Lessons").Find(&courses).Error
	if err != nil {
		return courses, err
	}
	return courses, nil
}
