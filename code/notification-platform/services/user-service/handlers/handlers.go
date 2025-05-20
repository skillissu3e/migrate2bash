package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"notification-platform/services/user-service/models"
	"notification-platform/services/user-service/utils"
)

// RegisterHandler обрабатывает регистрацию пользователя
func RegisterHandler(c echo.Context) error {
	var req models.AuthRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	user := models.User{
		ID:           "1",
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         "client",
		CreatedAt:    time.Now(),
	}

	token, err := utils.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusCreated, models.AuthResponse{Token: token})
}
