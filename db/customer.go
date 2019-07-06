package db

import (
	"context"

	"github.com/lib/pq"
	"github.com/m18/graphqldb/model"
)

// GetCustomers returns a list of customers
func (c *Client) GetCustomers(ctx context.Context, customerIDs []int32) ([]*model.Customer, error) {
	sql := `
		SELECT id, name
		FROM customers
		WHERE id = ANY($1)`

	rows, err := c.db.QueryContext(ctx, sql, pq.Array(customerIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		res  = make([]*model.Customer, 0, len(customerIDs))
		id   int32
		name string
	)
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		res = append(res, &model.Customer{
			ID:   id,
			Name: name,
		})
	}
	return res, nil
}
