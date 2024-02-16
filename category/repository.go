package category

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(category Category) (Category, error)
	FindByTitle(title string) (Category, error)
	FindByID(ID int) (Category, error)
	FindAll() ([]Category, error)
	Update(category Category) (Category, error)
	Delete(ID int) (Category, error)
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

func (r *repository) FindByTitle(title string) (Category, error) {
	var category Category
	err := r.db.Where("title = ?", title).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) FindByID(title int) (Category, error) {
	var category Category
	err := r.db.Where("id = ?", title).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (r *repository) Update(category Category) (Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) Delete(ID int) (Category, error) {
	var category Category

	err := r.db.Where("id = ?", ID).Delete(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
