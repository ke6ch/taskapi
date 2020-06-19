package lib

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		// Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// GetSession セッション情報を取得する
func GetSession(UUID string) error {
	val, err := client.Get(UUID).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
		return err
	} else if err != nil {
		return err
	} else {
		fmt.Println("key", val)
		return nil
	}
}

// SetSession Key: UUID, value: UserIDとしてredisにセッション情報を登録する
func SetSession(UUID string, userID string) error {
	err := client.Set(UUID, userID, 3*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	return nil
}
