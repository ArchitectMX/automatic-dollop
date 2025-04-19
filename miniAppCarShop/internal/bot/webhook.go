package bot

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WebhookHandler(c *gin.Context) {
	var update tgbotapi.Update

	// Пробуем привязать JSON от Telegram в структуру update
	if err := c.ShouldBindJSON(&update); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	// Если есть сообщение, отправляем ответ
	if update.Message != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет 👋! Это Telegram Mini App.")
		bot.Send(msg)
	}

	// Статус успешного получения
	c.Status(http.StatusOK)
}
