package telegram

import (
	"context"
	"errors"
	"log"
	"net/url"
	"src/clients/telegram"
	"src/lib/e"
	"src/lib/storage"
	"strings"
)

const (
	RndCmd = "/rnd"
	HelpCmd = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command %s from %s\n", text, username)

	if isAddCmd(text) {
		return p.savePage(text, chatID, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID, username)
	default:
		return p.tg.SendMessage(telegram.NewMessage(chatID, msgUnknownCommand, nil))
	}
}

func (p *Processor) savePage(pageURL string, chatID int, username string) (err error) {
	defer func () {err = e.WrapIfErr("cannot do command: save page", err)}()

	page := &storage.Page{
		URL: pageURL,
		UserName: username,
	}

	isExsists, err := p.storage.IsExistsPage(context.TODO(), page)
	if err != nil {
		return err
	}

	if isExsists {
		return p.tg.SendMessage(telegram.NewMessage(chatID, msgAllreadyExists, nil))
	}

	if err := p.storage.SavePage(context.TODO(), page); err != nil {
		return err
	}

	if err := p.tg.SendMessage(telegram.NewMessage(chatID, msgSaved, nil)); err != nil {
		return err
	}

	return nil
}

func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func ()  {
		err = e.WrapIfErr("cannot do cmd: sendRandom", err)
	}()

		page, err := p.storage.PickRandomPage(context.TODO(), username)

		if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
			return err
		} else if errors.Is(err, storage.ErrNoSavedPages) {
			return p.tg.SendMessage(telegram.NewMessage(chatID, msgNoSavedURL, nil))
		}

		if err := p.tg.SendMessage(telegram.NewMessage(chatID, page.URL, nil)); err != nil {
			return err
		}

	return p.storage.RemovePage(context.TODO(), page)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(telegram.NewMessage(chatID, msgHelp, nil))
}

func (p *Processor) sendHello(chatID int, username string) error {
	b, err := p.storage.IsExistsUser(context.TODO(), username)
	if err != nil {
		return e.Wrap("cannot check if user exist: ", err)
	}
	if !b {
		if err := p.storage.SaveNewUser(context.TODO(), username); err != nil {
			return e.Wrap("cannot save new user: ", err)
		}
	}
	return p.tg.SendMessage(telegram.NewMessage(chatID, msgHello, nil))
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return (err == nil && u.Host != "")
}
