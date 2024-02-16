package chapter

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(chapter Chapter) (Chapter, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(chapter Chapter) (Chapter, error) {
	err := r.db.Create(&chapter).Error
	if err != nil {
		return chapter, err
	}
	return chapter, nil
}
