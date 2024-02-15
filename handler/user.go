package handler

import (
	"be_online_course/auth"
	"be_online_course/helper"
	"be_online_course/user"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *UserHandler {
	return &UserHandler{userService, authService}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var input user.RegisterUserInput
	if err := c.BodyParser(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}
		response := helper.FiberAPIResponse(c, "Register account failedd", http.StatusUnprocessableEntity, "error", errorMessage)
		return response
	}
	checkEmail, err := h.userService.CheckingEmail(input.Email)
	fmt.Println(checkEmail, err)
	if err != nil {
		return err
	}
	if checkEmail {
		response := helper.FiberAPIResponse(c, "Email Already Registered", http.StatusBadRequest, "error", nil)
		return response
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.FiberAPIResponse(c, "Register account failed", http.StatusBadRequest, "error", nil)
		return response
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.FiberAPIResponse(c, "Register account failed", http.StatusBadRequest, "error", nil)
		return response
	}
	formatter := user.FormatUser(newUser, token)
	response := helper.FiberAPIResponse(c, "Account successfully registered", http.StatusOK, "success", formatter)
	return response
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var input user.LoginInput
	if err := c.BodyParser(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}
		response := helper.FiberAPIResponse(c, "Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		return response
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}
		response := helper.FiberAPIResponse(c, "Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		return response
	}
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.FiberAPIResponse(c, "Login failed", http.StatusBadRequest, "error", nil)
		return response
	}
	formatter := user.FormatUser(loggedinUser, token)
	response := helper.FiberAPIResponse(c, "Login success", http.StatusOK, "success", formatter)
	return response
}

func (h *UserHandler) CheckEmailAvailability(c *fiber.Ctx) error {
	var input user.CheckEmailInput
	if err := c.BodyParser(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}
		response := helper.FiberAPIResponse(c, "Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		return response
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := fiber.Map{"errors": "Server error"}
		response := helper.FiberAPIResponse(c, "Email Already Registered", http.StatusUnprocessableEntity, "error", errorMessage)
		return response
	}

	data := fiber.Map{"is_available": isEmailAvailable}
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.FiberAPIResponse(c, metaMessage, http.StatusOK, "success", data)
	return response
}

func (h *UserHandler) UploadAvatar(c *fiber.Ctx) error {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := fiber.Map{"is_uploaded": false}
		response := helper.FiberAPIResponse(c, "Failed to upload avatars", http.StatusBadRequest, "error", data)
		return response
	}

	currentUser := c.Locals("currentUser").(user.User)
	userID := currentUser.ID
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	path := fmt.Sprintf("images/%d-%d-%s", userID, currentTime, file.Filename)

	err = c.SaveFile(file, path)
	if err != nil {
		data := fiber.Map{"is_uploaded": false}
		response := helper.FiberAPIResponse(c, "Failed to upload avatar", http.StatusBadRequest, "error", data)
		return response
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := fiber.Map{"is_uploaded": false}
		response := helper.FiberAPIResponse(c, "Failed to upload avatar", http.StatusBadRequest, "error", data)
		return response
	}
	data := fiber.Map{"is_uploaded": true}
	response := helper.FiberAPIResponse(c, "success upload avatar", http.StatusOK, "success", data)
	return response
}
