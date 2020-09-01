package main

import (
	"context"
	"fmt"
	"log"

	"github.com/peanut-cc/ent_orm_notes/graph_traversal/ent/pet"
	"github.com/peanut-cc/ent_orm_notes/graph_traversal/ent/user"

	"github.com/peanut-cc/ent_orm_notes/graph_traversal/ent/group"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-cc/ent_orm_notes/graph_traversal/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(10.211.55.3:3306)/graph_traversal?parseTime=True",
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
	//Gen(ctx, client)
	//Gen2(ctx, client)
	//Traverse(ctx, client)
	//Traverse2(ctx, client)
	//edgerLoading(ctx, client)
	edgerLoading2(ctx, client)
}

func Gen(ctx context.Context, client *ent.Client) error {
	hub, err := client.Group.Create().SetName("Github").Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating the group: %v", err)
	}
	// Create the admin of the group.
	// Unlike `Save`, `SaveX` panics if an error occurs.
	dan := client.User.Create().SetAge(29).SetName("Dan").AddManage(hub).SaveX(ctx)
	// Create "Ariel" and its pets.
	// Create "Ariel" and its pets.
	a8m := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		AddGroups(hub).
		AddFriends(dan).
		SaveX(ctx)
	pedro := client.Pet.
		Create().
		SetName("Pedro").
		SetOwner(a8m).
		SaveX(ctx)
	xabi := client.Pet.
		Create().
		SetName("Xabi").
		SetOwner(a8m).
		SaveX(ctx)

	// Create "Alex" and its pets.
	alex := client.User.
		Create().
		SetAge(37).
		SetName("Alex").
		SaveX(ctx)
	coco := client.Pet.
		Create().
		SetName("Coco").
		SetOwner(alex).
		AddFriends(pedro).
		SaveX(ctx)

	fmt.Println("Pets created:", pedro, xabi, coco)
	// Output:
	// Pets created: Pet(id=1, name=Pedro) Pet(id=2, name=Xabi) Pet(id=3, name=Coco)
	return nil
}

func Gen2(ctx context.Context, client *ent.Client) {
	lab, err := client.Group.Create().SetName("Gitlab").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating the group: %v", err)
	}
	client.User.Create().SetAge(18).SetName("peanut").AddGroups(lab).SaveX(ctx)
}

func Traverse(ctx context.Context, client *ent.Client) error {
	owner, err := client.Group.
		Query().
		Where(group.Name("Github")).
		QueryAdmin().
		QueryFriends().
		QueryPets().
		QueryOwner().
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying the owner: %v", err)
	}
	fmt.Println(owner)
	return nil
}

func Traverse2(ctx context.Context, client *ent.Client) error {
	pets, err := client.Pet.
		Query().
		Where(
			pet.HasOwnerWith(
				user.HasFriendsWith(
					user.HasManage(),
				),
			),
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying the pets: %v", err)
	}
	fmt.Println(pets)
	// Output:
	// [Pet(id=1, name=Pedro) Pet(id=2, name=Xabi)]
	return nil
}

func edgerLoading(ctx context.Context, client *ent.Client) {
	users, err := client.User.Query().WithPets().All(ctx)
	if err != nil {
		log.Fatalf("user query failed:%v", err)
	}
	log.Println(users)
	for _, u := range users {
		for _, p := range u.Edges.Pets {
			log.Printf("user (%v) -- > Pet (%v)\n", u.Name, p.Name)
		}
	}

}

func edgerLoading2(ctx context.Context, client *ent.Client) {
	users, err := client.User.
		Query().
		Where(
			user.AgeGT(18),
		).
		WithPets().
		WithGroups(func(q *ent.GroupQuery) {
			q.Limit(5)
			q.WithUsers().Limit(5)
		}).All(ctx)
	if err != nil {
		log.Fatalf("user query failed:%v", err)
	}
	log.Println(users)
	for _, u := range users {
		for _, p := range u.Edges.Pets {
			log.Printf("user (%v) --> Pet (%v)\n", u.Name, p.Name)
		}
		for _, g := range u.Edges.Groups {
			log.Printf("user (%v) -- Group (%v)\n", u.Name, g.Name)
		}
	}

}
