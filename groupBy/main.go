package main

import (
	"context"
	"log"

	"github.com/peanut-cc/ent_orm_notes/groupBy/ent/user"

	_ "github.com/go-sql-driver/mysql"

	"github.com/peanut-cc/ent_orm_notes/groupBy/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(10.211.55.3:3306)/groupBy?parseTime=True",
		ent.Debug())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}
	//GenData(ctx, client)
	//Do(ctx, client)
	Do2(ctx, client)

}

func GenData(ctx context.Context, client *ent.Client) {
	client.User.Create().SetName("peanut").SetAge(18).SaveX(ctx)
	client.User.Create().SetName("jack").SetAge(20).SaveX(ctx)
	client.User.Create().SetName("steve").SetAge(22).SaveX(ctx)
	client.User.Create().SetName("peanut-cc").SetAge(18).SaveX(ctx)
	client.User.Create().SetName("jack-dd").SetAge(18).SaveX(ctx)
}

func Do(ctx context.Context, client *ent.Client) {
	var v []struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Sum   int    `json:"sum"`
		Count int    `json:"count"`
	}
	client.User.
		Query().
		GroupBy(
			user.FieldName, user.FieldAge,
		).
		Aggregate(
			ent.Count(),
			ent.Sum(user.FieldAge),
		).
		ScanX(ctx, &v)
	log.Println(v)

}

func Do2(ctx context.Context, client *ent.Client) {
	names := client.User.Query().GroupBy(user.FieldName).StringsX(ctx)
	log.Println(names)
}
