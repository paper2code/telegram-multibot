package cmd

import (
	"gopkg.in/telegram-bot-api.v4"

	// 	"github.com/lunny/html2md"
	// 	"github.com/pkg/errors"
	// 	"github.com/mmcdole/gofeed"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
)

func Feeds(chatID int64, ctx *context.MultiBotContext, input string, rkm *tgbotapi.ReplyKeyboardMarkup) {
	ctx.Log().Infof("input: %s", input)
	ctx.SendMessageText(chatID, "Feeds", 0, nil)
}
