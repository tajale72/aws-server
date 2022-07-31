package db

import (
	"context"
	"interview/internal/model"
	"log"
	"os"
	"sort"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// var _ Businesslogic = &Business{}

// type Businesslogic interface {
// 	InsertUser(body []byte) (*mongo.InsertOneResult, error)
// 	GetUser() (model.User, error)
// 	Mongo() (*mongo.Client, error)
// }

// type Business struct {
// 	dbclient *mongo.Client
// }

//InsertUser inserts data into the test database
func InsertUser(user model.User) (*mongo.InsertOneResult, error) {
	sort.Strings(user.Arr)
	dbclient, err := Mongo()
	if err != nil {
		return nil, err
	}
	collection := dbclient.Database("test").Collection("user")
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return res, nil

}

//GetUser get's data from the database
func GetUser(name string) ([]model.User, error) {
	var user model.User
	var a []model.User
	dbclient, err := Mongo()
	if err != nil {
		log.Println("error creating a client", err)
		return nil, err
	}
	collection := dbclient.Database("test").Collection("user")
	err = collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&user)
	a = append(a, user)
	if err != nil {
		log.Println("error getting response from collection client", err)
		return nil, err
	}
	return a, nil
}

func Mongo() (*mongo.Client, error) {
	mongostring := os.Getenv("ATLAS_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongostring))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetAllUser() ([]model.User, error) {
	var users []model.User
	dbclient, err := Mongo()
	if err != nil {
		log.Println("error creating a client", err)
		return nil, err
	}
	collection := dbclient.Database("test").Collection("user")
	a, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("error getting response from collection client", err)
		return nil, err
	}
	if err = a.All(context.Background(), &users); err != nil {
		log.Println(err)
	}
	return users, nil
}
