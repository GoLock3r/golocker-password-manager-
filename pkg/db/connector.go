package db

import (
	"context"
	"fmt"
	"golock3r/server/crypt"
	"golock3r/server/logger"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Loggers *logger.Loggers

var col *mongo.Collection
var ctx = context.TODO()

var URI = "mongodb://localhost:27017"

func Connect(collection string) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))

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

// Anything that is labeled 'password' or contains 'private' will be encrypted
func EncryptEntry(key []byte, entry map[string]string) map[string]string {
	crypt.Loggers = Loggers

	enc_entry := make(map[string]string)

	for k, v := range entry {
		if k == "password" || strings.Contains(k, "private") {
			enc_entry[k] = crypt.EncryptStringToHex(key, v)
		} else {
			enc_entry[k] = v
		}
	}
	return enc_entry
}

// Anything that is labeled 'password' or contains 'private' will be decrypted
func DecryptEntry(key []byte, entry map[string]string) map[string]string {
	crypt.Loggers = Loggers

	dec_entry := make(map[string]string)

	for k, v := range entry {
		if k == "password" || strings.Contains(k, "private") {
			dec_entry[k] = crypt.DecryptStringFromHex(key, v)
		} else {
			dec_entry[k] = v
		}
	}
	return dec_entry
}

// TODO Restate the above two functions, with varargs for 'labels' to encrypt. Encrypt labels if
// Varargs exists, encrypt everything if no varargs are passed
// TODO Functionality to exclude labels?

func WriteEntry(entry map[string]string) {
	to_insert := bson.M{}

	for k, v := range entry {
		to_insert[k] = v
	}
	to_insert["date"] = time.Now().String()

	result, err := col.InsertOne(ctx, to_insert)

	if err != nil {
		Loggers.LogError.Println("Could not write entry")
	} else {
		Loggers.LogInfo.Println("Wrote entry", result)
	}
}

func ReadFromTitle(entryTitle string) []map[string]string {
	filter := bson.D{{Key: "title", Value: entryTitle}}
	cursor, err := col.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var results_map []map[string]string

	for _, result := range results {
		val := make(map[string]string)
		for _, field := range result {
			val[field.Key] = fmt.Sprint(field.Value)
		}
		results_map = append(results_map, val)
	}
	return results_map
}

func ReadFromUsername(entryUsername string) []map[string]string {
	filter := bson.D{{Key: "username", Value: entryUsername}} //found help on mongo db documentation https://docs.mongodb.com/drivers/go/current/fundamentals/crud/query-document/
	cursor, err := col.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var results_map []map[string]string

	for _, result := range results {
		val := make(map[string]string)
		for _, field := range result {
			val[field.Key] = fmt.Sprint(field.Value)
		}
		results_map = append(results_map, val)
	}
	return results_map
}

func ReadAll() []map[string]string {

	cursor, err := col.Find(ctx, bson.D{})

	if err != nil {
		Loggers.LogError.Println("Could not read entries", err)
	}
	var results []bson.D

	if err = cursor.All(ctx, &results); err != nil {
		Loggers.LogError.Println("Cursor error", err)
	}

	var results_map []map[string]string

	for _, result := range results {
		val := make(map[string]string)
		for _, field := range result {
			val[field.Key] = fmt.Sprint(field.Value)
		}
		results_map = append(results_map, val)
	}
	return results_map
}

// Resource used: https://golangdocs.com/mongodb-golang
func UpdateEntry(entryTitle string) {
	filter := bson.D{{Key: "title", Value: entryTitle}}
	var input1, input2 = "", ""
	fmt.Println("Enter the field you would like to update: ")
	fmt.Scanln(&input1)
	fmt.Println("Enter the new value for your chosen field: ")
	fmt.Scanln(&input2)
	input1 = strings.TrimSpace(input1)
	input2 = strings.TrimSpace(input2)

	update := bson.D{{Key: "$set", Value: bson.D{{Key: input1, Value: input2}}}}

	_, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		Loggers.LogError.Println("Entry couldn't be updated", err)
	}
}

func DeleteEntry(entryTitle string) {
	_, err := col.DeleteOne(context.TODO(), bson.D{{Key: "title", Value: entryTitle}})
	if err != nil {
		Loggers.LogError.Println("Entry couldn't be deleted", err)
	} else {
		Loggers.LogInfo.Println("Entry deleted")
	}
}
