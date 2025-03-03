package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/telebot.v3"
	"go-chat-server/config"
	"go-chat-server/handlers"
)

func main() {
	// Load Environment Variables
	config.LoadEnv()

	// Connect Database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// Set up Telegram Bot
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Register Handlers
	handlers.RegisterHandlers(bot, db)

	fmt.Println("Bot is running...")
	bot.Start()
}