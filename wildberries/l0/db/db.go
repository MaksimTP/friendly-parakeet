package db

import (
	"database/sql"
	"fmt"
	"log"
	model "main/model"

	"github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "wb_lvl0"
)

type DataBase struct {
	db *sql.DB
}

var DbInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func (d *DataBase) Connect(driverName string, dbInfo string) {
	db, err := sql.Open(driverName, dbInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	d.db = db
}

func (d *DataBase) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

const (
	insertStatementPayment = `
INSERT INTO payment (id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	insertStatementDelivery = `
INSERT INTO delivery (id, name, phone, zip, city, address, region, email)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	insertStatementItem = `
INSERT INTO item (id, chrt_id, track_number, price, rid, sale, size, total_price, nm_id, brand, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	insertStatementOrder = `
INSERT INTO "order" (order_uid, track_number, entry, delivery_id, payment_id, items_ids, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
)

func (d *DataBase) GetNextIdToInsert(tableName string) int {
	query := "SELECT COUNT(*) FROM " + tableName
	rows, _ := d.db.Query(query)
	var id int
	rows.Next()
	err := rows.Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	id++
	return id
}

func (d *DataBase) InsertData(data model.Order) {

	deliveryID := d.GetNextIdToInsert("delivery")
	_, err := d.db.Exec(insertStatementDelivery, deliveryID, data.Delivery.Name, data.Delivery.Phone, data.Delivery.Zip, data.Delivery.City, data.Delivery.Address, data.Delivery.Region, data.Delivery.Email)
	if err != nil {
		log.Println(err.Error())
	}

	paymentId := d.GetNextIdToInsert("payment")
	_, err = d.db.Exec(insertStatementPayment, paymentId, data.Payment.Transaction, data.Payment.RequestID, data.Payment.Currency, data.Payment.Provider, data.Payment.Amount, data.Payment.PaymentDt, data.Payment.Bank, data.Payment.DeliveryCost, data.Payment.GoodsTotal, data.Payment.CustomFee)
	if err != nil {
		log.Println(err.Error())
	}
	itemId := d.GetNextIdToInsert("item")

	itemIds := make([]int, 0)
	itemIds = append(itemIds, itemId)
	for _, v := range data.Items {
		_, err = d.db.Exec(insertStatementItem, itemId, v.ChrtID, v.TrackNumber, v.Price, v.Rid, v.Sale, v.Size, v.TotalPrice, v.NmID, v.Brand, v.Status)
		if err != nil {
			log.Println(err.Error())
		}
		itemId++
		itemIds = append(itemIds, itemId)
	}

	_, err = d.db.Exec(insertStatementOrder, data.OrderUid, data.TrackNumber, data.Entry, deliveryID, paymentId, pq.Array(itemIds), data.Locale, data.InternalSignature, data.CustomerID, data.DeliveryService, data.Shardkey, data.SmID, data.DateCreated, data.OofShard)

	if err != nil {
		log.Println(err.Error())
	}
}

// func (d *DataBase) GetInfoById(id int) {
// 	query := "SELECT * FROM "
// }
