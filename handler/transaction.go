package handler

import (
	"be_online_course/helper"
	"be_online_course/transaction"
	"be_online_course/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *TransactionHandler {
	return &TransactionHandler{service}
}

// func (h *TransactionHandler) GetCampaignTransactions(c *fiber.Ctx) error {
// 	var input transaction.GetCourseTransactionsInput

// 	// Get route parameters directly from c.Params
// 	campaignID, err := c.ParamsInt("id")
// 	if err != nil {
// 		response := helper.APIResponse("Invalid campaign ID", http.StatusBadRequest, "error", nil)
// 		return c.Status(http.StatusBadRequest).JSON(response)
// 	}

// 	currentUser := c.Locals("currentUser").(*user.User)
// 	input.User = *currentUser
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
	currentUser := c.Locals("currentUser").(*user.User)
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

	if err := c.BodyParser(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	currentUser := c.Locals("currentUser").(*user.User)
	input.User = *currentUser

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
