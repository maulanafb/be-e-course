package transaction

import (
	"be_online_course/course"
	"be_online_course/user"
)

type GetCourseTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionsInput struct {
	Amount   int `json:"amount" binding:"required"`
	CourseID int `json:"course_id" binding:"required"`
	User     user.User
	Course   course.Course
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id" `
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

type GetUserIDInput struct {
	UserID int `uri:"user_id"`
}
