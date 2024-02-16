package category

import "errors"

type service struct {
	repository Repository
}

type Service interface {
	CreateCategory(input CreateCategoryInput) (Category, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCategory(input CreateCategoryInput) (Category, error) {
	category := Category{}
	category.Title = input.Title

	if input.Title == "" {
		return category, errors.New("Title cannot be empty")
	}

	return category, nil
}
