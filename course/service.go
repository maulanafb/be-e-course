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
	GetAllCourses() ([]Course, error)
	SaveCourseImage(input CreateCourseImageInput, fileLocation string) (CourseImage, error)
	GetCourseBySlug(input string) (Course, error)
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

func (s *service) GetAllCourses() ([]Course, error) {
	courses, err := s.repository.FindAll()
	if err != nil {
		return courses, err
	}
	return courses, nil
}

func (s *service) SaveCourseImage(input CreateCourseImageInput, fileLocation string) (CourseImage, error) {
	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimary(int(input.CourseID))
		if err != nil {
			return CourseImage{}, err
		}
	}

	courseImage := CourseImage{}
	courseImage.CourseID = input.CourseID
	courseImage.IsPrimary = isPrimary
	courseImage.FileName = fileLocation

	newCourseImage, err := s.repository.CreateImage(courseImage)
	if err != nil {
		return newCourseImage, err
	}

	return newCourseImage, nil
}

func (s *service) GetCourseBySlug(input string) (Course, error) {
	course, err := s.repository.FindBySlug(input)
	if err != nil {
		return course, err
	}

	return course, nil
}
