package loader

import (
	"context"
	"fmt"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/m18/graphqldb/db"
)

type contextKey string

const (
	customerLoaderKey     contextKey = "customer"
	orderProductLoaderKey contextKey = "orderProduct"
)

// Init initializes and returns Map
func Init(c *db.Client) Map {
	return Map{
		customerLoaderKey:     newCustomerBatchLoadFunc(c),
		orderProductLoaderKey: newOrderProductsBatchLoadFunc(c),
	}
}

// Map maps loader keys to batch-load funcs
type Map map[contextKey]dataloader.BatchFunc

// Attach attaches dataloaders to the request's context
func (m Map) Attach(ctx context.Context) context.Context {
	for k, batchFunc := range m {
		ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(batchFunc))
	}
	return ctx
}

func extract(ctx context.Context, k contextKey) (*dataloader.Loader, error) {
	res, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("cannot find a loader: %s", k)
	}
	return res, nil
}

func loadOne(ctx context.Context, loaderKey contextKey, k int32) (interface{}, error) {
	ldr, err := extract(ctx, loaderKey)
	if err != nil {
		return nil, err
	}
	v, err := ldr.Load(ctx, key(k))()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func loadMany(ctx context.Context, loaderKey contextKey, kk []int32) ([]interface{}, error) {
	ldr, err := extract(ctx, loaderKey)
	if err != nil {
		return nil, err
	}
	vv, errs := ldr.LoadMany(ctx, keys(kk))()
	if errs != nil && len(errs) > 0 {
		// TODO: join all errors into one & return it
		return nil, errs[0]
	}
	return vv, nil
}

func mustIndex(nn []int32, n int32) int {
	for i, v := range nn {
		if n == v {
			return i
		}
	}
	panic(fmt.Sprintf("could not find %d in %v", n, nn))
}

func key(k int32) dataloader.Key {
	return dataloader.StringKey(strconv.Itoa(int(k)))
}

func keys(kk []int32) []dataloader.Key {
	ss := make([]string, 0, len(kk))
	for _, k := range kk {
		ss = append(ss, strconv.Itoa(int(k)))
	}
	return dataloader.NewKeysFromStrings(ss)
}

func ints(kk dataloader.Keys) ([]int32, error) {
	res := make([]int32, 0, len(kk))
	for _, s := range kk.Keys() {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}
		res = append(res, int32(v))
	}
	return res, nil
}

func loadBatchError(err error, n int) []*dataloader.Result {
	r := &dataloader.Result{Error: err}
	res := make([]*dataloader.Result, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, r)
	}
	return res
}
