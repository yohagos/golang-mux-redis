package redis

import (
	"context"
	"fmt"
	"time"

	"golang-mux-redis/objects"
	"log"

	redis "github.com/go-redis/redis/v9"
)

var (
	ctx = context.TODO()
	cursor uint64
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
		/* id, err := client.Incr(context.TODO(),"user:next-id").Result()
		if err != nil {
			return
		}
		//key := fmt.Sprintf("user:%id", id)
		pipe := client.Pipeline()
		pipe.HSet(ctx, "id", id)
		pipe.HSet(ctx, "name", v.Name)
		pipe.HSet(ctx, "phone", v.Phone)
		pipe.HSet(ctx, "email", v.Email)
		pipe.HSet(ctx, "user:by-username", v.Name, id)
		_, err = pipe.Exec(ctx)
		if err != nil {
			panic(err)
		}
		log.Println("Saved Init Users successfully") */
		duration, _ := time.ParseDuration("1h")
		if err := client.Set(ctx, "user:"+v.Name, v.Name, duration).Err(); err != nil {
			panic(err)
		}
	}

	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(ctx, cursor, "user:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key ", key)
		}

		if cursor == 0 {
			break
		}
	}
}

