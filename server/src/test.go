package main

import (
    "context"
    "log"
    "time"
    "github.com/mongodb/mongo-go-driver/bson"
    "github.com/mongodb/mongo-go-driver/mongo"
)

type Book struct {
    Title    string
    OwnerId  string
}

func main() {
    client, err := mongo.NewClient("mongodb://localhost:27017")
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("test").Collection("test")

    filter := bson.D{{"title", "math"}}

    var book Book
    err = collection.FindOne(context.TODO(), filter).Decode(&book)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Found a book: %+v\n", book)
}
