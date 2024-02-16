package category

type CreateCategoryInput struct {
	Title string `json:"title" binding:"required"`
}

type GetCategoryTitle struct {
	Title string `uri:"title" binding:"required"`
}
