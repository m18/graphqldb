package resolver

import (
	"context"

	"github.com/m18/graphqldb/db"
)

// Query is the root resolver
type Query struct {
	c *db.Client
}

// Orders returns order resolvers
func (r *Query) Orders(ctx context.Context, args struct {
	First int32
}) ([]*OrderResolver, error) {
	orders, err := r.c.GetOrders(ctx, args.First)
	if err != nil {
		return nil, err
	}
	res := make([]*OrderResolver, 0, len(orders))
	for _, o := range orders {
		res = append(res, &OrderResolver{o})
	}
	return res, nil
}

// NewQuery creates a new *Query and returns it
func NewQuery(c *db.Client) *Query {
	return &Query{c}
}
