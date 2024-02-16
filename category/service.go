package category

import "errors"

type service struct {
	repository Repository
}

type Service interface {
	CreateCategory(input CreateCategoryInput) (Category, error)
	GetCategoryByTitle(title string) (Category, error)
	GetAllCategories() ([]Category, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCategory(input CreateCategoryInput) (Category, error) {
	// Validate input
	if input.Title == "" {
		return Category{}, errors.New("title cannot be empty")
	}
	// Create a new category instance
	newCategory := Category{
		Title: input.Title,
		// Additional fields if any...
	}
	createdCategory, err := s.repository.Save(newCategory)
	if err != nil {
		return Category{}, err
	}

	return createdCategory, nil
}
func (s *service) GetCategoryByTitle(title string) (Category, error) {
	category, err := s.repository.FindByTitle(title)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *service) GetAllCategories() ([]Category, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return categories, err
	}
	return categories, nil
}
