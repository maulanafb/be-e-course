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

		response := helper.APIResponse("Failed to create category", http.StatusUnprocessableEntity, "error", errorMessage)
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

	// Retrieve the category title from URL parameters
	param := c.Params("title")

	// Set the category title in the input struct
	input.Title = param

	// Call the service to get the category by title
	newCategory, err := h.service.GetCategoryByTitle(input.Title)
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error retrieving category:", err)

		response := helper.APIResponse("Category Not Found", http.StatusInternalServerError, "error", nil)
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

		response := helper.APIResponse("Failed to create category", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}
	updatedCategory, err := h.service.UpdateCategory(inputID, inputData)
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error retrieving category:", err)

		response := helper.APIResponse("Category Not Found", http.StatusInternalServerError, "error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("success to update campaign", http.StatusOK, "success", category.FormatCategory(updatedCategory))
	return c.Status(http.StatusOK).JSON(response)
}
