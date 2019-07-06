package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/m18/graphqldb/model"
)

// OrderResolver resolves order properties
type OrderResolver struct {
	order *model.Order
}

// ID resolves order ID
func (r *OrderResolver) ID() int32 {
	return r.order.ID
}

// Time resolves order time
func (r *OrderResolver) Time() graphql.Time {
	return graphql.Time{Time: r.order.Time}
}

// Customer returns customer resolver
func (r *OrderResolver) Customer(ctx context.Context) (*CustomerResolver, error) {
	return newCustomer(ctx, r.order.CustomerID)
}

// Products returns product resolvers
func (r *OrderResolver) Products(ctx context.Context) ([]*OrderProductResolver, error) {
	return newOrderProducts(ctx, r.order.ID)
}
