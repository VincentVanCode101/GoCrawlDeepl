package telegrambot

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SetupTelegramBot initializes and configures a new Telegram bot.
func SetupTelegramBot() (*tgbotapi.BotAPI, int64, error) {
	botToken, isPresent := os.LookupEnv("BOT_TOKEN")
	if !isPresent || botToken == "" {
		log.Println("No environment variable for BOT_TOKEN found. Telegram bot will not be used.")
		return nil, 0, nil
	}

	chatIDStr, isPresent := os.LookupEnv("CHAT_ID")
	if !isPresent || chatIDStr == "" {
		log.Println("No environment variable for CHAT_ID found. Telegram bot will not be used.")
		return nil, 0, nil
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		return nil, 0, fmt.Errorf("error converting CHAT_ID to int64: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create bot: %v", err)
	}

	return bot, chatID, nil
}
