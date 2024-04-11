package eventconsumer

import (
	"log"
	"time"

	"src/events"
)

type Consumer struct {
	fetcher events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) *Consumer {
	return &Consumer{
		fetcher: fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] fetcher: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Printf("[ERR] handle events: %s", err.Error())

			continue
		}
	}
}

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, ev := range events {
		log.Printf("got new event: %s", ev.Text)

		if err := c.processor.Process(ev); err != nil {
			log.Printf("[ERR] processor: %s", err.Error())

			continue
		}
	}

	return nil
}
