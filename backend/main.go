package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Book struct{
    Title string
    Author string
}

func main() {
    var collection *mongo.Collection
    ctx := context.TODO()
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Println(err)
    }
    defer func (){
        if err = client.Disconnect(ctx);err != nil {
            log.Panic(err)
        }
    }()
    err = client.Ping(ctx, nil)
    if err != nil{
        log.Fatal(err)
    }

    collection = client.Database("NotesPool").Collection("Notes")
    doc := Book{
        Title: "Harry Potter",
        Author: "JK Rowling",
    }

    result, err := collection.InsertOne(ctx, doc)
    fmt.Printf("Inserted document with _id: %v \n", result.InsertedID)
}
