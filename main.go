package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Порт, на котором будет работать сервер
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Если порт не задан, используем 3000 по умолчанию
	}

	// Обработчик для главной страницы, который отдаёт HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Отправляем HTML-контент
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="theme-color" content="#ffffff">
    <title>Магазин мини-приложение</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        html, body {
            height: 100%;
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            flex-direction: column;
            background-color: #f4f4f4;
            color: #333;
        }

        .container {
            text-align: center;
            padding: 20px;
            max-width: 600px;
            width: 100%;
        }

        h1 {
            font-size: 2.5rem;
            margin-bottom: 20px;
        }

        p {
            font-size: 1.2rem;
            margin-bottom: 30px;
        }

        .products {
            display: flex;
            justify-content: space-around;
            margin-top: 20px;
        }

        .product {
            background-color: #fff;
            padding: 15px;
            border-radius: 10px;
            width: 30%;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        }

        .product h3 {
            font-size: 1.5rem;
            margin-bottom: 10px;
        }

        .product p {
            font-size: 1rem;
        }

        footer {
            margin-top: 30px;
        }

        button {
            background-color: #007BFF;
            color: white;
            font-size: 1.2rem;
            border: none;
            border-radius: 5px;
            padding: 10px 20px;
            cursor: pointer;
        }

        button:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Добро пожаловать в магазин!</h1>
        <p>Здесь вы можете найти все, что хотите. Наслаждайтесь покупками!</p>

        <div class="products">
            <div class="product">
                <h3>Товар 1</h3>
                <p>Описание товара 1</p>
            </div>
            <div class="product">
                <h3>Товар 2</h3>
                <p>Описание товара 2</p>
            </div>
            <div class="product">
                <h3>Товар 3</h3>
                <p>Описание товара 3</p>
            </div>
        </div>

        <footer>
            <button onclick="window.location.href='https://t.me/YOUR_BOT_USERNAME'">Перейти в чат с ботом</button>
        </footer>
    </div>
</body>
</html>`)
	})

	// Запуск веб-сервера
	fmt.Println("Запуск сервера на порту", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
