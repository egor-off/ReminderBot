package telegram

import (
	"errors"
	"src/events"
	"src/lib/e"
)

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType = errors.New("unknown meta type")
)

// Process handles certain event.
func (p *Processor) Process(event events.Event) error {
	if event.Type != events.Unknown {
		return p.processEvent(event)
	} else {
		return e.Wrap("cannot process mesage", ErrUnknownEventType)
	}
}


func (p *Processor) processEvent(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

	switch event.Type {
	case events.Message:
		if err := p.doMessage(event.Text, meta.ChatID, meta.UserName); err != nil {
			return e.Wrap("cannot doCmd", err)
		}
	case events.CallbackQuery:
		if err := p.doCallback(event.Text, &meta); err != nil {
			return e.Wrap("cannot doCallback", err)
		}
	default:
		return nil
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("wrong meta", ErrUnknownMetaType)
	}
	return res, nil
}


