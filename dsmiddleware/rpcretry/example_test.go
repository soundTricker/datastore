package rpcretry_test

import (
	"context"
	"time"

	"go.mercari.io/datastore/clouddatastore"
	"go.mercari.io/datastore/dsmiddleware/rpcretry"
)

func Example_howToUse() {
	ctx := context.Background()
	client, err := clouddatastore.FromContext(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	mw := rpcretry.New(
		rpcretry.WithRetryLimit(5),
		rpcretry.WithMinBackoffDuration(10*time.Millisecond),
		rpcretry.WithMaxBackoffDuration(150*time.Microsecond),
		// rpcretry.WithMaxDoublings(2),
	)
	client.AppendMiddleware(mw)
}
