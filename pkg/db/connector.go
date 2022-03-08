package db

import (
	"context"
	"fmt"
	"golock3r/server/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Loggers *logger.Loggers

var col *mongo.Collection
var ctx = context.TODO()

type entry struct {
	url      string
	title    string
	username string
	password string
	notes    string
}

func Connect(collection string) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		Loggers.LogError.Println("Could not connect to database")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		Loggers.LogError.Println("Could not ping database")
	} else {
		Loggers.LogInfo.Println("Connected to database")
	}
	col = client.Database("golocker").Collection(collection)
}

func CreateEntry(url string, title string, username string, password string, other string) entry {
	var ent = entry{url, title, username, password, other}
	return ent
}

func WriteEntry(ent entry) {
	to_insert := bson.M{
		"url":      ent.url,
		"title":    ent.title,
		"username": ent.username,
		"password": ent.password,
		"notes":    ent.notes,
		"date":     time.Now().String()}

	result, err := col.InsertOne(ctx, to_insert)

	if err != nil {
		Loggers.LogError.Println("Could not write entry")
	} else {
		Loggers.LogInfo.Println("Wrote entry", result)
	}
}

func ReadFromTitle(title string) {

}

// func ReadFromUsername(username string) {
// 	filter := bson.D{{"username", username}} //found help on mongo db documentation https://docs.mongodb.com/drivers/go/current/fundamentals/crud/query-document/
// 	cursor, err := col.Find(context.TODO(), filter)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var results []bson.D
// 	if err = cursor.All(context.TODO(), &results); err != nil {
// 		panic(err)
// 	}
// 	for _, result := range results {
// 		fmt.Println(result)
// 	}

// 	for _, result := range results {
// 		for _, field := range result {
// 			fmt.Println(field)
// 		}
// 		fmt.Println("")
// 	}

// }

func ReadAll() {

	cursor, err := col.Find(ctx, bson.D{})

	if err != nil {
		panic(err)
		Loggers.LogError.Println("Could not read entries")
	}
	var results []bson.D

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
		Loggers.LogError.Println("cursor error")
	}

	for _, result := range results {
		for _, field := range result {
			fmt.Println(field)
		}
		fmt.Println("")
	}
}

func UpdateEntry() {

}

func DeleteEntry() {

}
