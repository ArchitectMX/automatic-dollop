const express = require('express');
const app = express();

const PORT = process.env.PORT || 3000;

app.get('/', (req, res) => {
  res.send(`
    <!DOCTYPE html>
    <html lang="ru">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Магазин</title>
        <style>
            body {
                font-family: sans-serif;
                display: flex;
                flex-direction: column;
                align-items: center;
                justify-content: center;
                height: 100vh;
                margin: 0;
                background: #f0f0f0;
            }
            .container {
                background: white;
                padding: 2rem;
                border-radius: 10px;
                box-shadow: 0 0 10px rgba(0,0,0,0.1);
                text-align: center;
            }
            h1 {
                margin-bottom: 1rem;
            }
            .products {
                display: flex;
                gap: 1rem;
                margin: 1rem 0;
            }
            .product {
                padding: 1rem;
                background: #fafafa;
                border: 1px solid #ccc;
                border-radius: 5px;
                flex: 1;
            }
            button {
                padding: 0.6rem 1.2rem;
                font-size: 1rem;
                border: none;
                border-radius: 5px;
                background: #007bff;
                color: white;
                cursor: pointer;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Добро пожаловать!</h1>
            <p>Выберите товар ниже:</p>
            <div class="products">
                <div class="product">
                    <h3>Товар 1</h3>
                    <p>Описание</p>
                </div>
                <div class="product">
                    <h3>Товар 2</h3>
                    <p>Описание</p>
                </div>
                <div class="product">
                    <h3>Товар 3</h3>
                    <p>Описание</p>
                </div>
            </div>
            <button onclick="window.location.href='https://t.me/YOUR_BOT_USERNAME'">Написать боту</button>
        </div>
    </body>
    </html>
  `);
});

app.listen(PORT, () => {
  console.log(`🚀 Сервер запущен: http://localhost:${PORT}`);
});
