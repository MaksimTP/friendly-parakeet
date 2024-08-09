package main

import (
	"fmt"
	"log"
	"main/cache"
	"main/db"
	"main/model"

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

const (
	clusterID = "test-cluster"
	clientID  = "subscriber"
	url       = "nats://localhost:4222"
)

func NewSubscriber() *Subscriber {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url))
	if err != nil {
		log.Fatalf("Error connecting to NATS %v\n", err)
	}
	return &Subscriber{Sc: sc}
}

func main() {
	c := cache.Cache{}
	d := db.DataBase{}
	d.Connect("postgres", db.DbInfo)
	c.RestoreDataFromDB(d)

	sub := NewSubscriber()

	sub.Sc.Subscribe("order", func(m *stan.Msg) {
		fmt.Println("Received message:", string(m.Data))
		structData, err := model.ReadJSON(m.Data)
		if err != nil {
			log.Fatalf("error: %s", err.Error())
		}
		c.SaveData(structData)
		d.InsertData(structData)
	},
		stan.DeliverAllAvailable())

	select {}
}
