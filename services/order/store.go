package order

import (
	"database/sql"
	"fmt"

	"github.com/matimortari/go-backend/types"
)

type Store struct {
	db *sql.DB
}

// Create a new Store struct
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Create a new order in the database
func (s *Store) CreateOrder(order types.Order) (int, error) {
	query := `
		INSERT INTO orders (userId, total, status, address)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	var id int
	err := s.db.QueryRow(query, order.UserID, order.Total, order.Status, order.Address).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not create order: %w", err)
	}

	return id, nil
}

// Create a new order item in the database
func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	query := `
		INSERT INTO order_items (orderId, productId, quantity, price)
		VALUES ($1, $2, $3, $4)`

	_, err := s.db.Exec(query, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	if err != nil {
		return fmt.Errorf("could not create order item: %w", err)
	}

	return nil
}
