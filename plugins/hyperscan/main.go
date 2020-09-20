package main

import (
	"gopkg.in/telegram-bot-api.v4"

	"github.com/flier/gohs/hyperscan"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
)

var ctx *context.MultiBotContext

// InitPlugin initialize plugin if it needed
func InitPlugin(mbc *context.MultiBotContext) error {
	ctx = mbc
	return nil
}

// GetName function returns plugin name
func GetName() string {
	return "hyperscan"
}

// GetDescription function returns plugin description
func GetDescription() string {
	return "Hyperscan module for parsing users' messages"
}

// GetCommands return plugin commands for bot
func GetCommands() []string {
	return []string{}
}

// GetKeyboard return plugin keyboard buttons for bot
func GetKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	return nil
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
