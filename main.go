package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const (
	uri    = "mongodb+srv://%s:%s@cluster0.yt1pq.mongodb.net/?retryWrites=true&w=majority"
	dbName = "mend"
	dbUser = "tamirm"
	dbPass = "12345"
)

var client *mongo.Client

func main() {
	fmt.Println("Starting the mend-crud app...")
	client, _ = initDB()
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{firstname}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{firstname}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initDB() (*mongo.Client, error) {
	connectionUri := fmt.Sprintf(uri, dbUser, dbPass)
	clientOptions := options.Client().ApplyURI(connectionUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("db connected")
	return client, nil
}
