package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"

	"example.com/ex_bun/configs"
)

func main() {
	ctx := context.Background()
	sqldb, err := sql.Open(sqliteshim.ShimName, configs.Cfg.DB.DSN)
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	// add BUNDEBUG
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	var num int
	err = db.NewSelect().ColumnExpr("1").Scan(ctx, &num)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get num:%d\n", num)
}
