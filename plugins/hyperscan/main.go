package main

import (
	"fmt"

	"github.com/flier/gohs/hyperscan"
	"github.com/k0kubun/pp"
	"gopkg.in/telegram-bot-api.v4"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
	"github.com/paper2code/golang-telegram-multibot/v2/pkg/models"

	"github.com/paper2code/golang-telegram-multibot/v2/plugins/hyperscan/cmd"
	hsm "github.com/paper2code/golang-telegram-multibot/v2/plugins/hyperscan/models"
)

var (
	flag   string
	cmdMap = map[string]*models.Command{
		"home": &models.Command{
			Func:        cmd.Home,
			Name:        "home",
			Keyboard:    "Home",
			Description: "Home",
		},
		"find": &models.Command{
			Func:        cmd.Find,
			Name:        "find",
			Keyboard:    "Find",
			Description: "Find patterns",
		},
	}
	flags, token string
	ctx          *context.MultiBotContext
	db           hyperscan.BlockDatabase
	options      map[string]interface{}
	scratch      hsm.Scratch
	regexMap     map[int]hsm.RegexLine
)

// InitPlugin initialize plugin if it needed
func InitPlugin(mbc *context.MultiBotContext) error {
	ctx = mbc
	options = ctx.GetOptions(GetName())
	if st, ok := options["flags"]; ok && st != nil {
		flags = st.(string)
	}
	if st, ok := options["token"]; ok && st != nil {
		token = st.(string)
	}
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
func GetCommands() (commands []string) {
	for cmd, _ := range cmdMap {
		commands = append(commands, cmd)
	}
	return
}

// GetKeyboard return plugin keyboard buttons for bot
func GetKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	length := len(cmdMap)
	buttons := make([][]tgbotapi.KeyboardButton, length)
	count := 0
	for _, info := range cmdMap {
		if info.Keyboard != "" {
			button := tgbotapi.NewKeyboardButton(info.Keyboard)
			buttons[count%3] = append(buttons[count%3], button)
			count++
		}
	}
	pp.Println("buttons", buttons)
	rkm := tgbotapi.NewReplyKeyboard(buttons...)
	rkm.OneTimeKeyboard = true
	rkm.Selective = true
	rkm.ResizeKeyboard = true
	return &rkm
}

// UpdateHandler function call for each update
func UpdateHandler(update tgbotapi.Update) (err error) {
	for _, info := range cmdMap {
		ctx.Log().Debugf("update.Message.Text: %s, info.Keyboard=%s", update.Message.Text, info.Keyboard)
		if update.Message.Text == info.Keyboard || update.Message.Text == fmt.Sprintf("/%s", GetName()) {
			info.Func.(func(int64, *context.MultiBotContext, string, *tgbotapi.ReplyKeyboardMarkup))(update.Message.Chat.ID, ctx, update.Message.Text, GetKeyboard())
		}
	}
	return nil
}

// RunCommand handler start if bot get one of commands
func RunCommand(command string, update tgbotapi.Update) (err error) {
	for cmd, info := range cmdMap {
		ctx.Log().Debugf("cmd: %s, match=%s", cmd, fmt.Sprintf("/%s_%s", GetName()))
		if update.Message.Text == fmt.Sprintf("/%s_%s", GetName(), cmd) {
			info.Func.(func(int64, *context.MultiBotContext, string, *tgbotapi.ReplyKeyboardMarkup))(update.Message.Chat.ID, ctx, "", GetKeyboard())
		}
	}
	return
}

// StartCommand handler start if bot get one command 'start'
func StartCommand(update tgbotapi.Update) (err error) {
	home, err := Asset("views/start.tmpl")
	if err != nil {
		return err
	}
	ctx.SendMessageMarkdown(update.Message.Chat.ID, string(home), 0, nil)
	return
}
