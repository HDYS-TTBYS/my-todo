package main

import (
	"HDYS-TTBYS/my-todo/ent"
	"HDYS-TTBYS/my-todo/ent/migrate"
	"context"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=postgres port=5432 user=api_user dbname=api_db password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
