package handler

import (
	"be_online_course/category"
	"be_online_course/helper"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type categoryHandler struct {
	service category.Service
}

func NewCategoryHandler(service category.Service) *categoryHandler {
	return &categoryHandler{service}
}

// api/v1/category

func (h *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	var input category.CreateCategoryInput
	if err := c.BodyParser(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	newCategory, err := h.service.CreateCategory(input)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("success to create campaign", http.StatusOK, "success", newCategory)
	return c.Status(http.StatusOK).JSON(response)
}
