package cmd

import (
	// "github.com/dghubble/sling"
	"gopkg.in/telegram-bot-api.v4"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
	"github.com/paper2code/golang-telegram-multibot/v2/pkg/models"
)

func Question(chatID int64, ctx *context.MultiBotContext, req *models.Request, rkm *tgbotapi.ReplyKeyboardMarkup) {
	ctx.Log().Infof("Payload: %s", req.Payload)
	askQuestion(req)
	ctx.SendMessageText(chatID, "Haystack Question", 0, rkm)
}

func askQuestion(req *models.Request) {

}
