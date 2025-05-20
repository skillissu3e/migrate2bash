package models

import (
	"fmt"
	"strings"
	"time"
)

// AuthRequest используется для входа/регистрации
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate проверяет корректность данных
func (r *AuthRequest) Validate() error {
	if !strings.Contains(r.Email, "@") {
		return fmt.Errorf("invalid email format")
	}
	if len(r.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	return nil
}

// AuthResponse возвращается после успешной аутентификации
type AuthResponse struct {
	Token string `json:"token"`
}

// User представляет пользователя в системе
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"` // "admin", "client"
	CreatedAt    time.Time `json:"created_at"`
}
