package helper

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

var ttl time.Duration = time.Duration(5 * time.Minute)

func SetRedis(Rds *redis.Client, key string, data interface{}) bool {
	serialized, _ := json.Marshal(data)

	sr := Rds.Set(context.Background(), key, serialized, ttl)
	if err := sr.Err(); err != nil {
		PanicIfError(err)
	}

	return true
}

func GetRedis(Rds *redis.Client, key string) interface{} {
	tmp := Rds.Get(context.Background(), key)
	res, _ := tmp.Result()

	if res != "" {
		var jsonData = []byte(res)
		var unserialized interface{}
		err := json.Unmarshal(jsonData, &unserialized)
		PanicIfError(err)

		return unserialized
	}

	return res
}
