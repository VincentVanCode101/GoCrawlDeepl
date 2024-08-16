package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"crawl/deepl/utils/browser"
	"crawl/deepl/utils/cliargs"
	"crawl/deepl/utils/envutil"
	"crawl/deepl/utils/network"
	"crawl/deepl/utils/telegrambot"
	"crawl/deepl/utils/url"
)

const (
	baseURL                             = "https://www.deepl.com/en/translator#"
	mainTranslationXpath                = `//*[@id="textareasContainer"]/div[3]/section/div[1]/d-textarea/div/p/span`
	inputSpanForToBeTranslatedWordXpath = `//*[@id="textareasContainer"]/div[1]/section/div/div[1]/d-textarea/div[1]/p/span`
	typeOfToBeTranslatedWordXpath       = `//*[@id="headlessui-tabs-panel-:R6l8cqbspl:"]/div/div[2]/div/section/div/div/div/div/div/div[1]/span[1]/span[1]`
)

// Translation holds the phrase to be translated along with its translations.
type Translation struct {
	Phrase       string
	Translations map[string]string
}

// TelegramBot encapsulates the details needed to interact with the Telegram bot API.
type TelegramBot struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

func main() {
	totalExecTime := time.Now()

	fromLang, toLang := getLanguages()
	toBeTranslatedPhrases := getPhrasesToTranslate()

	checkInternetConnection()

	chromeCtx, cancelChrome := setupBrowser()
	defer cancelChrome()

	telegramBot := setupTelegramBot()

	processTranslations(chromeCtx, fromLang, toLang, toBeTranslatedPhrases, telegramBot)

	fmt.Printf("Total execution time of the program: %v\n\n", time.Since(totalExecTime))

	waitForBrowserClosure()
}

func getLanguages() (string, string) {
	fromLang, toLang, err := envutil.GetLanguages()
	if err != nil {
		log.Fatalf("Failed to retrieve languages from environment: %v", err)
	}
	return fromLang, toLang
}

func getPhrasesToTranslate() []string {
	toBeTranslatedPhrases, err := cliargs.FilterNonEmptyArgs()
	if err != nil {
		log.Fatalf("Failed to get phrases to translate: %v", err)
	}
	return toBeTranslatedPhrases
}

func checkInternetConnection() {
	if err := network.CheckInternetConnection(); err != nil {
		log.Fatalf("Failed to establish an internet connection: %v", err)
	}
}

func setupBrowser() (context.Context, func()) {
	keepBrowserOpen := os.Getenv("KEEP_BROWSER_OPEN") == "true"
	chromeCtx, cancelChrome, cancelExecAllocator, err := browser.GetChromeContext(!keepBrowserOpen)
	if err != nil {
		log.Fatalf("Failed to initialize Chrome context: %v", err)
	}

	cancelFunc := func() {
		if !keepBrowserOpen {
			cancelExecAllocator()
			cancelChrome()
		}
	}

	return chromeCtx, cancelFunc
}

func setupTelegramBot() *TelegramBot {
	bot, chatID, err := telegrambot.SetupTelegramBot()
	if err != nil {
		log.Printf("Failed to initialize Telegram bot: %v", err)
		return nil
	}

	return &TelegramBot{Bot: bot, ChatID: chatID}
}

func processTranslations(ctx context.Context, fromLang, toLang string, phrases []string, telegramBot *TelegramBot) {
	for _, phrase := range phrases {
		fetchType := !strings.ContainsAny(phrase, " \t")
		// When a word contains no whitespaces or tabs, we get the type of the word aswell,
		// otherwise we can't be sure if it is a compound noun (e.g. car door) (where fetching
		// the type would make sense) or if it is a noun phrase (e.g. flying banana ->
		// which consists of a verb (or participle) "flying", which is used as an adjective,
		// modifying the noun "banana"), where deepl doesnt show the type. And trying to fetch
		// the type of a noun phrase does not work (the crawler outputs random types which do
		// not represent acurate types)

		// Since I cant think of a better way to distinguish between a compound noun ('car door')
		// and a noun phrase ('flying banana') I only fetch the type if the word is a single word
		// accapting that deepl shows the type of "car door" which I do not fetch :(

		translation := translatePhrase(ctx, phrase, fromLang, toLang, fetchType)
		outputTranslation(translation, telegramBot)
	}
}

func translatePhrase(ctx context.Context, phrase, fromLang, toLang string, fetchType bool) Translation {
	startTime := time.Now()
	defer func() {
		fmt.Printf("Execution time for translation: %v\n\n", time.Since(startTime))
	}()

	translation := Translation{
		Phrase:       phrase,
		Translations: make(map[string]string),
	}

	deeplURL := url.BuildDeeplURL(baseURL, fromLang, toLang, phrase)

	translation.Translations = fetchTranslation(ctx, deeplURL, fetchType)

	return translation
}

func fetchTranslation(ctx context.Context, deeplURL string, fetchType bool) map[string]string {
	var mainTranslationNodes, typeOfWordNodes []*cdp.Node

	tasks := chromedp.Tasks{
		chromedp.Navigate(deeplURL),
		chromedp.WaitVisible(mainTranslationXpath),
		chromedp.Nodes(mainTranslationXpath, &mainTranslationNodes, chromedp.BySearch),
	}

	if fetchType {
		tasks = append(tasks,
			chromedp.Click(inputSpanForToBeTranslatedWordXpath, chromedp.NodeVisible),
			chromedp.Nodes(typeOfToBeTranslatedWordXpath, &typeOfWordNodes, chromedp.BySearch),
		)
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatalf("Failed to run translation task: %v", err)
	}

	translationResult := map[string]string{
		"mainTranslations": extractTextFromNodes(ctx, mainTranslationNodes),
	}

	if fetchType {
		translationResult["typeOfToBeTranslatedWord"] = extractTextFromNodes(ctx, typeOfWordNodes)
	}

	return translationResult
}

func extractTextFromNodes(ctx context.Context, nodes []*cdp.Node) string {
	var texts []string
	for _, node := range nodes {
		var text string
		if err := chromedp.Run(ctx, chromedp.Text(node.FullXPath(), &text)); err != nil {
			log.Printf("Failed to extract text from node: %v", err)
			continue
		}
		texts = append(texts, text)
	}
	return strings.Join(texts, " ")
}

func formatTranslation(phrase string, translations map[string]string) string {
	translationText := fmt.Sprintf("Input:\n%s\n\nMain translation:\n%s", phrase, translations["mainTranslations"])
	if typeOfWord, ok := translations["typeOfToBeTranslatedWord"]; ok {
		translationText = fmt.Sprintf("Input:\n%s (%s)\n\nMain translation:\n%s", phrase, typeOfWord, translations["mainTranslations"])
	}
	return translationText
}

func outputTranslation(translation Translation, telegramBot *TelegramBot) {
	text := formatTranslation(translation.Phrase, translation.Translations)

	if telegramBot != nil {
		sendTelegramMessage(telegramBot, text)
	}

	fmt.Println("---------------------------------")
	fmt.Println(text)
	fmt.Println("---------------------------------")
}

func sendTelegramMessage(telegramBot *TelegramBot, text string) {
	msg := tgbotapi.NewMessage(telegramBot.ChatID, text)
	if _, err := telegramBot.Bot.Send(msg); err != nil {
		log.Printf("Failed to send message via Telegram bot: %v", err)
	}
}

func waitForBrowserClosure() {
	if os.Getenv("KEEP_BROWSER_OPEN") == "true" {
		fmt.Println("Press 'Enter' to close the browser...")
		fmt.Scanln()
	}
}
