package read

import (
	"database/sql"
)

type OrderItem struct {
	Uuid        string
	OrderId     string
	Title       string
	Description string
	ProductId   string
	Price       int
}

// IOrderItemFinderByOrderId finds all order items for order.
type IOrderItemFinderByOrderId interface {
	Find(orderUuid string) ([]OrderItem, error)
}

// OrderItemFinderById implementation.
type OrderItemFinderById struct {
	db *sql.DB
}

// NewOrderItemFinderById Constructor.
func NewOrderItemFinderById(db *sql.DB) *OrderItemFinderById {
	return &OrderItemFinderById{
		db: db,
	}
}

func (f OrderItemFinderById) Find(orderUuid string) ([]OrderItem, error) {
	// Prepare a SQL statement
	stmt, err := f.db.Prepare(`
SELECT order_items.uuid, order_items.order_id, order_items.title, order_items.description, order_items.product_id, order_items.price
FROM order_items
WHERE order_items.order_id = ?;
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(orderUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderItems []OrderItem
	for rows.Next() {
		var orderItem OrderItem
		err = rows.Scan(&orderItem.Uuid, &orderItem.OrderId, &orderItem.Title, &orderItem.Description, &orderItem.ProductId, &orderItem.Price)
		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}
