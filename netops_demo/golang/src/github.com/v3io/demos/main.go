package main

import (
	"os"

	"github.com/v3io/demos/functions/ingester"
	"github.com/v3io/demos/functions/querier"

	"github.com/nuclio/nuclio-sdk-go"
)

func Ingest(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	return ingest.Ingest(context, event)
}

func Query(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	return query.Query(context, event)
}

// InitContext runs only once when the function runtime starts
func InitContext(context *nuclio.Context) error {
	switch os.Getenv("NUCLIO_HANDLER") {
	case "main:Ingest":
		return ingest.InitContext(context)
	case "main:Query":
		return query.InitContext(context)
	}

	return nil
}
