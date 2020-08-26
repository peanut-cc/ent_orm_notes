package main

import (
	"context"
	"log"

	"github.com/peanut-cc/ent_orm_notes/schema_notes/ent/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-cc/ent_orm_notes/schema_notes/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(10.211.55.3:3306)/schema_notes?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}
	client.User.Create().SetRequiredName("peanut").Save(ctx)
	client.User.Create().SetRequiredName("syncd").
		SetOptionalName("option_name").
		SetNilableName("nil_name").
		SetNilableName2("nil_name2").
		SetAge(18).
		SetAge2(20).
		SaveX(ctx)
	u := client.User.Query().Where(user.RequiredNameEQ("peanut")).OnlyX(ctx)
	log.Printf("required_name is:%v option_name is:%v nil_name is:%v nil_name2 is:%v age is :%v age2 is:%v\n", u.RequiredName,
		u.OptionalName, u.NilableName, u.NilableName2, u.Age, u.Age2)
	u2 := client.User.Query().Where(user.RequiredNameEQ("syncd")).OnlyX(ctx)
	log.Printf("required_name is:%v option_name is:%v nil_name is:%v nil_name2 is:%v age is :%v age2 is:%v\n", u2.RequiredName,
		u2.OptionalName, u2.NilableName, u2.NilableName2, u2.Age, u2.Age2)
}
