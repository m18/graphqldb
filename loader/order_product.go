package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/m18/graphqldb/db"
	"github.com/m18/graphqldb/model"
)

// LoadOrderProducts load order products via dataloader
func LoadOrderProducts(ctx context.Context, orderID int32) ([]*model.OrderProduct, error) {
	v, err := loadOne(ctx, orderProductLoaderKey, orderID)
	if err != nil || v == nil {
		return nil, err
	}
	res, ok := v.([]*model.OrderProduct)
	if !ok {
		return nil, fmt.Errorf("wrong type: %T", v)
	}
	return res, nil
}

type orderProductLoader struct {
	c *db.Client
}

func newOrderProductsBatchLoadFunc(c *db.Client) dataloader.BatchFunc {
	return (&orderProductLoader{c}).loadBatch
}

func (l *orderProductLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)
	ids, err := ints(keys)
	if err != nil {
		return loadBatchError(err, n)
	}

	ops, err := l.c.GetOrderProducts(ctx, ids)
	if err != nil {
		return loadBatchError(err, n)
	}

	res := make([]*dataloader.Result, n)
	for k, v := range ops {
		// results must be in the same order as keys
		i := mustIndex(ids, k)
		res[i] = &dataloader.Result{Data: v}
	}

	return res
}
