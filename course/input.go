package course

import "be_online_course/user"

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

type CreateCourseImageInput struct {
	CourseID  int  `form:"course_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
	User      user.User
}

type GetCourseBySlugInput struct {
	Slug string `uri:"slug" binding:"required"`
}
