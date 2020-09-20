package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
	"github.com/paper2code/golang-telegram-multibot/v2/pkg/models"
	"gopkg.in/telegram-bot-api.v4"

	"github.com/paper2code/golang-telegram-multibot/v2/plugins/paper2code/cmd"
)

var (
	cmdMap = map[string]*models.Command{
		// "home":     &models.Command{Func: cmd.Home, Name: "home", Keyboard: "Paper2code Home"},
		"add":      &models.Command{Func: cmd.Add, Name: "add", Keyboard: "Add Paper/Repository"},
		"search":   &models.Command{Func: cmd.Search, Name: "search", Keyboard: "Search Paper"},
		"question": &models.Command{Func: cmd.Question, Name: "question", Keyboard: "Question ?"},
		"mashup":   &models.Command{Func: cmd.Mashup, Name: "mashup", Keyboard: "Mashup"},
		"trending": &models.Command{Func: cmd.Search, Name: "trending", Keyboard: "Trending"},
		"feeds":    &models.Command{Func: cmd.Feeds, Name: "feeds", Keyboard: "Feeds"},
	}
	ctx             *context.MultiBotContext
	options         map[string]interface{}
	endpoint, token string
)

// InitPlugin initialize plugin if it needed
func InitPlugin(mbc *context.MultiBotContext) error {
	ctx = mbc
	options = ctx.GetOptions(GetName())
	if st, ok := options["endpoint"]; ok && st != nil {
		endpoint = st.(string)
	}
	if st, ok := options["token"]; ok && st != nil {
		token = st.(string)
	}
	return nil
}

// GetName function returns plugin name
func GetName() string {
	return "paper2code"
}

// GetDescription function returns plugin description
func GetDescription() string {
	return "Paper2code Module"
}

// GetCommands return plugin commands for bot
func GetCommands() (commands []string) {
	for cmd, _ := range cmdMap {
		commands = append(commands, cmd)
	}
	return
}

// GetKeyboard return plugin keyboard buttons for bot
func GetKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	length := len(cmdMap) / 2
	buttons := make([][]tgbotapi.KeyboardButton, length)
	count := 0
	for _, info := range cmdMap {
		if info.Keyboard != "" {
			button := tgbotapi.NewKeyboardButton(info.Keyboard)
			buttons[count%3] = append(buttons[count%3], button)
			count++
		}
	}
	for i, button := range buttons {
		if len(button) == 0 {
			buttons[i] = nil
		}
	}
	pp.Println("buttons:", buttons)
	rkm := tgbotapi.NewReplyKeyboard(buttons...)
	rkm.OneTimeKeyboard = true
	rkm.Selective = true
	rkm.ResizeKeyboard = true
	return &rkm
}

// UpdateHandler function call for each update
func UpdateHandler(update tgbotapi.Update) (err error) {
	for _, info := range cmdMap {
		if update.Message.Text == info.Keyboard {
			info.Func.(func(int64, *context.MultiBotContext, string, *tgbotapi.ReplyKeyboardMarkup))(update.Message.Chat.ID, ctx, update.Message.Text, GetKeyboard())
		}
	}
	return nil
}

// RunCommand handler start if bot get one of commands
func RunCommand(command string, update tgbotapi.Update) (err error) {
	for cmd, info := range cmdMap {
		if update.Message.Text == fmt.Sprintf("/%s_%s", GetName(), cmd) {
			info.Func.(func(int64, *context.MultiBotContext, string, *tgbotapi.ReplyKeyboardMarkup))(update.Message.Chat.ID, ctx, "", GetKeyboard())
		}
	}
	return
}

// StartCommand handler start if bot get one command 'start'
func StartCommand(update tgbotapi.Update) (err error) {
	// pluginName := GetName()
	return
}
