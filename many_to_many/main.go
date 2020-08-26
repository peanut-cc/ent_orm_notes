package main

import (
	"context"
	"fmt"
	"log"

	"github.com/peanut-cc/ent_orm_notes/many_to_many/ent/group"
	"github.com/peanut-cc/ent_orm_notes/many_to_many/ent/user"

	_ "github.com/go-sql-driver/mysql"

	"github.com/peanut-cc/ent_orm_notes/many_to_many/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(192.168.1.100:3306)/many_to_many?parseTime=True")
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
	hub := client.Group.Create().SetName("GitHub").SaveX(ctx)
	lab := client.Group.Create().SetName("GitLab").SaveX(ctx)
	peanut := client.User.Create().SetName("peanut").AddGroups(hub, lab).SaveX(ctx)
	jack := client.User.Create().SetName("jack").AddGroups(hub).SaveX(ctx)

	groups, err := peanut.QueryGroups().All(ctx)
	if err != nil {
		return fmt.Errorf("querying a8m groups: %v", err)
	}
	fmt.Println(groups)

	groups, err = jack.QueryGroups().All(ctx)
	if err != nil {
		return fmt.Errorf("querying nati groups: %v", err)
	}
	fmt.Println(groups)
	users, err := peanut.QueryGroups().
		Where(group.Not(group.HasUsersWith(user.Name("jack")))).
		QueryUsers().
		QueryGroups().
		QueryUsers().
		All(ctx)
	if err != nil {
		return fmt.Errorf("traversing the graph: %v", err)
	}
	fmt.Println(users)
	return nil
}
