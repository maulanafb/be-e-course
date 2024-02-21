package course

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(course Course) (Course, error)
	FindAll() ([]Course, error)
	FindByID(ID int) (Course, error)
	MarkAllImagesAsNonPrimary(courseID int) (bool, error)
	FindBySlug(slug string) (Course, error)
	CreateImage(courseImage CourseImage) (CourseImage, error)
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
	err := r.db.Preload("Category").Preload("Chapter").Preload("Mentor").Preload("Chapter.Lessons").Preload("CourseImage").Find(&courses).Error
	if err != nil {
		return courses, err
	}
	return courses, nil
}

func (r *repository) FindByID(ID int) (Course, error) {
	var course Course
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&course).Error

	if err != nil {
		return course, err
	}
	return course, nil
}

func (r *repository) CreateImage(courseImage CourseImage) (CourseImage, error) {
	err := r.db.Create(&courseImage).Error
	if err != nil {
		return courseImage, err
	}

	return courseImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(courseID int) (bool, error) {
	err := r.db.Model(&CourseImage{}).Where("course_id = ?", courseID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) FindBySlug(slug string) (Course, error) {
	var course Course
	err := r.db.Preload("Category").Preload("Chapter").Preload("Mentor").Preload("Chapter.Lessons").Preload("CourseImage").Where("slug = ?", slug).Find(&course).Error
	if err != nil {
		return course, err
	}
	return course, nil
}
