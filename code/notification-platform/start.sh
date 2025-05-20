#!/bin/bash

# 1. Проверка Docker
echo "Проверка Docker..."
if ! systemctl is-active --quiet docker; then
    echo "Запуск Docker..."
    sudo systemctl start docker || { echo "Не удалось запустить Docker"; exit 1; }
fi

# 2. Запуск RabbitMQ
echo "Запуск RabbitMQ..."
cd /home/dara/code/notification-platform
docker-compose down
docker-compose up -d

# Подождите, пока RabbitMQ инициализируется
sleep 10

# 3. Установка зависимостей
echo "Установка зависимостей..."
cd /home/dara/code/notification-platform
go mod tidy
go get github.com/rabbitmq/amqp091-go
go get github.com/go-telegram-bot-api/telegram-bot-api/v5
go get gopkg.in/gomail.v2

# 4. Создание директории для логов
mkdir -p /home/dara/code/notification-platform/logs

# 5. Запуск User Service
echo "Запуск User Service..."
cd /home/dara/code/notification-platform/services/user-service
go run main.go > /home/dara/code/notification-platform/logs/user-service.log 2>&1 &

# 6. Запуск API Gateway
echo "Запуск API Gateway..."
cd /home/dara/code/notification-platform/services/api-gateway
go run main.go > /home/dara/code/notification-platform/logs/api-gateway.log 2>&1 &

# 7. Запуск Delivery Service
echo "Запуск Delivery Service..."
cd /home/dara/code/notification-platform/services/delivery-service
go run main.go > /home/dara/code/notification-platform/logs/delivery.log 2>&1 &

echo "Все сервисы запущены!"
