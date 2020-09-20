package cmd

import (
	"gopkg.in/telegram-bot-api.v4"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
)

func It(chatID int64, ctx *context.MultiBotContext, input string, rkm *tgbotapi.ReplyKeyboardMarkup) {
	ctx.Log().Infof("input: %s", input)
	ctx.SendMessageText(chatID, "Searx IT", 0, rkm)
}
