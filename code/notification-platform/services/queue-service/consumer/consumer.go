package main

import (
	"encoding/json"
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

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Ошибка получения сообщений:", err)
	}

	for msg := range msgs {
		var task NotificationTask
		if err := json.Unmarshal(msg.Body, &task); err != nil {
			log.Println("Ошибка разбора JSON:", err)
			continue
		}

		switch task.Channel {
		case "email":
			sendEmail(task.UserID, task.Message)
		case "telegram":
			sendTelegram(task.UserID, task.Message)
		default:
			log.Printf("Неизвестный канал: %s", task.Channel)
		}
	}
}
