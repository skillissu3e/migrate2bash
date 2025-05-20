package main

import (
	"crypto/tls" // Добавлен для работы с TLS
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/streadway/amqp"
	"gopkg.in/gomail.v2"
)

// NotificationTask представляет задачу на отправку уведомления
type NotificationTask struct {
	UserID  string `json:"user_id"` // Может быть email или Telegram chat ID
	Channel string `json:"channel"` // "email", "telegram"
	Message string `json:"message"`
}

func main() {
	var conn *amqp.Connection
	var err error

	// Повторные попытки подключения к RabbitMQ
	for {
		conn, err = amqp.Dial("amqp://admin:admin@localhost:5672/")
		if err == nil {
			break
		}
		log.Println("Попытка повторного подключения к RabbitMQ через 5 секунд:", err)
		time.Sleep(5 * time.Second)
	}
	defer conn.Close()

	// Создание канала
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка создания канала: %v", err)
	}
	defer ch.Close()

	// Объявление очереди
	q, err := ch.QueueDeclare("notification-tasks", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Ошибка объявления очереди: %v", err)
	}

	// Получение сообщений
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Ошибка получения сообщений: %v", err)
	}

	// Обработка сообщений
	for msg := range msgs {
		var task NotificationTask
		if err := json.Unmarshal(msg.Body, &task); err != nil {
			log.Println("Ошибка разбора JSON:", err)
			continue
		}

		switch task.Channel {
		case "email":
			sendEmail(task.Message) // Теперь принимает только message
		case "telegram":
			sendTelegram(task.UserID, task.Message)
		default:
			log.Printf("Неизвестный канал: %s", task.Channel)
		}
	}
}

// sendEmail отправляет email через Yandex.Mail
func sendEmail(message string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "ki743v@yandex.ru")
	mailer.SetHeader("To", "re.kitaev@gmail.com")
	mailer.SetHeader("Subject", "Уведомление")

	// Текстовая версия (для клиентов без поддержки HTML)
	mailer.SetBody("text/plain", "Привет! Это текстовая версия уведомления.")

	// HTML-версия
	mailer.SetBody("text/html", "<html><body><h1>Уведомление</h1><p>"+message+"</p></body></html>")

	dialer := gomail.NewDialer("smtp.yandex.ru", 587, "ki743v@yandex.ru", "hcrhkeqgeviwmvxl")
	dialer.SSL = false
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Printf("Ошибка отправки Email: %v", err)
	} else {
		log.Println("Email успешно отправлен")
	}
}

// sendTelegram отправляет сообщение через Telegram-бота
func sendTelegram(userID, message string) {
	bot, err := tgbotapi.NewBotAPI("7787045212:AAHlQOfw_Dw1zzrGr_rz4-aUYE6hF9mOKk4")
	if err != nil {
		log.Printf("Ошибка инициализации Telegram-бота: %v", err)
		return
	}

	chatID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		log.Printf("Неверный формат Telegram chatID: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(chatID, message)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка отправки Telegram: %v", err)
	} else {
		log.Println("Telegram-сообщение отправлено")
	}
}
