package database

import (
	"Not-So--Social-Media/model"
	errorUtils "Not-So--Social-Media/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb://pushkarmouryagames:gC8MP1VDJthq6XSz@nssm.kbc7ipw.mongodb.net/?retryWrites=true&w=majority&appName=NSSM"
var dbName = "NSSM"
var collectionName = "User"

var collection *mongo.Collection
var ctx context.Context = context.TODO()

func StartDatabase() {

	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	errorUtils.HandleError(err)

	client.Database(dbName).Collection(collectionName)
	fmt.Println("You successfully connected to MongoDB!")
}

func CreateUser(user model.User) {
	inserted, err := collection.InsertOne(ctx, user)

	errorUtils.HandleError(err)
	fmt.Printf("Inserted Scuccessfully with an Id: %v", inserted.InsertedID)
}
