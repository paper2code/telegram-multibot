package main

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telegram-bot-api.v4"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/models"
)

var bot *tgbotapi.BotAPI

var homeReplyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Search Paper2code"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Latest Trends"),
		tgbotapi.NewKeyboardButton("Latest Mashup"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Semantic Segementation"),
		tgbotapi.NewKeyboardButton("Summarize a text"),
		tgbotapi.NewKeyboardButton("Classifiy a text/url"),
	),
)

// homeReplyKeyboard.OneTimeKeyboard = true

// BotServe function run telegram bot listener
func BotServe() (err error) {

	var updates <-chan tgbotapi.Update
	defer wg.Done()

	updOptions := tgbotapi.NewUpdate(0)
	updOptions.Timeout = 60

	if updates, err = bot.GetUpdatesChan(updOptions); err != nil {
		return
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		go commandsHandler(update)
		go botEachUpdateHandler(update)

		if update.Message.Command() != "" {
			if update.Message.Command() == "start" {
				go botStartHandler(update)
			} else {
				go botCommandsHandler(update)
			}
		}
	}

	return
}

func botEachUpdateHandler(update tgbotapi.Update) {
	for name, botPlugin := range botPlugins {
		if err := botPlugin.EachUpdateHandler(update); err != nil {
			log.Errorf("Error in plugin %s on each update handler: %s", name, err)
		}
	}
}

func botCommandsHandler(update tgbotapi.Update) {
	cmd := update.Message.Command()
	if botPlugin, ok := botPluginsByCommand[cmd]; ok {
		shortCmd := strings.TrimPrefix(cmd, fmt.Sprintf("%s_", botPlugin.Name))
		if err := botPlugin.RunCommandHandler(shortCmd, update); err != nil {
			log.Errorf("Unable to run command '%s' for plugin '%s': %s", cmd, botPlugin.Name, err)
		}
	}
}

func botStartHandler(update tgbotapi.Update) {
	// startCommand(update)
	// botContext.SendMessageMarkdown(update.Message.Chat.ID, "*Hi!*", 0, nil)
	for name, botPlugin := range botPlugins {
		if err := botPlugin.StartCommandHandler(update); err != nil {
			log.Errorf("Unable to run command start for plugin '%s': %s", name, err)
		}
	}
}

// startCommand - Command
func startCommand(update tgbotapi.Update) {
	stories := []models.Message{
		models.Message{
			ChatID:   update.Message.Chat.ID,
			Duration: 1,
			Content:  "Hi 🙂!",
			MsgType:  "Message",
		},
		models.Message{
			ChatID:   update.Message.Chat.ID,
			Duration: 2,
			Content:  "My name is Seoz",
			MsgType:  "Message",
		},
		models.Message{
			ChatID:   update.Message.Chat.ID,
			Duration: 1,
			Content:  "How can I help you?",
			MsgType:  "Message",
			Keyboard: &homeReplyKeyboard,
		},
	}

	for _, story := range stories {
		models.ConsumeChainMessage(bot, story)
	}
}

func commandsHandler(update tgbotapi.Update) {
	command, _, ok := breakCommand(update.Message.Text)
	if ok {
		switch command {
		case "Back", "/Back", "/back":
			backCommand(update)
		// case "Start":
		// 	startCommand(update)
		case "Start", "/Start", "/start":
			startCommand(update)
			// case "/start":
			// 	startCommand(update)
			// case "Search Paper2code":
			// 	searchP2CCommand(update)
			// case "Latest Trends":
			// 	latestTrends(update)
			// case "Latest Mashup":
			// 	latestMashup(update)
			// case "Semantic Segementation":
			// 	semanticSegmentation(update)
			// case "Summarize a text":
			// 	summarizeText(update)
			// case "Classifiy a text/url":
			// 	classifierTextURL(update)
		}
	}
}

// backCommand - Command
func backCommand(update tgbotapi.Update) {
	stories := []models.Message{
		models.Message{
			ChatID:   update.Message.Chat.ID,
			Duration: 2,
			Content:  "What are you interested in?",
			MsgType:  "Message",
			Keyboard: &homeReplyKeyboard,
		},
	}
	for _, story := range stories {
		models.ConsumeChainMessage(bot, story)
	}
}

func breakCommand(message string) (string, []string, bool) {
	var command []string
	var arguments []string
	if message == "" {
		return "", arguments, false
	}
	command = strings.Split(message, " ")
	if len(command) >= 2 {
		arguments = strings.Split(command[1], ",")
	}
	return command[0], arguments, true
}
