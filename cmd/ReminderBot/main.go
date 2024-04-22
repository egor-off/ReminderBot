package main

import (
	"context"
	"flag"
	"log"
	tgClient "ReminderBot/pkg/clients/telegram"
	eventconsumer "ReminderBot/internal/consumer/eventConsumer"
	"ReminderBot/internal/events/telegram"
	storage "ReminderBot/internal/storage/sqlite"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
	pathDB = "./DB/tg_bot.db"
)

func main() {
	db, err := storage.New(pathDB)
	if err != nil {
		log.Fatalln("cannot start DB ", err)
	}

	db.Init(context.TODO())
	evetsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		db,
	)

	log.Println("service started")

	consumer := eventconsumer.New(evetsProcessor, evetsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped", err)
	}

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
