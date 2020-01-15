package storage

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type SimpleStorage struct {
	path string
	db   *mongo.Client
}

func CreateSimpleStorage() (s Storage, err error) {
	ss := &SimpleStorage{
		path: "mongodb://localhost:27017",
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(ss.path))
	if err != nil {
		logrus.Fatal("Create database error ")
		return nil, err
	}
	ss.db = client
	return ss, err
}

func (ss *SimpleStorage) Insert() {
	collection := ss.db.Database("testing").Collection("handle record")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {

	}
	fmt.Println(res)
	//

}
