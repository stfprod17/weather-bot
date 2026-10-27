# Weather Bot

A Telegram bot for getting the current weather in any city around the world.

The bot allows users to save their city and receive the current temperature using the OpenWeather API.

## Features

- Telegram bot integration
- Search weather by city
- Save a personal city for each user
- Get the current temperature
- PostgreSQL storage
- Database migrations with Goose
- Docker-based PostgreSQL setup
- User data persistence

## Commands

| Command | Description |
|---|---|
| `/start` | Show a welcome message |
| `/city <city>` | Save or update your city |
| `/weather` | Get the current weather in your saved city |

Examples:

```text
/city London
/weather
```

## Tech Stack

- Go
- PostgreSQL
- Docker
- Telegram Bot API
- OpenWeather API
- pgx
- Goose
- godotenv

## Requirements

- Go 1.25+
- Docker Desktop
- Goose
- Telegram account
- OpenWeather API key

## Configuration

Create a `.env` file:

```bash
cp .env.example .env
```

Configure the following variables:

```env
BOT_TOKEN=your_telegram_bot_token
OPENWEATHERAPI_KEY=your_openweather_api_key

DATABASE_URL=postgres://postgres:your_password@localhost:5432/postgres

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://postgres:your_password@localhost:5432/postgres
GOOSE_MIGRATION_DIR=./migrations
```

Never commit `.env` to the repository.

## Database Setup

Start PostgreSQL:

```bash
docker compose up -d weather-postgres
```

Install Goose if necessary:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Run database migrations:

```bash
goose postgres "$GOOSE_DBSTRING" -dir migrations up
```

Check migration status:

```bash
goose postgres "$GOOSE_DBSTRING" -dir migrations status
```

## Run the Bot

Install dependencies:

```bash
go mod download
```

Start the bot:

```bash
go run main.go
```

Open the bot in Telegram and send:

```text
/start
```

## Project Structure

```text
weather-bot/
├── clients/
│   └── openweather/
├── handler/
├── migrations/
├── models/
├── repo/
├── .env.example
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

## How It Works

1. The user sends a command to the Telegram bot.
2. The bot saves the user's city in PostgreSQL.
3. The bot requests the city coordinates from OpenWeather.
4. The bot receives the current weather data.
5. The bot sends the temperature back to the user.

## Database

The application stores:

- Telegram user ID
- Username
- Saved city
- Account creation time

Database migrations are located in the `migrations` directory.

## Development

Stop PostgreSQL:

```bash
docker compose down
```

View PostgreSQL logs:

```bash
docker compose logs -f weather-postgres
```

## Future Improvements

- Add weather description
- Add feels-like temperature
- Add humidity and wind speed
- Add a multi-day forecast
- Add inline Telegram buttons
- Add automatic weather updates
- Add city autocomplete
- Add automated tests
- Improve error handling and logging

## License

This project is intended for educational and personal use.
