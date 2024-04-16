package telegram

import (
	"log"
	"src/clients/telegram"
	"strings"
)

func (p *Processor) doCallback(text string, meta *Meta) error {
	text = strings.TrimSpace(text)

	log.Printf("got new callback %s from %s\n", text, meta.UserName)

	switch meta.Data {
	case RndCmd:
		return p.sendRandom(meta.ChatID, meta.UserName)
	case HelpCmd:
		return p.sendHelp(meta.ChatID)
	case StartCmd:
		return p.sendHello(meta.ChatID, meta.UserName)
	case deleteUrlData:
		return p.deleteURL(text, meta)
	default:
		return p.tg.SendMessage(telegram.NewMessage(meta.ChatID, msgUnknownCommand, nil))
	}
}
