package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-cc/ent_orm_notes/schema_edges/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(192.168.1.100:3306)/schema_edges?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}
}
