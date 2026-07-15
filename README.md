Weather Telegram Bot

A Telegram bot for getting current weather in any city around the world. Built with Go, OpenWeather API, and PostgreSQL for storing user data.


This project was built as part of learning Go. The /start command with a welcome message was independently added on top of the tutorial.




Bot Commands

CommandDescription/startWelcome message and list of available commands/city LondonSave your city/weatherGet current weather in your saved city


Tech Stack


Go — main programming language
PostgreSQL — database for storing user cities
Docker — running the database in a container
go-telegram-bot-api — Telegram Bot API library
OpenWeather API — weather data provider
pgx — PostgreSQL driver for Go
goose — database migrations
godotenv — loading environment variables



Where to get the keys


BOT_TOKEN — create a bot via @BotFather in Telegram
OPENWEATHERAPI_KEY — register at openweathermap.org and get a free API key
