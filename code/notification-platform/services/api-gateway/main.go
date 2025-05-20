package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
)

var jwtKey = []byte("secret_key") // В продакшене используйте os.Getenv("JWT_SECRET")

// ValidateJWT проверяет токен
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})
}

// AuthMiddleware проверяет JWT и передает данные в контекст
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
		}

		// Удаление префикса "Bearer " (если есть)
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("email", claims["email"])
		c.Set("role", claims["role"])

		return next(c)
	}
}

// NotificationTask представляет задачу на отправку уведомления
type NotificationTask struct {
	UserID  string `json:"user_id"` // Может быть email или Telegram chat ID
	Channel string `json:"channel"` // "email", "telegram"
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Защищенные маршруты
	e.POST("/send", func(c echo.Context) error {
		var task NotificationTask
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Подключение к RabbitMQ
		conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
		if err != nil {
			log.Printf("Ошибка подключения к RabbitMQ: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "RabbitMQ unavailable"})
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Printf("Ошибка создания канала: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Channel creation failed"})
		}
		defer ch.Close()

		q, err := ch.QueueDeclare("notification-tasks", true, false, false, false, nil)
		if err != nil {
			log.Printf("Ошибка объявления очереди: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Queue declaration failed"})
		}

		body, _ := json.Marshal(task)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		if err != nil {
			log.Printf("Ошибка отправки в очередь: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to queue task"})
		}

		return c.JSON(http.StatusOK, map[string]string{"status": "queued"})
	}, AuthMiddleware)

	port := ":8080"
	e.Logger.Fatal(e.Start(port))
	fmt.Printf("API Gateway запущен на http://localhost%s\n", port)
}
