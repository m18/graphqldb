package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/m18/graphqldb/db"
	"github.com/m18/graphqldb/loader"
	"github.com/m18/graphqldb/resolver"

	_ "github.com/lib/pq"
)

func main() {
	db, err := db.New("localhost", 5432, "store", "p", "store")
	if err != nil {
		log.Fatal(err)
	}

	s := graphql.MustParseSchema(schema, resolver.NewQuery(db))
	loaders := loader.Init(db)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: newHandler(s, loaders),
	}

	handleShutdown(db, srv)

	if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleShutdown(c *db.Client, srv *http.Server) {
	ch := make(chan os.Signal, 1)
	go func() {
		<-ch

		err := c.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = srv.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()
	signal.Notify(ch, os.Interrupt, os.Kill)
}
