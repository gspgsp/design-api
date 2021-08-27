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
	Id      primitive.ObjectID `bson:"_id"`
	Mobile  string
	Code    string
	CodeKey string
}

func NewMgo(collection string) *mgo {
	return &mgo{config.Config.Mongodb.MongodbDatabase, collection}
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

func (m *mgo) GetOne(value interface{}, data interface{}) {
	client := Db.Mongo

	collection := client.Database(m.database).Collection(m.collection)
	err := collection.FindOne(context.Background(), value).Decode(data)
	if err != nil {
		client.Database(m.database).Collection("sms_log").InsertOne(context.TODO(), bson.D{{"findDataErr", err.Error()}})
	}
}

func (m *mgo) UpdateOne(value interface{}, new interface{}) {
	client := Db.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	n, ok := new.(SmsMongoInfo)
	if ok {
		collection.UpdateOne(context.Background(), bson.M{"mobile": n.Mobile}, new)
	}
}
