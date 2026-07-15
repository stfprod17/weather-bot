package main

import (
	"context"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stfprod17/weatherbot/clients/openweather"
	"github.com/stfprod17/weatherbot/handler"
	"github.com/stfprod17/weatherbot/repo"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to connect to database: %v\n", err)

	}
	defer conn.Close()

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Error ping db")
	}

	conn.QueryRow(context.Background(), "")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN")) //инициализация бота, которая принимает токен (нью апи)
	if err != nil {
		log.Panic(err) //если не смогли инициализировать бота, то логгируем ошибку и выходим из программы паникой
	}

	bot.Debug = true // выставляем режим дебаг для бота, чтобы отслеживать и контролировать все детальные логи. полезно, когда с ботом что-то не так и можно отследить в чем было конкретно дело

	log.Printf("Authorized on account %s", bot.Self.UserName) //логируем об успешной инициализации бота и передаем юзернейм бота

	owClient := openweather.New(os.Getenv("OPENWEATHERAPI_KEY"))

	userRepo := repo.New(conn)

	botHandler := handler.New(bot, owClient, userRepo)

	botHandler.Start()
}
