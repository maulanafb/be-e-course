package main

import (
	"be_online_course/auth"
	"be_online_course/handler"
	"be_online_course/helper"
	"be_online_course/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/be-e-course?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Database connected successfully")

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	authService := auth.NewService()

	// Uncomment the line below to perform database migrations
	// if err := migrations.Migrate(db); err != nil {
	// 	log.Fatal(err)
	// }
	userHandler := handler.NewUserHandler(userService, authService)

	app := fiber.New()

	// Enable CORS
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Static("/images", "./images")

	api := app.Group("/api/v1")
	api.Post("/users", userHandler.RegisterUser)
	api.Post("/sessions", userHandler.Login)

	// Use the authMiddleware
	// api.Use(authMiddleware(authService, userService))

	app.Listen(":8088")
}

func authMiddleware(authService auth.Service, userService user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// Set the user in the context
		c.Locals("currentUser", user)
		return c.Next()
	}
}
