package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-cc/ent_orm_notes/user_index2/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(192.168.1.104:3306)/use_index2?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}
	err = Do(ctx, client)
	fmt.Println(err)
}

func Do(ctx context.Context, client *ent.Client) error {
	// Unlike `Save`, `SaveX` panics if an error occurs.
	tlv := client.City.
		Create().
		SetName("TLV").
		SaveX(ctx)
	nyc := client.City.
		Create().
		SetName("NYC").
		SaveX(ctx)
	// Add a street "ST" to "TLV".
	client.Street.
		Create().
		SetName("ST").
		SetCity(tlv).
		SaveX(ctx)
	// This operation will fail because "ST"
	// is already created under "TLV".
	_, err := client.Street.
		Create().
		SetName("ST").
		SetCity(tlv).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("expecting creation to fail")
	}
	// Add a street "ST" to "NYC".
	client.Street.
		Create().
		SetName("ST").
		SetCity(nyc).
		SaveX(ctx)
	return nil
}
