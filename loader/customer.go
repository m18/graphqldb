package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/m18/graphqldb/db"
	"github.com/m18/graphqldb/model"
)

// LoadCustomer loads customer via dataloader
func LoadCustomer(ctx context.Context, id int32) (*model.Customer, error) {
	v, err := loadOne(ctx, customerLoaderKey, id)
	if err != nil || v == nil {
		return nil, err
	}
	res, ok := v.(*model.Customer)
	if !ok {
		return nil, fmt.Errorf("wrong type: %T", v)
	}
	return res, nil
}

type customerLoader struct {
	c *db.Client
}

func newCustomerBatchLoadFunc(c *db.Client) dataloader.BatchFunc {
	return (&customerLoader{c}).loadBatch
}

func (l *customerLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)
	ids, err := ints(keys)
	if err != nil {
		return loadBatchError(err, n)
	}

	cc, err := l.c.GetCustomers(ctx, ids)
	if err != nil {
		return loadBatchError(err, n)
	}

	res := make([]*dataloader.Result, n)
	for _, c := range cc {
		// results must be in the same order as keys
		i := mustIndex(ids, c.ID)
		res[i] = &dataloader.Result{Data: c}
	}

	return res
}
