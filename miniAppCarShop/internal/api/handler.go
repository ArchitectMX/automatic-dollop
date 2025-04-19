package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	// Пример получения данных пользователя
	initData := c.Query("initData")

	// Возвращаем данные с именем пользователя (можно адаптировать под свои нужды)
	c.JSON(http.StatusOK, gin.H{
		"username": "Иван", // Здесь можно подставить данные реального пользователя
	})
}
