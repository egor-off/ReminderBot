package telegram

import (
	"log"
	"src/lib/e"
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

	switch text {
	case RndCmd:
	case HelpCmd:
	case StartCmd:
	default:
	}
}

func (p *Processor) svaePage(text string, chatID int, username string) (err error) {
	defer func () {err = e.WrapIfErr("cannot do command: save page", err)}()


}
