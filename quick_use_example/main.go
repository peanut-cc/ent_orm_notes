package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peanut-pg/ent_orm_notes/quick_user_example/ent"
	"github.com/peanut-pg/ent_orm_notes/quick_user_example/ent/car"
	"github.com/peanut-pg/ent_orm_notes/quick_user_example/ent/user"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(192.168.1.104:3306)/ent_orm?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}
	//CreateUser(ctx, client)
	//peanut, err := QueryUser(ctx, client)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Fatalf("query user name is:%v, aget is %v", peanut.Name, peanut.Age)
	//CreateCars(ctx, client)
	//peanut_pg, err := QueryUserByName(ctx, client, "peanut_pg")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//QueryCars(ctx, peanut_pg)
	//car, err := QueryCarByModel(ctx, client)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//QueryCarUser(ctx, car)
	//CreateGraph(ctx, client)

}

// CreateUser 创建用户 name=peanut, age=18
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(18).
		SetName("peanut").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

// QueryUser 查询用户 where name=peanut
func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.NameEQ("peanut")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

// CreateCars 创建 Tesla 和Ford 汽车，加该汽车属于user: peanut_pg
func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// creating new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}

	// creating new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}
	log.Println("car was created: ", ford)

	// create a new user, and add it the 2 cars.
	peanut_pg, err := client.User.
		Create().
		SetAge(18).
		SetName("peanut_pg").
		// AddCars 将车属于user peanut_pg
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", peanut_pg)
	return peanut_pg, nil
}

// QueryUserByName 通过name 查询
func QueryUserByName(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.NameEQ(name)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

// QueryCars 查询用户是否有Ford 这个车
func QueryCars(ctx context.Context, peanut_pg *ent.User) error {
	cars, err := peanut_pg.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println("returned cars:", cars)

	// what about filtering specific cars.
	ford, err := peanut_pg.QueryCars().
		Where(car.ModelEQ("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println(ford)
	return nil
}

// QueryCarByModel 查询car.model=Tesla
func QueryCarByModel(ctx context.Context, client *ent.Client) (*ent.Car, error) {
	car, err := client.Car.Query().
		Where(car.ModelEQ("Tesla")).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed query car")
	}
	return car, nil
}

// QueryCarUser 查询car.model=Tesla的所属者是谁
func QueryCarUser(ctx context.Context, car *ent.Car) error {
	owner, err := car.QueryOwner().Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying car %q owner:%v", car.Model, err)
	}
	log.Printf("car %q owner: %q\n", car.Model, owner.Name)
	return nil
}

// CreateGraph 创建基础数据
func CreateGraph(ctx context.Context, client *ent.Client) error {
	// first, create the users.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}
	// then, create the cars, and attach them to the users in the creation.
	_, err = client.Car.
		Create().
		SetModel("TeslaY").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("TeslaX").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("TeslaS").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(neta).              // attach this graph to Neta.
		Save(ctx)
	if err != nil {
		return err
	}
	// create the groups, and add their users in the creation.
	_, err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}
