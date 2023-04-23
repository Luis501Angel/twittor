package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConn = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://luis501angel:<password>@mycluster.4nrs7yu.mongodb.net/test")

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa en la BD")
	return client
}

func CheckConnection() bool {
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
