package telegram

import (
	"src/clients/telegram"
	"src/events"
	"src/lib/e"
	"src/lib/storage"
	"log"
)

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

func event(upd telegram.Update) events.Event {
	printUpdate(&upd)
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	switch updType {
	case events.Message:
		res.Meta = Meta{
			ChatID: upd.Message.Chat.ID,
			UserName: upd.Message.From.UserName,
		}
	case events.CallbackQuery:
		res.Meta = Meta{
			ChatID: upd.CallbackData.Message.Chat.ID,
			UserName: upd.CallbackData.From.UserName,
			Data: upd.CallbackData.Data,
			ID: upd.CallbackData.ID,
		}
	}

	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil && upd.CallbackData == nil {
		return ""
	} else if upd.Message != nil {
		return upd.Message.Text
	} else if upd.CallbackData != nil {
		return upd.CallbackData.Message.Text
	}
	return ""
}

func fetchType(upd telegram.Update) events.Type{
	if upd.CallbackData != nil {
		return events.CallbackQuery
	} else if upd.Message != nil {
		return events.Message
	}
	return events.Unknown
}


func printUpdate(upd *telegram.Update) {
	if upd.Message != nil {
		log.Printf("update %v\nans: %v\n", upd, upd.Message)
	} else if upd.CallbackData != nil {
		log.Printf("update %v\nans: %v\nmessage: %v", upd, upd.CallbackData, upd.CallbackData.Message)
	}
}
