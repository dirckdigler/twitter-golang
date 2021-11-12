package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model of the user in the database.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name, omitempty"`
	Lastname  string             `bson:"lastname" json:"lastname, omitempty"`
	Date      time.Time          `bson:"date" json:"date, omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password, omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar, omitempty"`
	Banner    string             `bson:"banner" json:"banner, omitempty"`
	Biography string             `bson:"biography" json:"biography, omitempty"`
	Location  string             `bson:"location" json:"location, omitempty"`
	Websie    string             `bson:"website" json:"website, omitempty"`
}
