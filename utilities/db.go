package utilities

import (
	"fmt"
	"github.com/go-pg/pg"
)

type IDB interface {
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

func NewOrmDB(configUtil IConfigUtil) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     configUtil.GetConfig("dbHost") + ":" + configUtil.GetConfig("dbPort"),
		User:     configUtil.GetConfig("dbUser"),
		Password: configUtil.GetConfig("dbPassword"),
		Database: configUtil.GetConfig("dbName"),
	})

	db.AddQueryHook(dbLogger{})

	_, err := db.Exec("SELECT 1;")
	if err != nil {
		panic(err.Error())
	}
	return db
}
