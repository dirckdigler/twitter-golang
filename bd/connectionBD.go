package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is a connection object to the DB.
var MongoCN = ConnectionDB()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://root:master123@cluster0.o29yt.mongodb.net/twitter?retryWrites=true&w=majority")

// ConnectionDB provides connection with the data base.
func ConnectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Succesfully connection to MongoDB")
	return client
}

// CheckConnection is a ping to the database.
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
