package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbNo,
	})

	/*
		// For debuging purpose
		ping, err := rdb.Ping(Ctx).Result()
		if err != nil {
			fmt.Println("Error while connecting db", err)
		} else {
			fmt.Println("DB created succfully ", ping)
		}
	*/

	return rdb
}
