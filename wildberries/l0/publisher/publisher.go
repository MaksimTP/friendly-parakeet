package main

import (
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

type Publisher struct {
	Sc stan.Conn
}

const (
	clusterID = "test-cluster"
	clientID  = "publisher"
	url       = "nats://localhost:4222"
)

func NewPublisher() *Publisher {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url))
	if err != nil {
		log.Fatalf("Error connecting to NATS %v\n", err)
	}
	return &Publisher{sc}
}

func (p *Publisher) Publish(channel string, data []byte) {
	err := p.Sc.Publish(channel, data)
	if err != nil {
		log.Println("error happened:", err.Error())
	}
}

func main() {
	pub := NewPublisher()
	data, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	pub.Publish("order", data)
}
