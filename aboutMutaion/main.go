package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-cc/ent_orm_notes/aboutMutaion/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(10.211.55.3:3306)/aboutMutaion?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}
	//Do(ctx, client)
	Do2(ctx, client)
}

func Do(ctx context.Context, client *ent.Client) {
	creator := client.User.Create()
	SetAgeName(creator.Mutation())
	creator.SaveX(ctx)
	updater := client.User.UpdateOneID(1)
	SetAgeName(updater.Mutation())
	updater.SaveX(ctx)
}

// SetAgeName sets the age and the name for any mutation.
func SetAgeName(m *ent.UserMutation) {
	m.SetAge(32)
	m.SetName("Ariel")
}

func Do2(ctx context.Context, client *ent.Client) {
	creator1 := client.User.Create().SetAge(18)
	SetName(creator1.Mutation(), "a8m")
	creator1.SaveX(ctx)
	creator2 := client.Pet.Create().SetAge(16)
	SetName(creator2.Mutation(), "pedro")
	creator2.SaveX(ctx)
}

// SetNamer wraps the 2 methods for getting
// and setting the "name" field in mutations.
type SetNamer interface {
	SetName(string)
	Name() (string, bool)
}

func SetName(m SetNamer, name string) {
	if _, exist := m.Name(); !exist {
		m.SetName(name)
	}
}
