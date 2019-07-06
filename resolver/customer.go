package resolver

import (
	"context"

	"github.com/m18/graphqldb/loader"
	"github.com/m18/graphqldb/model"
)

// CustomerResolver resolves customer properties
type CustomerResolver struct {
	cust *model.Customer
}

// ID resolves customer ID
func (r *CustomerResolver) ID() int32 {
	return r.cust.ID
}

// Name resolves customer name
func (r *CustomerResolver) Name() string {
	return r.cust.Name
}

func newCustomer(ctx context.Context, customerID int32) (*CustomerResolver, error) {
	c, err := loader.LoadCustomer(ctx, customerID)
	if err != nil {
		return nil, err
	}
	return &CustomerResolver{c}, nil
}
