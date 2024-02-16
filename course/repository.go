package course

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(course Course) (Course, error)
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
