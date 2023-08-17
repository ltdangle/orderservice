package read

import (
	"database/sql"
	"log"
	"time"
)

type Order struct {
	Uuid        string
	CustomerId  string
	PaymentId   string
	PaymentDate time.Time
	Items       []OrderItem
	CreatedAt   time.Time
}

type OrderFinderById interface {
	Find(uuid string) (*Order, error)
}

// FinderById implementation.
type FinderById struct {
	db              *sql.DB
	orderItemFinder IOrderItemFinderByOrderId
}

// NewOrderFinderById NewSqlOrderRepo Constructor.
func NewOrderFinderById(db *sql.DB, orderItemFinder IOrderItemFinderByOrderId) *FinderById {
	return &FinderById{
		db:              db,
		orderItemFinder: orderItemFinder,
	}
}

func (f FinderById) Find(uuid string) (*Order, error) {
	// Prepare a SQL statement
	stmt, err := f.db.Prepare(`
SELECT orders.uuid, orders.created_at, p.payment_id, p.date
FROM orders 
LEFT JOIN payments p 
ON orders.uuid = p.order_id
WHERE orders.uuid = ?;
`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the prepared statement with a specific value
	row := stmt.QueryRow(uuid)

	var order Order
	err = row.Scan(&order.Uuid, &order.CreatedAt, &order.PaymentId, &order.PaymentDate)
	if err != nil {
		return nil, err
	}

	order.Items, err = f.orderItemFinder.Find(uuid)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
