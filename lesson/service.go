package lesson

type service struct {
	repository Repository
}

type Service interface {
	CreateLesson(input CreateLessonInput) (Lesson, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateLesson(input CreateLessonInput) (Lesson, error) {
	lesson := Lesson{}
	lesson.Title = input.Title
	lesson.Content = input.Content
	lesson.IsFree = input.IsFree
	lesson.MentorNote = input.MentorNote
	lesson.IsCompleted = input.IsCompleted
	lesson.ChapterID = input.ChapterID
	newLesson, err := s.repository.Save(lesson)
	if err != nil {
		return newLesson, err
	}
	return newLesson, nil

}
