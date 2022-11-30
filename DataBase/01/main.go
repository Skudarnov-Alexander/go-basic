package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	db, err := sql.Open("sqlite3",
		"db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// работаем с базой
	// ...
	// можем продиагностировать соединение
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}
	// в процессе работы
}
