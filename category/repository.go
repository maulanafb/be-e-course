package category

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(category Category) (Category, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}
