package telegram

import (
	"reminder-tg-bot/pkg/clients/telegram"
	"reminder-tg-bot/internal/storage"
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
