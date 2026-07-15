Weather Telegram Bot

A Telegram bot for getting current weather in any city around the world. Built with Go, OpenWeather API, and PostgreSQL for storing user data.


Bot Commands

1.CommandDescription/startWelcome message 
2.List of available commands:
"/city London" - save your city
"/weather" - current weather in your saved city


Tech Stack


1. Go — main programming language
2. PostgreSQL — database for storing user cities
3. Docker — running the database in a container
4. go-telegram-bot-api — Telegram Bot API library
5. OpenWeather API — weather data provider
6. pgx — PostgreSQL driver for Go
7. goose — database migrations
8. godotenv — loading environment variables



Where to get the keys


1. BOT_TOKEN — create a bot via @BotFather in Telegram
2. OPENWEATHERAPI_KEY — register at openweathermap.org and get a free API key
