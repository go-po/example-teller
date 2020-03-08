package main

import (
	"context"
	"flag"
	"github.com/go-po/example-teller/internal/api"
	"github.com/go-po/example-teller/internal/app"
	"github.com/go-po/example-teller/internal/handlers"
	"github.com/go-po/po"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	flagBroker string
	flagStore  string

	flagHelp bool

	defaultRabbitUrl = "amqp://teller:teller@localhost:5670"
	defaultPGUrl     = "postgres://teller:teller@localhost:5430/teller?sslmode=disable"
)

func main() {

	flag.StringVar(&flagStore, "store", "inmemory", `Set the store
	inmemory: all data in memory
	pg: uses a connection url matching the docker-compose.yaml file
	provide your own connection url
`)
	flag.StringVar(&flagBroker, "broker", "channels", `Set the broker
	channels: use the channels broker
	rabbit: uses the RabbitMQ matching the docker-compose.yaml file
	provide your own RabbitMQ url
`)

	flag.BoolVar(&flagHelp, "help", false, "Print this help message")

	flag.Parse()

	if flagHelp {
		flag.Usage()
		os.Exit(0)
	}

	var broker po.Option
	if flagBroker == "rabbit" {
		broker = po.WithBrokerRabbit(defaultRabbitUrl, "teller", "app")
	} else if flagBroker != "channels" && strings.TrimSpace(flagBroker) != "" {
		broker = po.WithBrokerRabbit(flagBroker, "teller", "app")
	} else {
		broker = po.WithBrokerChannel()
	}

	var store po.Option
	if flagStore == "pg" {
		store = po.WithStorePGUrl(defaultPGUrl)
	} else if flagStore != "inmemory" && strings.TrimSpace(flagStore) != "" {
		store = po.WithStorePGUrl(flagStore)
	} else {
		store = po.WithStoreInMemory()
	}

	log.Printf("Starting Teller with store [%s] and broker [%s]", flagStore, flagBroker)

	rootCtx := context.Background()

	// initialize PO
	dao, err := po.NewFromOptions(broker, store)
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
