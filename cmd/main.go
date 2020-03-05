package main

import (
	"context"
	"github.com/go-po/example-teller/internal/api"
	"github.com/go-po/example-teller/internal/app"
	"github.com/go-po/example-teller/internal/handlers"
	"github.com/go-po/po"
	"log"

	"net/http"
)

func main() {
	rootCtx := context.Background()

	// initialize PO
	dao, err := po.NewFromOptions(
		po.WithChannelBroker(),
		po.WithInMemoryStore(),
	)
	if err != nil {
		log.Fatalf("failed initialize po: %s", err)
	}

	// setup subscriber to act on all commands issued
	subscriptionId := "command handler"
	streamId := "vars:commands"
	handler := handlers.NewCommandSubscriber(dao)
	err = dao.Subscribe(rootCtx, subscriptionId, streamId, handler)
	if err != nil {
		log.Fatalf("failed subscribing: %s", err)
	}

	app := app.New(dao)

	err = http.ListenAndServe(":8000", api.Root(app))
	log.Printf("server stopped: %s", err)
}
