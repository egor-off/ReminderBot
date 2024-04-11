package telegram

import (
	"errors"
	"src/clients/telegram"
	"src/events"
	"src/lib/e"
	"src/lib/storage"
)

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType = errors.New("unknown meta type")
)

type Processor struct {
	tg *telegram.Client
	offset int
	storage storage.Storage
}

type Meta struct {
	ChatID int
	UserName string
}

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg: client,
		storage: storage,
	}
}

// Fetch get updates from telegram and makes slice of events.
func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, v := range updates {
		res = append(res, event(v))
	}

	p.offset = updates[len(updates) - 1].ID + 1

	return res, nil
}

// Process handles certain event.
func (p *Processor) Process(event events.Event) error {
	switch event.Type {
		case events.Message:
			return p.processMessage(event)
		// case events.CallbackQuery:
		// 	return p.ProcessCallBackQuery(event)
		default:
			return e.Wrap("cannot process message", ErrUnknownEventType)
	}
}


func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

	if err := p.doCmd(event.Text, meta.ChatID, meta.UserName); err != nil {
		return e.Wrap("cannot doCmd", err)
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("can't get meta", ErrUnknownMetaType)
	}
	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID: upd.Message.Chat.ID,
			UserName: upd.Message.From.UserName,
		}
	} else if updType == events.CallbackQuery {
		res.Meta = Meta{
			ChatID: upd.CallbackData.Message.Chat.ID,
			UserName: upd.CallbackData.Message.From.UserName,
		}
	}

	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil && upd.CallbackData == nil {
		return ""
	} else if upd.Message != nil {
		return upd.Message.Text
	} else {
		return upd.CallbackData.Message.Text
	}
}

func fetchType(upd telegram.Update) events.Type{
	if upd.CallbackData != nil {
		return events.CallbackQuery
	} else if upd.Message != nil {
		return events.Message
	}
	return events.Unknown
}
