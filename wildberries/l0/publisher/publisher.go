package publisher

import (
	"log"

	"github.com/nats-io/stan.go"
)

type Publisher struct {
	Sc stan.Conn
}

func (p *Publisher) Publish(channel string, data []byte) {
	err := p.Sc.Publish(channel, data)
	if err != nil {
		log.Println("error happened:", err.Error())
	}
}
