package order

import (
	"database/sql"

	"github.com/alissoncorsair/goapi/types"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	var id int
	err := s.db.QueryRow("INSERT into orders (user_id, total, status, address) VALUES ($1, $2, $3, $4) RETURNING id", order.UserID, order.Total, order.Status, order.Address).Scan(&id)

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT into order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)

	return err
}
