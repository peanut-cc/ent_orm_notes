package main

import (
	"context"
	"log"

	"github.com/peanut-cc/ent_orm_notes/hook_runtime_example/ent/hook"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-cc/ent_orm_notes/hook_runtime_example/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(10.211.55.3:3306)/hook_runtime?parseTime=True")
	if err != nil {
		log.Fatalf("ent open db error:%v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}

	//client.Use(func(next ent.Mutator) ent.Mutator {
	//	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	//		start := time.Now()
	//		defer func() {
	//			log.Printf("Op=%s\tType=%s\tTime=%s\tConreteType=%T", m.Op(), m.Type(), time.Since(start), m)
	//		}()
	//		return next.Mutate(ctx, m)
	//	})
	//})
	//Do(ctx, client)

	client.User.Use(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			return next.Mutate(ctx, m)
		})
	})

	client.Use(hook.On(Logger()))

}

func Do(ctx context.Context, client *ent.Client) {
	client.User.Create().SetName("peanut").SaveX(ctx)
}
