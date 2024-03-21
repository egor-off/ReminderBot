package main

import (
	"flag"
	"log"
	"src/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
	fetcher := fetcher.New()
	processor := processor.New(tgClient, )
}

func mustToken() string {
	var token string
	flag.StringVar(&token, "token-tg-bot", "", "input tg token for starting bot")
	flag.Parse()
	if token == "" {
		log.Fatal("token is empty")
	}
	return token
}
