package handler

import (
	"be_online_course/course"
	"be_online_course/helper"
	"be_online_course/transaction"
	"be_online_course/user"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service          transaction.Service
	courseRepository course.Repository
}

func NewTransactionHandler(service transaction.Service, courseRepository course.Repository) *TransactionHandler {
	return &TransactionHandler{service, courseRepository}
}

// func (h *TransactionHandler) GetCampaignTransactions(c *fiber.Ctx) error {
// 	var input transaction.GetCourseTransactionsInput

// 	// Get route parameters directly from c.Params
// 	campaignID, err := c.ParamsInt("id")
// 	if err != nil {
// 		response := helper.APIResponse("Invalid campaign ID", http.StatusBadRequest, "error", nil)
// 		return c.Status(http.StatusBadRequest).JSON(response)
// 	}

// 	currentUser := c.Locals("currentUser").(user.User)
// 	input.User = currentUser
// 	input.ID = campaignID

// 	transactions, err := h.service.GetTransactionsByCourseID(input)
// 	if err != nil {
// 		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
// 		return c.Status(http.StatusBadRequest).JSON(response)
// 	}

// 	response := helper.APIResponse("Campaign's transactions", http.StatusOK, "success", transactions)
// 	return c.Status(http.StatusOK).JSON(response)
// }

func (h *TransactionHandler) GetUserTransactions(c *fiber.Ctx) error {
	currentUser := c.Locals("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("User's transactions", http.StatusOK, "success", transactions)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var input transaction.CreateTransactionsInput

	param := c.Params("course_slug")
	inputCourse, err := h.courseRepository.FindBySlug(param)
	if err != nil {
		response := helper.APIResponse("Failed to create user's transactions", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	fmt.Println(inputCourse)
	currentUser := c.Locals("currentUser").(user.User)
	input.User = currentUser
	input.Course = inputCourse
	input.CourseID = int(inputCourse.ID)
	input.Amount = inputCourse.Price

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("Success to create transaction", http.StatusOK, "success", newTransaction)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *TransactionHandler) GetNotification(c *fiber.Ctx) error {
	var input transaction.TransactionNotificationInput

	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := h.service.ProcessPayment(input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(input)
}

func (h *TransactionHandler) GetPaidCourses(c *fiber.Ctx) error {
	currentUser := c.Locals("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetUserByPaidStatus(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's courses", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	fmt.Println(transactions)
	response := helper.APIResponse("User's courses ", http.StatusOK, "success", transactions)
	return c.Status(http.StatusOK).JSON(response)
}
