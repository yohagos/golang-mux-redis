package redis

import (
	"context"
	"fmt"
	"strings"

	"golang-mux-redis/objects"
	"log"

	redis "github.com/go-redis/redis/v9"
)

var (
	ctx = context.TODO()
	client *redis.Client
)

func RedisInit() {
	client = redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
			Password: "", // no password set
        	DB:       0,  // use default DB
		})

	log.Println(client.Ping(ctx))

	users := objects.AllUsers
	for _, v := range users {
		obj := fmt.Sprintf("%v", v)
		if err := client.HSet(ctx, "users", v.Name, obj).Err(); err != nil {
			panic(err)
		}
	}

}

func GetUsers() *[]objects.User{
	var users []objects.User
	x, err := client.HGetAll(ctx, "users").Result()
	if err != nil {
		panic(err)
	}

	for _, v := range x {
		v = strings.Replace(v, "{", "", -1)
		v = strings.Replace(v, "}", "", -1)
		strip := strings.Split(v, " ")
		
		users = append(users, objects.User{Name: strip[0], Email: strip[2], Phone: strip[1]})
	}
	//fmt.Println(users)
	return &users
}

func CreateNewUser(user objects.User) {
	obj := fmt.Sprintf("%v",user)
	if err := client.HSet(ctx, "users", user.Name, obj).Err(); err != nil {
		panic(err)
	}
}