package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	ID        primitive.ObjectID `json:"id" binding:"omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" binding:"omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastName" binding:"omitempty" bson:"lastname,omitempty"`
	Age       int                `json:"age,omitempty" binding:"omitempty" bson:"age,omitempty"`
}

// GetUserByID - returns a single user from the db
func (u *User) GetUserByID(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	err := db.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	return err
}

// GetAllUsers - returns all users from the db
func GetAllUsers() ([]User, error) {
	var users []User
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var user User
		_ = cur.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

// CreateUser - creates a single user
func (u *User) CreateUser() error {
	res, err := db.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": res.InsertedID}).Decode(&u)
	return err
}

// UpdateUser - updates a user
func (u *User) UpdateUser(id string) error {
	var updateQuery bson.M
	oid, _ := primitive.ObjectIDFromHex(id)
	data, err := bson.Marshal(u)
	if err != nil {
		return err
	}
	_ = bson.Unmarshal(data, &updateQuery)
	_, err = db.UpdateOne(
		context.TODO(),
		bson.M{"_id": oid},
		bson.D{{"$set", u}},
	)
	if err != nil {
		return err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	return err
}

// DeleteUser - deletes one user from the db
func (u *User) DeleteUser(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	err := db.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	return err
}
