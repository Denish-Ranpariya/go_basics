package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Denish-Ranpariya/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://*****@cluster0.h4mgl.mongodb.net/test"
const dbName = "Netflix"
const colName = "Watchlist"

var collection *mongo.Collection

//connect with mongoDB

func init() {
	//client options
	clientOptions := options.Client()
	clientOptions.ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MOngoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection reference is ready
	fmt.Println("Collection reference is ready")
}

//mongodb helpers

// insert one record / doc
func insertOneMovie(movie model.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db with id: ", result.InsertedID)
}

func updateOneMovie(movieId string, movie model.Netflix) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.MatchedCount, " document updated successfully with given id.")
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount, " document deleted successfully with given id.")
}

func deleteAllMovie() {
	filter := bson.M{{}}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount, " documents deleted successfully with given id.")
}

func getAllMovies() {

}
