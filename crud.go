package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

const tableName = "people"

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func GetPerson(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Inside GetPerson ...")
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection := client.Database(dbName).Collection(tableName)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func GetPeople(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []Person
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection := client.Database(dbName).Collection(tableName)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func CreatePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := client.Database(dbName).Collection(tableName)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func UpdatePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	params := mux.Vars(request)
	name, _ := params["firstname"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"firstname", name}}
	update := bson.M{"$set": bson.M{"lastname": person.Lastname}}
	collection := client.Database(dbName).Collection(tableName)
	result, _ := collection.UpdateOne(ctx, filter, update)
	json.NewEncoder(response).Encode(result)
}

func DeletePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	name, _ := params["firstname"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"firstname", name}}
	collection := client.Database(dbName).Collection(tableName)
	result, _ := collection.DeleteOne(ctx, filter)
	json.NewEncoder(response).Encode(result)
}
