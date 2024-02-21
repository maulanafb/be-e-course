package handler

import (
	"be_online_course/course"
	"be_online_course/helper"
	"be_online_course/user"
	"fmt"
	"net/http"
	"time"

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
func (h *courseHandler) UploadImage(c *fiber.Ctx) error {
	var input course.CreateCourseImageInput

	if err := c.BodyParser(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("Failed to upload course images1", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	currentUser := c.Locals("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := fiber.Map{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload course image2", http.StatusBadRequest, "error", data)

		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Get the current date and format it as "2006-01-02" (Year-Month-Day)
	currentDate := time.Now().Format("2006-01-02")

	// Append the formatted date to the file name
	path := fmt.Sprintf("images/%d-%s-%s", userID, currentDate, file.Filename)

	err = c.SaveFile(file, path)
	if err != nil {
		data := fiber.Map{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload courses image3", http.StatusBadRequest, "error", data)

		return c.Status(http.StatusBadRequest).JSON(response)
	}

	_, err = h.service.SaveCourseImage(input, path)
	if err != nil {
		data := fiber.Map{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload course image4", http.StatusBadRequest, "error", data)

		return c.Status(http.StatusBadRequest).JSON(response)
	}

	data := fiber.Map{"is_uploaded": true}
	response := helper.APIResponse("courses image successfully uploaded", http.StatusOK, "success", data)

	return c.Status(http.StatusOK).JSON(response)
}

func (h *courseHandler) GetCourseBySlug(c *fiber.Ctx) error {
	var input course.GetCourseBySlugInput

	param := c.Params("slug")

	input.Slug = param

	newCourse, err := h.service.GetCourseBySlug(input.Slug)
	if err != nil {

		response := helper.APIResponse("Failed to upload course image4", http.StatusBadRequest, "error", nil)

		return c.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("courses retrieved successfully", http.StatusOK, "success", newCourse)
	return c.Status(http.StatusOK).JSON(response)
}
func (h *courseHandler) GetCourseDetailBySlug(c *fiber.Ctx) error {
	var input course.GetCourseBySlugInput

	param := c.Params("slug")

	input.Slug = param

	newCourse, err := h.service.GetCourseBySlug(input.Slug)
	if err != nil {

		response := helper.APIResponse("Failed to upload course image4", http.StatusBadRequest, "error", nil)

		return c.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("courses retrieved successfully", http.StatusOK, "success", newCourse)
	return c.Status(http.StatusOK).JSON(response)
}
