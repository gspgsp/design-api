package log

import "github.com/go-redis/redis"

func NewRd(database int) *redis.Client {
	client := Db.Redis
	client.Do("select", database)

	return client
}
