package log

import (
	"context"
	"design-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type mgo struct {
	database   string
	collection string
}

type SmsMongoInfo struct {
	Id       primitive.ObjectID `bson:"_id"`
	Mobile   string
	Code     string
	CodeKey  string
	ExpireAt int64
}

// NewMgo /**
func NewMgo(collection string) *mgo {
	return &mgo{config.Config.Mongodb.MongodbDatabase, collection}
}

// InsertOne /**
func (m *mgo) InsertOne(value interface{}) *mongo.InsertOneResult {
	insertResult, err := m.GetCollection().InsertOne(context.TODO(), value)

	if err != nil {
		log.Fatalln("mongodb 插入数据错误:" + err.Error())
	}

	return insertResult
}

// GetOne /**
func (m *mgo) GetOne(value interface{}, data interface{}) {
	client := Db.Mongo

	collection := client.Database(m.database).Collection(m.collection)
	err := collection.FindOne(context.Background(), value).Decode(data)
	if err != nil {
		m.collection = "sms_log"
		m.GetCollection().InsertOne(context.TODO(), bson.D{{"findDataErr", err.Error()}})
	}
}

// 更新的 没有弄，直接调用

// GetCollection /**
func (m *mgo) GetCollection() *mongo.Collection {
	client := Db.Mongo
	collection := client.Database(m.database).Collection(m.collection)

	return collection
}
