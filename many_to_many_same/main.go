package main

import (
	"context"
	"fmt"
	"log"

	"github.com/peanut-cc/ent_orm_notes/many_to_many_same/ent/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-cc/ent_orm_notes/many_to_many_same/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(192.168.1.100:3306)/many_to_many_same?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}
	Do(ctx, client)
}

func Do(ctx context.Context, client *ent.Client) error {
	a8m := client.User.Create().SetName("a8m").SaveX(ctx)

	nati := client.User.Create().SetName("nati").AddFollowers(a8m).SaveX(ctx)

	flw := a8m.QueryFollowing().AllX(ctx)
	fmt.Println(flw)

	flr := a8m.QueryFollowers().AllX(ctx)
	fmt.Println(flr)

	flw = nati.QueryFollowing().AllX(ctx)
	fmt.Println(flw)

	flr = nati.QueryFollowers().AllX(ctx)
	fmt.Println(flr)

	names := client.User.
		Query().
		Where(user.Not(user.HasFollowers())).
		GroupBy(user.FieldName).
		StringsX(ctx)
	fmt.Println(names)
	return nil

}
