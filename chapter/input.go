package chapter

type CreateChapterInput struct {
	Title    string `json:"title"`
	CourseID uint   `json:"course_id"`
}
