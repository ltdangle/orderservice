package read

import (
	"database/sql"
	"log"
)

type Order struct {
	Uuid        sql.NullString
	CustomerId  sql.NullString
	PaymentId   sql.NullString
	PaymentDate sql.NullTime
	Items       []OrderItem
	Total       int
	Status      sql.NullString
	CreatedAt   sql.NullTime
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
SELECT orders.uuid, orders.created_at, orders.status, p.payment_id, p.date
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
	err = row.Scan(&order.Uuid, &order.CreatedAt, &order.Status, &order.PaymentId, &order.PaymentDate)
	if err != nil {
		return nil, err
	}

	order.Items, err = f.orderItemFinder.Find(uuid)
	if err != nil {
		return nil, err
	}

	// Calculate total.
	for _, item := range order.Items {
		order.Total += item.Price
	}

	return &order, nil
}

type OrderFinderActiveById interface {
	FindActive(uuid string) (*Order, error)
}

type FinderActiveById struct {
	db *sql.DB
}

func NewOrderFinderActiveById(db *sql.DB) *FinderActiveById {
	return &FinderActiveById{
		db: db,
	}
}

func (f FinderActiveById) FindActive(uuid string) (*Order, error) {
	stmt, err := f.db.Prepare(`
SELECT *
FROM orders
WHERE status = 'created'
AND uuid = ?;
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var order Order
	for rows.Next() {
		err = rows.Scan(&order.Uuid, &order.Status, &order.CustomerId, &order.PaymentId, &order.CreatedAt)

		if err != nil {
			return nil, err
		}
	}

	return &order, nil
}
