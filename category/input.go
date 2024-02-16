package category

type CreateCategoryInput struct {
	Title string `json:"title" binding:"required"`
}

type GetCategoryTitle struct {
	Title string `uri:"title" binding:"required"`
}

type InputIDCategory struct {
	ID int `uri:"id"`
}

type InputDataCategory struct {
	Title string `json:"title" binding:"required"`
}
