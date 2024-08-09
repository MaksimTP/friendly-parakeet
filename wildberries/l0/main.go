package main

import (
	"fmt"
	"html/template"
	"log"
	"main/model"
	"net/http"
)

// import (
// 	"log"

// 	"github.com/nats-io/stan.go"
// )

// const (
// 	clusterID = "test-cluster"
// 	clientID  = "subscriber"
// 	url       = "nats://localhost:4222"
// )

// type STANConn struct {
// 	Sc stan.Conn
// }

// func (s *STANConn) Close() {
// 	if s.Sc != nil {
// 		s.Sc.Close()
// 	}
// }

// func NewSTANConn() *STANConn {
// 	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url))
// 	if err != nil {
// 		log.Fatalf("Error connecting to NATS %v\n", err)
// 	}
// 	return &STANConn{sc}
// }
// func main() {

// 	// sc := NewSTANConn()

// 	// sc.Subscribe("order", func(m *stan.Msg) {
// 	// 	fmt.Println("Received message:", string(m.Data))
// 	// 	structData, err := model.ReadJSON(m.Data)
// 	// 	if err != nil {
// 	// 		log.Fatalf("error: %s", err.Error())
// 	// 	}
// 	// 	d.InsertData(structData)
// 	// },
// 	// 	stan.DeliverAllAvailable())

// 	// data, err := os.ReadFile("model.json")
// 	// if err != nil {
// 	// 	log.Fatalf("error: %s", err.Error())
// 	// }
// 	// sc.Publish("order", data)

// }

func main() {
	order := model.Order{}
	order.Id = 2
	order.OrderUid = "123"
	order.TrackNumber = "222"
	tmpl, err := template.ParseFiles("html_template/index.html")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("hello")
		tmpl.Execute(w, order)
	})
	http.ListenAndServe(":8000", nil)
}
