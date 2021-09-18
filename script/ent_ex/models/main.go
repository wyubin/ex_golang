package models

import (
	"context"
	"database/sql"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"example.com/ent_ex/configs"
	"example.com/ent_ex/ent"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB = ConnectDB()
)

func ConnectDB() *ent.Client {
	db, err := sql.Open(configs.Cfg.DB.Type, configs.Cfg.DB.DSN)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", configs.Cfg.DB.Type, err)
		return nil
	}
	drv := entsql.OpenDB(dialect.SQLite, db)
	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil
	}
	return client
}
