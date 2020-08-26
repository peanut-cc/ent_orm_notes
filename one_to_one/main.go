package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-cc/ent_orm_notes/one_to_one/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(192.168.1.100:3306)/one_to_one?parseTime=True")
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
	peanut, err := client.User.
		Create().
		SetAge(18).
		SetName("peanut").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %v", err)
	}
	fmt.Println("user:", peanut)
	expired, err := time.Parse(time.RFC3339, "2019-12-08T15:04:05Z")
	if err != nil {
		return err
	}
	card1, err := client.Card.
		Create().
		SetOwner(peanut).
		SetExpired(expired).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %v", err)
	}
	fmt.Println("card:", card1)
	// Only returns the card of the user,
	// and expects that there's only one.
	// 这里是正向查找，查找用户所持有的卡
	card2, err := peanut.QueryCard().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying card: %v", err)
	}
	fmt.Println("card:", card2)
	// The Card entity is able to query its owner using
	// its back-reference.
	// 这里根据card反向查找所属的用户
	owner, err := card2.QueryOwner().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying owner: %v", err)
	}
	fmt.Println("owner:", owner)
	return nil
}
