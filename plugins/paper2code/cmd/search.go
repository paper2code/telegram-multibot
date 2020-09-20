package cmd

import (
	"gopkg.in/telegram-bot-api.v4"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
)

func Search(chatID int64, ctx *context.MultiBotContext, input string, rkm *tgbotapi.ReplyKeyboardMarkup) {
	ctx.Log().Infof("input: %s", input)
	ctx.SendMessageText(chatID, "Search", 0, nil)
}
