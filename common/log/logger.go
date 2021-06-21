package log

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
	"log"
)

type NoSqlDataBase struct {
	Mongo *mongo.Client
}

var Db *NoSqlDataBase

/**
初始化连接
 */
func Init(mongoUrl string) {
	Db = &NoSqlDataBase{
		Mongo: setMongodbConnect(mongoUrl),
	}
}

/**
设置mongodb连接pool
 */
func setMongodbConnect(mongoUrl string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl).SetMaxPoolSize(10))
	if err != nil {
		log.Println("mongo 连接错误" + err.Error())
	}

	return client
}
