package database

import (
	"context"
	"log"
	"os"

	"github.com/cronui/cronui/ent"
	"github.com/cronui/cronui/internal/config"
)

func NewDatabase(conf *config.Database) (*ent.Client, error) {
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		log.Println("cannot create data directory")
		return nil, err
	}

	client, err := ent.Open(conf.Driver, conf.GetDatabaseDsn())
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
