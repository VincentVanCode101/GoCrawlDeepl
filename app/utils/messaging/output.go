package messaging

import (
	"crawl/deepl/utils/telegram"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// OutputTranslation sends the translation to Telegram and prints it to the console.
func OutputTranslation(telegramBot *telegram.TelegramBot, text string) {

	if telegramBot != nil {
		sendTelegramMessage(telegramBot, text)
	}

	fmt.Println("---------------------------------")
	fmt.Println(text)
	fmt.Println("---------------------------------")
}

func sendTelegramMessage(telegramBot *telegram.TelegramBot, text string) {
	msg := tgbotapi.NewMessage(telegramBot.ChatID, text)
	if _, err := telegramBot.Bot.Send(msg); err != nil {
		log.Printf("Failed to send message via Telegram bot: %v", err)
	}
}
