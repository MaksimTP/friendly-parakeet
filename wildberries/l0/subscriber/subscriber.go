package subscriber

import (
	"log"

	"github.com/nats-io/stan.go"
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
