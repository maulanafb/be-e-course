package handler

import (
	"be_online_course/helper"
	"be_online_course/lesson"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type lessonHandler struct {
	service lesson.Service
}

func NewLessonHandler(service lesson.Service) *lessonHandler {
	return &lessonHandler{service}
}

func (h *lessonHandler) CreateLesson(c *fiber.Ctx) error {
	var input lesson.CreateLessonInput
	if err := c.BodyParser(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("Failed to create lesson", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}
	newLesson, err := h.service.CreateLesson(input)
	if err != nil {
		response := helper.APIResponse("Failed to create lesson", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("lesson created successfully", http.StatusOK, "success", newLesson)
	return c.Status(http.StatusOK).JSON(response)
}
