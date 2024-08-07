package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	// d := &db.DataBase{}
	// d.Connect("postgres", db.DbInfo)
	// d.Close()
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Publish("foo", []byte("Hello World!"))
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})
	nc.Drain()
}
