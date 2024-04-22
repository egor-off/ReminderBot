package telegram

import (
	"context"
	"encoding/json"
	"log"
	"net/url"
	"reminder-tg-bot/pkg/clients/telegram"
	"reminder-tg-bot/pkg/e"
	"reminder-tg-bot/internal/storage"
	"strings"
)

const (
	StartCmd = "/start"
)

// TODO: rework commands with Meta

func (p *Processor) doMessage(text string, meta *Meta) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command %s from %s\n", text, meta.UserName)

	if isAddCmd(text) {
		return p.savePage(text, meta)
	}

	switch text {
	case StartCmd:
		return p.sendHello(meta)
	default:
		return p.unknownMessage(meta)
	}
}

func (p *Processor) savePage(pageURL string, meta *Meta) (err error) {
	defer func () {err = e.WrapIfErr("cannot do command: save page", err)}()

	page := &storage.Page{
		URL: pageURL,
		UserName: meta.UserName,
	}

	if err := p.tg.DeleteMessage(meta.ChatID, meta.MessageID); err != nil {
		return err
	}

	u, err := p.storage.PickUserInfo(context.TODO(), meta.UserName)
	if err != nil {
		return err
	}

	isExsists, err := p.storage.IsExistsPage(context.TODO(), page)
	if err != nil {
		return err
	}

	if isExsists {
		return p.tg.EditMessage(telegram.NewEditMessage(u.MessageID, u.ChatID, msgAllreadyExists , defaultKeyboard))
	}

	isRemoved, err := p.storage.IsRemovedURL(context.TODO(), page.URL, meta.UserName)
	if err != nil {
		return err
	}

	if isRemoved {
		err = p.storage.UpdateURLRemoved(context.TODO(), pageURL, page.UserName)
		if err != nil {
			return err
		}
	} else if err := p.storage.SavePage(context.TODO(), page); err != nil {
		return err
	}

	if err := p.tg.EditMessage(telegram.NewEditMessage(u.MessageID, u.ChatID, msgSaved, defaultKeyboard)); err != nil {
		return err
	}

	return nil
}

func (p *Processor) unknownMessage(meta *Meta) error {
	u, err := p.storage.PickUserInfo(context.TODO(), meta.UserName)
	if err != nil {
		return err
	}

	if err := p.tg.EditMessage(telegram.NewEditMessage(u.MessageID, meta.ChatID, msgUnknownCommand, defaultKeyboard)); err != nil {
		return e.Wrap("cannot send message", err)
	}
	return p.tg.DeleteMessage(meta.ChatID, meta.MessageID)
}

func (p *Processor) sendHello(meta *Meta) error {
	b, err := p.storage.IsExistsUser(context.TODO(), meta.UserName)
	if err != nil {
		return e.Wrap("cannot check if user exist: ", err)
	}
	if !b {
		if err := p.storage.SaveNewUser(context.TODO(), meta.UserName); err != nil {
			return e.Wrap("cannot save new user: ", err)
		}
	}
	r, err := p.tg.SendMessage(telegram.NewMessage(meta.ChatID, msgHello, defaultKeyboard))
	if err != nil {
		return e.Wrap("cannot send first message: ", err)
	}
	var rez telegram.MessageResponse
	if err := json.Unmarshal(r, &rez); err != nil {
		e.Wrap("cannot unmarshal mesage response: ", err)
	}

	if err := p.storage.UpdateUserInfo(context.TODO(), meta.UserName, rez.Result.MessageID, rez.Result.Chat.ID); err != nil {
		e.Wrap("cannot update info", err)
	}

	if err := p.tg.DeleteMessage(meta.ChatID, meta.MessageID); err != nil {
		return e.Wrap("cannot delete message: ", err)
	}

	return nil
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return (err == nil && u.Host != "")
}
