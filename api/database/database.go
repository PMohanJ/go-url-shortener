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

	/* ping, err := rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("Cr db err ", err)
	} else {
		fmt.Println("db created succfully ", ping)
	} */

	return rdb
}
