package db

import (
	"context"
	"time"

	"github.com/m18/graphqldb/model"
)

// GetOrders returns a list of orders
func (c *Client) GetOrders(ctx context.Context, first int32) ([]*model.Order, error) {
	sql := `
		SELECT id, customer_id, time
		FROM orders 
		ORDER BY time
		LIMIT $1`

	rows, err := c.db.QueryContext(ctx, sql, first)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		res        []*model.Order
		id         int32
		customerID int32
		time       time.Time
	)
	for rows.Next() {
		err = rows.Scan(&id, &customerID, &time)
		if err != nil {
			return nil, err
		}
		res = append(
			res,
			&model.Order{
				ID:         id,
				CustomerID: customerID,
				Time:       time,
			})
	}
	return res, nil
}
