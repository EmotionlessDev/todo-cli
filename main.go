/*
Copyright © 2024 Svirin Artyom <emotionlesscode@gmail.com>
*/
package main

import (
	"context"
	"database/sql"
	"log"
	"time"
	"todo/cmd"

	_ "github.com/lib/pq"
)

func openDB() (*sql.DB, error) {
	dsn := "postgres://todo:todo@localhost:5432/todo?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("can't open db: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("can't ping db: %v", err)
	}

	return db, nil

}

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatalf("can't open db: %v", err)
	}
	defer db.Close()

	cmd.Execute()
}
