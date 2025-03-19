package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://Hassaan:policeline12@hassaancluster.qyywv.mongodb.net/?retryWrites=true&w=majority&appName=HassaanCluster").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = client.Disconnect(context.TODO())
		if err != nil {
			panic("Couldnot close connection")
		}
	}()

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

var Client *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, dbName, colName string) *mongo.Collection {
	return client.Database(dbName).Collection(colName)
}
