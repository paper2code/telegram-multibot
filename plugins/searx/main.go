package main

import (
	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
	"gopkg.in/telegram-bot-api.v4"
	// "github.com/paper2code/golang-telegram-multibot/v2/plugins/searx/models"
)

var ctx *context.MultiBotContext

// InitPlugin initialize plugin if it needed
func InitPlugin(mbc *context.MultiBotContext) error {
	ctx = mbc
	return nil
}

// GetName function returns plugin name
func GetName() string {
	return "searx"
}

// GetDescription function returns plugin description
func GetDescription() string {
	return "Searx is search engine aggregator"
}

// GetCommands return plugin commands for bot
func GetCommands() []string {
	return []string{}
}

// GetKeyboard return plugin keyboard buttons for bot
func GetKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("General"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("News"),
			tgbotapi.NewKeyboardButton("Social"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Technology"),
			tgbotapi.NewKeyboardButton("Science"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Images"),
			tgbotapi.NewKeyboardButton("Videos"),
			tgbotapi.NewKeyboardButton("Music"),
		),
	)
	return &replyKeyboard
}

// UpdateHandler function call for each update
func UpdateHandler(update tgbotapi.Update) (err error) {
	return nil
}

// RunCommand handler start if bot get one of commands
func RunCommand(command string, update tgbotapi.Update) (err error) {
	return
}

// StartCommand handler start if bot get one command 'start'
func StartCommand(update tgbotapi.Update) (err error) {
	return
}

func searchSearx() {
}
