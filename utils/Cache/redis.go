package Cache

import (
	v8 "github.com/go-redis/redis/v8"
)

var RedisClient *v8.Client

func init() {

	RedisClient = v8.NewClient(&v8.Options{
		Addr:     "127.0.0.1:6379",
		Username: "",
		Password: "12345",
		DB:       0,
	})
}
