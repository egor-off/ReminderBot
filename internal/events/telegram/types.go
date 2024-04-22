package telegram

import (
	"ReminderBot/pkg/clients/telegram"
	"ReminderBot/internal/storage"
)

type Processor struct {
	tg *telegram.Client
	offset int
	storage storage.Storage
}

type Meta struct {
	ChatID int
	UserName string
	Data string
	MessageID int
	InlineMessageID string
}
