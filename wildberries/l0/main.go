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
	order := model.Order{}

	c := cache.NewCache()
	d := db.DataBase{}
	d.Connect("postgres", db.DbInfo)
	// c.RestoreDataFromDB(d)

	sub := subscriber.NewSubscriber()

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

	// select {}

	tmpl, err := template.ParseFiles("html_template/index.html")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl.Execute(w, order)
	})
	http.ListenAndServe(":8000", nil)
}
