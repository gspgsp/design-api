package log

import (
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"log"
)

type mgo struct {
	database   string
	collection string
}

func NewMgo(database, collection string) *mgo {
	return &mgo{database, collection}
}

func (m *mgo) InsertOne(value interface{}) *mongo.InsertOneResult {
	client := Db.Mongo

	collection := client.Database(m.database).Collection(m.collection)
	insertResult, err := collection.InsertOne(context.TODO(), value)

	if err != nil {
		log.Fatalln("mongodb 插入数据错误:" + err.Error())
	}

	return insertResult
}
