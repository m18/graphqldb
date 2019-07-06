package resolver

import (
	"context"

	"github.com/m18/graphqldb/loader"
	"github.com/m18/graphqldb/model"
)

// OrderProductResolver resolves order product properties
type OrderProductResolver struct {
	prod *model.OrderProduct
}

// ID resolves product ID
func (r *OrderProductResolver) ID() int32 {
	return r.prod.ID
}

// Name resolves product name
func (r *OrderProductResolver) Name() string {
	return r.prod.Name
}

// Price resolves product price
func (r *OrderProductResolver) Price() float64 {
	return float64(r.prod.Price) / 100
}

// Quantity resolves product quantity in the order
func (r *OrderProductResolver) Quantity() int32 {
	return r.prod.Quantity
}

// TotalPrice resolves total product price for the quantity
func (r *OrderProductResolver) TotalPrice() float64 {
	return r.Price() * float64(r.Quantity())
}

func newOrderProducts(ctx context.Context, orderID int32) ([]*OrderProductResolver, error) {
	ops, err := loader.LoadOrderProducts(ctx, orderID)
	if err != nil {
		return nil, err
	}
	res := make([]*OrderProductResolver, 0, len(ops))
	for _, op := range ops {
		res = append(res, &OrderProductResolver{op})
	}
	return res, nil
}
