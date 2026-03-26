package repository

import (
	"database/sql"
	"fmt"

	orderModel "github.com/daniloAleite/go-orders-api/internal/model/order"
)

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{
		db: db,
	}
}

func (r *PostgresOrderRepository) Create(customer string, amount float64) (orderModel.Order, error) {
	query := `
		INSERT INTO orders (customer, amount, status)
		VALUES ($1, $2, $3)
		RETURNING id, customer, amount, status, created_at
	`

	var order orderModel.Order

	err := r.db.QueryRow(query, customer, amount, "created").Scan(
		&order.ID,
		&order.Customer,
		&order.Amount,
		&order.Status,
		&order.CreatedAt,
	)
	if err != nil {
		return orderModel.Order{}, fmt.Errorf("insert order: %w", err)
	}

	return order, nil
}

func (r *PostgresOrderRepository) List() ([]orderModel.Order, error) {
	query := `
		SELECT id, customer, amount, status, created_at
		FROM orders
		ORDER BY id ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("select orders: %w", err)
	}
	defer rows.Close()

	orders := make([]orderModel.Order, 0)

	for rows.Next() {
		var order orderModel.Order

		if err := rows.Scan(
			&order.ID,
			&order.Customer,
			&order.Amount,
			&order.Status,
			&order.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan order row: %w", err)
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate order rows: %w", err)
	}

	return orders, nil
}
