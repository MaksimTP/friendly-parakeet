package main

import (
	"fmt"
	"html/template"
	"log"
	"main/cache"
	"main/db"
	"main/model"
	"main/subscriber"
	"net/http"

	"github.com/nats-io/stan.go"
)

func main() {
	c := cache.NewCache()
	d := db.DataBase{}
	d.Connect("postgres", db.DbInfo)
	c.RestoreDataFromDB(d)
	sub := subscriber.NewSubscriber()

	sub.Sc.Subscribe("order", func(m *stan.Msg) {
		fmt.Println("Sub got new message")
		structData, err := model.ReadJSON(m.Data)
		if err != nil {
			log.Fatalf("error: %s", err.Error())
		}
		c.SaveData(structData)
		d.InsertData(structData)
	},
		stan.DeliverAllAvailable())

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl.Execute(w, model.Order{})
	})
	http.HandleFunc("/order", func(w http.ResponseWriter, req *http.Request) {
		uid := req.URL.Query().Get("uid")
		order, err := c.GetOrderById(uid)
		fmt.Println(order)
		if err != nil {
			log.Println("no order with that uid")
		}
		tmpl.Execute(w, order)

	})

	http.ListenAndServe(":8080", nil)
}
