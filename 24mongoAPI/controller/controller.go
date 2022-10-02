package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Denish-Ranpariya/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://****@cluster0.h4mgl.mongodb.net/?retryWrites=true&w=majority"
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

func updateOneMovie(movieId string) {
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

func deleteAllMovies() int64 {
	filter := bson.M{}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount, " documents deleted successfully with given id.")
	return result.DeletedCount
}

// get all movies from database
func getAllMovies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)

	}
	defer cursor.Close(context.Background())
	return movies
}

//actual controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		log.Fatal(err)
	}

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkMovieAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie marked as watched with id : " + params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Deleted movie with id : " + params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deletedMoviesCount := deleteAllMovies()
	json.NewEncoder(w).Encode(strconv.Itoa(int(deletedMoviesCount)) + " movies deleted")
}
