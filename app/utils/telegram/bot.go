package telegram

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramBot encapsulates the details needed to interact with the Telegram bot API.
type TelegramBot struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

// SetupTelegramBot initializes and configures a new Telegram bot.
func SetupTelegramBot() (*TelegramBot, error) {
	botToken, isPresent := os.LookupEnv("BOT_TOKEN")
	if !isPresent || botToken == "" {
		log.Println("No environment variable for BOT_TOKEN found. Telegram bot will not be used.")
		return nil, nil
	}

	chatIDStr, isPresent := os.LookupEnv("CHAT_ID")
	if !isPresent || chatIDStr == "" {
		log.Println("No environment variable for CHAT_ID found. Telegram bot will not be used.")
		return nil, nil
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Printf("Error converting CHAT_ID to int64: %v", err)
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Printf("Failed to create bot: %v", err)
		return nil, err
	}

	return &TelegramBot{Bot: bot, ChatID: chatID}, nil
}
