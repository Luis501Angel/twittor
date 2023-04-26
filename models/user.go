package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	LastName  string             `bson:"lastname" json:"lastname"`
	Birthday  time.Time          `bson:"birthday" json:"birthday"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password, omitempty" json:"password"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	Website   string             `bson:"website" json:"website,omitempty"`
}
