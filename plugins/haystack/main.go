package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"gopkg.in/telegram-bot-api.v4"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
	"github.com/paper2code/golang-telegram-multibot/v2/pkg/models"

	"github.com/paper2code/golang-telegram-multibot/v2/plugins/haystack/cmd"
)

var (
	ctx     *context.MultiBotContext
	options map[string]interface{}
	cmdMap  = map[string]*models.Command{
		"home": &models.Command{
			Func:        cmd.Home,
			Name:        "home",
			Keyboard:    "Home",
			Description: "Home",
		},
	}
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
	return "haystack"
}

// GetDescription function returns plugin description
func GetDescription() string {
	return "This is haystack plugin for multibot"
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
	buttons := make([][]tgbotapi.KeyboardButton, len(cmdMap))
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
			info.Func.(func(int64, *context.MultiBotContext, *models.Request, *tgbotapi.ReplyKeyboardMarkup))(update.Message.Chat.ID, ctx, &models.Request{Payload: update.Message.Text}, GetKeyboard())
		}
	}
	return nil
}

// RunCommand handler start if bot get one of commands
func RunCommand(command string, update tgbotapi.Update) (err error) {
	for cmd, info := range cmdMap {
		ctx.Log().Debugf("cmd: %s, match=%s", cmd, fmt.Sprintf("/%s_%s", GetName()))
		if update.Message.Text == fmt.Sprintf("/%s_%s", GetName(), cmd) {
			info.Func.(func(int64, *context.MultiBotContext, *models.Request, *tgbotapi.ReplyKeyboardMarkup))(update.Message.Chat.ID, ctx, &models.Request{}, GetKeyboard())
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
