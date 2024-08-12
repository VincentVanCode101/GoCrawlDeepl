package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"crawl/deepl/utils/browser"
	"crawl/deepl/utils/cliargs"
	"crawl/deepl/utils/envutil"
	"crawl/deepl/utils/network"
	"crawl/deepl/utils/telegrambot"
	"crawl/deepl/utils/url"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const (
	baseURL = "https://www.deepl.com/en/translator#"
	xpath   = `//*[@id="textareasContainer"]/div[3]/section/div[1]/d-textarea/div/p/span`
)

// Translation holds the phrase to be translated along with its translations.
type Translation struct {
	toBeTranslatedPhrase string
	translatedPhrases    []string
}

// TelegramBot encapsulates the details needed to interact with the Telegram bot API.
type TelegramBot struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func main() {
	totalExecTime := time.Now()

	fromLang, toLang, err := envutil.GetLanguages()
	if err != nil {
		log.Fatalf("Failed to retrieve languages from environment: %v", err)
	}

	toBeTranslatedPhrases, err := cliargs.FilterNonEmptyArgs()
	if err != nil {
		log.Fatal(err)
	}

	if err := network.CheckInternetConnection(); err != nil {
		log.Fatalf("Failed to establish an internet connection: %v", err)
	}

	chromeCtx, cancelChrome, cancelExecAllocator, err := browser.GetChromeContext()
	if err != nil {
		log.Fatalf("Failed to initialize Chrome context: %v", err)
	}
	defer cancelExecAllocator()
	defer cancelChrome()

	bot, chatID, err := telegrambot.SetupTelegramBot()
	if err != nil {
		log.Printf("Failed to initialize Telegram bot: %v", err)
		return // the program should continue without the telegram-sending funcionality, thats why we just log and not abort here
	}

	var telegramBot *TelegramBot
	if bot != nil {
		telegramBot = &TelegramBot{bot: bot, chatID: chatID}
	}

	fmt.Printf("Amount of unrelated words to translate %d: %s\n", len(toBeTranslatedPhrases), strings.Join(toBeTranslatedPhrases, ", "))
	for _, toBeTranslatedPhrase := range toBeTranslatedPhrases {
		handleTranslation(chromeCtx, toBeTranslatedPhrase, fromLang, toLang, telegramBot)
	}

	fmt.Printf("Total execution time of the program: %v\n\n", time.Since(totalExecTime))
}

func handleTranslation(chromeCtx context.Context, toBeTranslatedPhrase string, fromLang string, toLang string, telegrambot *TelegramBot) {

	startTime := time.Now()

	translation := Translation{
		toBeTranslatedPhrase: toBeTranslatedPhrase,
		translatedPhrases:    []string{},
	}

	var deeplURL string = url.BuildDeeplURL(baseURL, fromLang, toLang, translation.toBeTranslatedPhrase)

	translatedPhrases, err := runTranslationTask(chromeCtx, deeplURL, xpath)
	if err != nil {
		log.Fatal(err)
	}
	translation.translatedPhrases = append(translation.translatedPhrases, translatedPhrases...)

	outputTranslation(translation, telegrambot)

	fmt.Printf("Execution time: %v\n\n", time.Since(startTime))
}

func formatTranslation(input string, translations []string) string {
	return fmt.Sprintf("Input:\n%s\n\nMain translation:\n%s", input, strings.Join(translations, "\n"))
}

func outputTranslation(translation Translation, telegramBot *TelegramBot) {
	text := formatTranslation(translation.toBeTranslatedPhrase, translation.translatedPhrases)

	// if telegramBot != nil {
	// 	msg := tgbotapi.NewMessage(telegramBot.chatID, text)
	// 	if _, err := telegramBot.bot.Send(msg); err != nil {
	// 		log.Printf("Failed to send message via Telegram bot: %v", err)
	// 	}
	// }
	fmt.Println("---------------------------------")
	fmt.Println(text)
	fmt.Println("---------------------------------")
}

func runTranslationTask(ctx context.Context, url, xpath string) ([]string, error) {
	var nodes []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(xpath),
		chromedp.Nodes(xpath, &nodes, chromedp.BySearch),
	); err != nil {
		return nil, err
	}
	return getTextFromNodes(ctx, nodes)
}

func getTextFromNodes(ctx context.Context, nodes []*cdp.Node) ([]string, error) {
	var translatedTexts []string
	for _, node := range nodes {
		var text string
		if err := chromedp.Run(ctx, chromedp.Text(node.FullXPath(), &text)); err != nil {
			return nil, err
		}
		translatedTexts = append(translatedTexts, text)
	}
	return translatedTexts, nil
}
