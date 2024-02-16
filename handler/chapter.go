package handler

import (
	"be_online_course/chapter"

	"be_online_course/helper"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type chapterHandler struct {
	service chapter.Service
}

func NewChapterHandler(service chapter.Service) *chapterHandler {
	return &chapterHandler{service}
}

func (h *chapterHandler) CreateChapter(c *fiber.Ctx) error {
	var input chapter.CreateChapterInput

	if err := c.BodyParser(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("Failed to create chapter", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	newChapter, err := h.service.CreateChapter(input)
	if err != nil {
		response := helper.APIResponse("Failed to create chapter", http.StatusBadRequest, "error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("chapter created successfully", http.StatusOK, "success", newChapter)
	return c.Status(http.StatusOK).JSON(response)
}
