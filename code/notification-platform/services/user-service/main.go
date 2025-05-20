package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"notification-platform/services/user-service/handlers"
	"notification-platform/services/user-service/utils"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Роуты
	e.POST("/register", handlers.RegisterHandler)

	e.POST("/validate-token", func(c echo.Context) error {
		tokenString := c.FormValue("token")
		_, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}
		return c.JSON(http.StatusOK, map[string]string{"valid": "true"})
	})

	// Запуск сервера
	port := ":8081"
	e.Logger.Fatal(e.Start(port))
	fmt.Printf("User Service запущен на http://localhost%s\n", port)
}
