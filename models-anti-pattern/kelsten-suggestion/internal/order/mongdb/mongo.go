package mongdb

import (
	"blog/models-anti-pattern/kelsten-suggestion/internal/order"
	"fmt"
	"time"
)

// db provides underlying MongoDB implementation of order functionality.
type db struct {
	conn string
}

// New returns configured MongoDB client that after successful DB connection
// can manipulate order collection content.
func New(conn string) *db {
	return &db{conn: conn}
}

func (d *db) Pinger() error {
	return nil
}

func (d *db) CreateOrder(date time.Time, item, category string) error {
	order := &order.Order{
		Date:     date,
		Item:     item,
		Category: category,
	}
	fmt.Println(order)
	//... rest of the insertion logic that uses order var as document to insert
	return nil
}

func (d *db) RemoveOrder(ID int) error {
	return nil
}

func (d *db) UpdateOrder(new, old order.Order) error {
	return nil
}

func (d *db) FindOrder(ID int) (order.Order, error) {
	var o order.Order
	return o, nil
}
