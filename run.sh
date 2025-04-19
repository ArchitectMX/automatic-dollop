#!/bin/bash

# Убедимся, что Go установлен
if ! command -v go &> /dev/null
then
    echo "Go не найден, установите Go"
    exit
fi

# Установим зависимости, если они есть
go mod tidy

# Запуск приложения
go run main.go
