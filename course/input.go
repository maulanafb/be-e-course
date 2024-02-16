package course

type CreateCourseInput struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Thumbnail   string `json:"thumbnail"`
	Price       int    `json:"price"`
	Level       string `json:"level"`
	Description string `json:"description"`
	MentorID    uint   `json:"mentor_id"`
	CategoryID  uint   `json:"category_id"`
}
