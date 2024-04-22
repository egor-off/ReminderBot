package telegram

import (
	"context"
	"errors"
	"log"
	"ReminderBot/pkg/clients/telegram"
	"ReminderBot/pkg/e"
	"ReminderBot/internal/storage"
)

func (p *Processor) doCallback(text string, meta *Meta) error {
	// text = strings.TrimSpace(text)

	log.Printf("got new callback %s from %s\n", meta.Data, meta.UserName)

	switch meta.Data {
	case randomData:
		return p.randomEdit(meta)
	case helpData:
		return p.editHelp(meta)
	case keepSaveData:
		return p.defaultEdit(keepSaveText + "\n\n", meta)
	case deleteUrlData:
		return p.deleteURL(text, meta)
	default:
		return p.tg.EditMessage(telegram.NewEditMessage(meta.MessageID ,meta.ChatID, msgUnknownCommand, defaultKeyboard))
	}
}

func (p *Processor) randomEdit(meta *Meta) (err error) {
	defer func ()  {
		err = e.WrapIfErr("cannot do cmd: sendRandom", err)
	}()

	page, err := p.storage.PickRandomPage(context.TODO(), meta.UserName)

	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return err
	} else if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.EditMessage(telegram.NewEditMessage(meta.MessageID, meta.ChatID, msgNoSavedURL, defaultKeyboard))
	}

	if err := p.tg.EditMessage(telegram.NewEditMessage(meta.MessageID, meta.ChatID, page.URL, afterRndKeyboard)); err != nil {
		return err
	}

	return nil
}

func (p *Processor) defaultEdit(text string, meta *Meta) error {
	if err := p.tg.EditMessage(telegram.NewEditMessage(meta.MessageID, meta.ChatID, text + msgDeafult, defaultKeyboard)); err != nil {
		return e.Wrap("cannot send default", err)
	}
	return nil
}

func (p *Processor) deleteURL(text string, meta *Meta) error {
	page := &storage.Page{
		URL: text,
		UserName: meta.UserName,
	}

	if err := p.storage.RemovePage(context.TODO(), page); err != nil {
		return e.Wrap("cannot delete page", err)
	}

	if err := p.tg.EditMessage(telegram.NewEditMessage(meta.MessageID, meta.ChatID, msgDeleted, defaultKeyboard)); err != nil {
		return e.Wrap("cannot edit message delete", err)
	}

	return nil
}

func (p *Processor) editHelp(meta *Meta) error {
	if err := p.tg.EditMessage(telegram.NewEditMessage(meta.MessageID, meta.ChatID, msgHelp, defaultKeyboard)); err != nil {
		return e.Wrap("cannot send default", err)
	}
	return nil
}
