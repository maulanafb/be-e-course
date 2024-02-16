package course

import (
	"time"

	"github.com/gosimple/slug"
)

type service struct {
	repository Repository
}

// Service interface defines the expected methods
type Service interface {
	CreateCourse(input CreateCourseInput) (Course, error)
}

// NewService creates a new instance of the service
func NewService(repository Repository) *service {
	return &service{repository}
}

// CreateCourse method in the service struct
func (s *service) CreateCourse(input CreateCourseInput) (Course, error) {
	course := Course{}
	course.Name = input.Name
	today := time.Now().Format("2006-01-02")
	course.Slug = slug.Make(input.Name + "-" + today)
	course.Price = input.Price
	course.Level = input.Level
	course.Description = input.Description
	course.MentorID = input.MentorID
	newCourse, err := s.repository.Save(course)
	if err != nil {
		return newCourse, err
	}
	return newCourse, nil
}
