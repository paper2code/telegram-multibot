package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telegram-bot-api.v4"

	"github.com/paper2code/golang-telegram-multibot/v2/context"
)

var ctx *context.MultiBotContext

type Log struct {
	Timestamp time.Time
	Text      string
	*tgbotapi.Message
}

// InitPlugin initialize plugin if it needed
func InitPlugin(mbc *context.MultiBotContext) (err error) {
	ctx = mbc

	err = ctx.DBCreateTable(&Log{})
	return
}

// GetName function returns plugin name
func GetName() string {
	return "log_messages"
}

// GetDescription function returns plugin description
func GetDescription() string {
	return "Simple plugin save all messages to file for multibot"
}

// GetCommands return plugin commands for bot
func GetCommands() []string {
	return []string{}
}

// UpdateHandler function call for each update
func UpdateHandler(update tgbotapi.Update) (err error) {
	log.Debugf("%s", update.Message.Text)
	l := Log{time.Now(), update.Message.Text, update.Message}
	err = ctx.GetDB().Save(&l).Error
	return
}

// RunCommand handler start if bot get one of commands
func RunCommand(command string, update tgbotapi.Update) (err error) {
	return
}

// StartCommand handler start if bot get one command 'start'
func StartCommand(update tgbotapi.Update) (err error) {
	return
}
