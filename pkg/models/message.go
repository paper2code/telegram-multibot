package models

import (
	"time"

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
