package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-cc/ent_orm_notes/index_use/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(192.168.1.104:3306)/use_index?parseTime=True")
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
