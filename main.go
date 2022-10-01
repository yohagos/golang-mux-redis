package main

import (
	"golang-mux-redis/objects"
	"golang-mux-redis/redis"
)

func main() {
	objects.UserInit()
/* 	_, err := objects.GetUser("Yosef")
	if err != nil {
		log.Println(err)
	} */

	redis.RedisInit()
}