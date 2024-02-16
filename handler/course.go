package handler

import (
	"be_online_course/course"
	"be_online_course/helper"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type courseHandler struct {
	service course.Service
}

func NewCourseHandler(service course.Service) *courseHandler {
	return &courseHandler{service}
}

func (h *courseHandler) CreateCourse(c *fiber.Ctx) error {
	var input course.CreateCourseInput

	if err := c.BodyParser(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("Failed to create courses", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	newCourse, err := h.service.CreateCourse(input)
	if err != nil {
		response := helper.APIResponse("Failed to create course", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("course created successfully", http.StatusOK, "success", newCourse)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *courseHandler) GetCourses(c *fiber.Ctx) error {
	courses, err := h.service.GetAllCourses()
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error retrieving courses:", err)

		response := helper.APIResponse("courses Not Found", http.StatusInternalServerError, "error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	// formattedCategories := course.FormatCategories(categories)
	response := helper.APIResponse("courses retrieved successfully", http.StatusOK, "success", courses)
	return c.Status(http.StatusOK).JSON(response)
}
