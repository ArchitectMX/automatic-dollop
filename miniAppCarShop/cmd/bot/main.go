package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Получаем токен бота из переменной окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Указываем URL для вебхука, который предоставлен Glitch
	webhookURL := "https://your-project-name.glitch.me/webhook/" + token // Замените "your-project-name" на свой проект

	// Создаем экземпляр бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Устанавливаем вебхук для бота
	wh, _ := tgbotapi.NewWebhook(webhookURL)
	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	// Логируем информацию о вебхуке
	info, _ := bot.GetWebhookInfo()
	log.Println("Webhook set to:", info.URL)

	// Бот работает, держим программу в ожидании
	select {}
}
