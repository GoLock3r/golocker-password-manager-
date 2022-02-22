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

// In format
// url: <site url, can be null>
// title: <entry title, can be null>
// username: <username, can be null>
// password: <password, cannot be null>
// other: <self ex. can be null
// creation_date: <date the entry was created>
//

var Loggers *logger.Loggers

var col *mongo.Collection
var ctx = context.TODO()

type entry struct {
	url      string
	title    string
	username string
	password string
	other    string
}

func Connect(collection string) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
		Loggers.LogError.Println("Could not connect to database")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
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
		"other":    ent.other,
		"date":     time.Now().String()}

	result, err := col.InsertOne(ctx, to_insert)

	if err != nil {
		panic(err)
		Loggers.LogError.Println("Could not write entry")
	} else {
		Loggers.LogInfo.Println("Wrote entry", result)
	}
}

func ReadFromTitle(title string) {

}

func ReadFromUsername(username string) {

}

func ReadAll() {

	cursor, err := col.Find(ctx, bson.D{})

	if err != nil {
		panic(err)
		Loggers.LogError.Println("Could not read entries")
	}
	var results []bson.M

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
		Loggers.LogError.Println("cursor error")
	}

	for _, result := range results {
		for _, field := range result {
			fmt.Println(field)
		}
	}
}

func UpdateEntry() {

}

func DeleteEntry() {

}
