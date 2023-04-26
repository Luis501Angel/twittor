package db

import (
	"context"
	"time"

	"github.com/luis501angel/go-course/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("twittor")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.Id.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
