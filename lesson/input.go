package lesson

type CreateLessonInput struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsFree      bool   `json:"is_free"`
	MentorNote  string `json:"mentor_note"`
	IsCompleted bool   `json:"is_completed"`
	ChapterID   uint   `json:"chapter_id"`
}
