package handler

import (
	"be_online_course/category"
	"be_online_course/helper"
	"fmt"
	"net/http"
	"strconv"

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

		response := helper.APIResponse("Failed to create categoryy", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	newCategory, err := h.service.CreateCategory(input)
	if err != nil {
		response := helper.APIResponse("Failed to create category", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("Category created successfully", http.StatusOK, "success", newCategory)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *categoryHandler) GetCategoryByTitle(c *fiber.Ctx) error {
	var input category.GetCategoryTitle

	param := c.Params("title")

	input.Title = param

	newCategory, err := h.service.GetCategoryByTitle(input.Title)
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error retrieving categoryy:", err)

		response := helper.APIResponse("Category Not Foundd", http.StatusInternalServerError, "error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Call the CategoryFormatter function to format the retrieved category
	formattedCategory := category.FormatCategory(newCategory)

	// Respond with the formatted category
	response := helper.APIResponse("Category retrieved successfully", http.StatusOK, "success", formattedCategory)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *categoryHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error retrieving category:", err)

		response := helper.APIResponse("Category Not Found", http.StatusInternalServerError, "error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	formattedCategories := category.FormatCategories(categories)
	response := helper.APIResponse("Category retrieved successfully", http.StatusOK, "success", formattedCategories)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *categoryHandler) UpdateCategory(c *fiber.Ctx) error {
	var inputID category.InputIDCategory

	param, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	inputID.ID = param

	var inputData category.InputDataCategory

	if err := c.BodyParser(&inputData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("Failed to update category", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}
	updatedCategory, err := h.service.UpdateCategory(inputID, inputData)
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error update category:", err)

		response := helper.APIResponse("Category Not Found", http.StatusInternalServerError, "error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("success to update campaign", http.StatusOK, "success", category.FormatCategory(updatedCategory))
	return c.Status(http.StatusOK).JSON(response)
}

func (h *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	// Retrieve the category ID from URL parameters
	param, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.APIResponse("Invalid category ID", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Call the service to delete the category
	deletedCategory, err := h.service.DeleteCategory(param)
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error deleting category:", err)

		response := helper.APIResponse("Failed to delete category", http.StatusInternalServerError, "error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Category deleted successfully", http.StatusOK, "success", deletedCategory)
	return c.Status(http.StatusOK).JSON(response)
}
