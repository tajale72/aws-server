package controller

import (
	"encoding/json"
	"errors"
	"interview/internal/db"
	"interview/internal/model"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// type Service struct {
// 	db db.Business
// }

//GetUser function handles the business logic for getting a user.
func GetUser(name string) ([]model.User, error) {
	log.Println("Get user from db")
	return db.GetUser(name)
}

//GetUser function handles the business logic for getting a user.
func InsertUser(body []byte) (*mongo.InsertOneResult, error) {
	log.Println("insert user into the db")
	var user model.User
	json.Unmarshal(body, &user)
	if user.Name == "" {
		log.Println()
		return nil, errors.New("please enter the name")
	} else {
		return db.InsertUser(user)
	}
}

func GetAllUser() ([]model.User, error) {
	return db.GetAllUser()
}
