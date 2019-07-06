package db

import (
	"context"

	"github.com/lib/pq"
	"github.com/m18/graphqldb/model"
)

// GetOrderProducts returns products by order
func (c *Client) GetOrderProducts(ctx context.Context, orderIDs []int32) (map[int32][]*model.OrderProduct, error) {
	sql := `
		SELECT op.order_id, p.id, p.name, p.price, op.quantity
		FROM order_products op INNER JOIN products p ON op.product_id = p.id
		WHERE op.order_id = ANY($1)
		ORDER by op.order_id`

	rows, err := c.db.QueryContext(ctx, sql, pq.Array(orderIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		res       = make(map[int32][]*model.OrderProduct, len(orderIDs))
		tmp       []*model.OrderProduct
		groupID   int32
		orderID   int32
		productID int32
		name      string
		price     int32
		quantity  int32
	)
	for rows.Next() {
		err = rows.Scan(&orderID, &productID, &name, &price, &quantity)
		if err != nil {
			return nil, err
		}
		if groupID == 0 {
			groupID = orderID
		}
		op := &model.OrderProduct{
			ID:       productID,
			Name:     name,
			Price:    price,
			Quantity: quantity,
		}
		if orderID != groupID {
			res[groupID] = tmp
			tmp = nil
			groupID = orderID
		}
		tmp = append(tmp, op)
	}
	res[groupID] = tmp
	return res, nil
}
