package chapter

type service struct {
	repository Repository
}

type Service interface {
	CreateChapter(input CreateChapterInput) (Chapter, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateChapter(input CreateChapterInput) (Chapter, error) {
	chapter := Chapter{}
	chapter.Title = input.Title
	chapter.CourseID = input.CourseID

	newChapter, err := s.repository.Save(chapter)
	if err != nil {
		return newChapter, err
	}
	return newChapter, nil

}
