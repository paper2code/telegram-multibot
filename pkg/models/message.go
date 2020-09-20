package models

import (
	"time"

	log "github.com/sirupsen/logrus"

	"gopkg.in/telegram-bot-api.v4"
)

type Messages map[string][]*Message

// Message - export
type Message struct {
	ChatID   int64                         `json:"chat_id" yaml:"chat-id"`
	MsgType  string                        `json:"msg_type" yaml:"msg-type"`
	Duration time.Duration                 `json:"duration" yaml:"duration"`
	Content  string                        `json:"content" yaml:"content"`
	Keyboard *tgbotapi.ReplyKeyboardMarkup `json:"-" yaml:"-"`
}

func ConsumeChainMessage(bot *tgbotapi.BotAPI, structure Message) {
	switch structure.MsgType {
	case "Message":
		var response tgbotapi.MessageConfig = tgbotapi.NewMessage(structure.ChatID, structure.Content)
		response.ParseMode = "Markdown"
		if structure.Keyboard != nil {
			response.ReplyMarkup = structure.Keyboard
		}
		sendMsg(bot, response)

	case "NewDocumentUpload":
		var response tgbotapi.DocumentConfig = tgbotapi.NewDocumentUpload(structure.ChatID, structure.Content)
		if structure.Keyboard != nil {
			response.ReplyMarkup = structure.Keyboard
		}
		sendMsg(bot, response)

	case "NewPhotoUpload":
		var response tgbotapi.PhotoConfig = tgbotapi.NewPhotoUpload(structure.ChatID, structure.Content)
		if structure.Keyboard != nil {
			response.ReplyMarkup = structure.Keyboard
		}
		sendMsg(bot, response)
	}
	time.Sleep(structure.Duration * time.Second)
}

// SendMsg - Send telegram message
func sendMsg(bot *tgbotapi.BotAPI, response tgbotapi.Chattable) {
	if _, err := bot.Send(response); err != nil {
		log.Panicln(err)
	}
}
