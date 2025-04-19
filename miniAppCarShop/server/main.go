package main

import (
	"os"
	"tg-miniapp/internal/api"
	"tg-miniapp/internal/bot"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация роутера
	router := gin.Default()

	// Получаем токен из переменных окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Обработка вебхука
	router.POST("/webhook/"+token, bot.WebhookHandler)

	// Статичный HTML для WebApp
	router.StaticFile("/", "./web/index.html")

	// API для получения данных пользователя
	router.GET("/api/user", api.UserHandler)

	// Запуск сервера
	router.Run(":8080")
}
