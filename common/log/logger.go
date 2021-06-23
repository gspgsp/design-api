package log

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/go-redis/redis"
	"context"
	"time"
	"log"
	"design-api/config"
)

type NoSqlDataBase struct {
	Mongo *mongo.Client
	Redis *redis.Client
}

var Db *NoSqlDataBase

/**
连接配置常量
 */
const (
	MONGODB_POOL_SIZE = 10
	REDIS_POLL_SIZE   = 10
	MIN_IDLE_CONNS    = 2
)

/**
初始化连接，自动调用
 */
func init() {
	Db = &NoSqlDataBase{
		Mongo: setMongodbConnect("mongodb://" + config.Config.Mongodb.MongodbUsername + ":" + config.Config.Mongodb.MongodbPassword + "@" + config.Config.Mongodb.MongodbHost + ":" + config.Config.Mongodb.MongodbPort),
		Redis: setRedisConnect(),
	}
}

/**
设置mongodb连接pool
 */
func setMongodbConnect(mongoUrl string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl).SetMaxPoolSize(MONGODB_POOL_SIZE))
	if err != nil {
		log.Println("mongo 连接错误" + err.Error())
		return nil
	}

	return client
}

/**
设置redis连接pool
 */
func setRedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         config.Config.Redis.RedisHost + ":" + config.Config.Redis.RedisPort,
		Password:     config.Config.Redis.RedisPassword,
		DB:           config.Config.Redis.RedisDatabase,
		PoolSize:     REDIS_POLL_SIZE,
		MinIdleConns: MIN_IDLE_CONNS,
	})

	//defer client.Close()

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("redis 连接错误" + err.Error())
		return nil
	}

	return client
}
