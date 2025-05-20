package main

import (
	"log"

	"github.com/streadway/amqp"
)

type NotificationTask struct {
	UserID  string `json:"user_id"` // Может быть email или Telegram chat ID
	Channel string `json:"channel"` // "email", "telegram"
	Message string `json:"message"`
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatal("Ошибка подключения к RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Ошибка создания канала:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("notification-tasks", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Ошибка объявления очереди:", err)
	}

	// Пример отправки задачи
	body := `{"user_id":"test@example.com","channel":"email","message":"Привет!"}`
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{ContentType: "application/json", Body: []byte(body)},
	)
	if err != nil {
		log.Println("Ошибка отправки в очередь:", err)
		return
	}

	log.Println("Сообщение отправлено в RabbitMQ")
}
