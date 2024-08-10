package subscriber

import (
	"log"

	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "subscriber"
	url       = "nats://localhost:4222"
)

type Subscriber struct {
	Sc  stan.Conn
	Sub stan.Subscription
}

func (s *Subscriber) Subscribe(channel string, callback stan.MsgHandler, opts ...stan.SubscriptionOption) {
	sub, err := s.Sc.Subscribe(channel, callback, opts...)
	if err != nil {
		log.Println("error happened:", err.Error())
	} else {
		s.Sub = sub
	}
}

func (s *Subscriber) Unsubscribe() {
	if s.Sub != nil {
		s.Sub.Close()
	}
}

func NewSubscriber() *Subscriber {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url))
	if err != nil {
		log.Fatalf("Error connecting to NATS %v\n", err)
	}
	return &Subscriber{Sc: sc}
}
