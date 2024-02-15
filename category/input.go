package category

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
	// Role     string `json:"role"`
	Image    string `json:"image" `
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
