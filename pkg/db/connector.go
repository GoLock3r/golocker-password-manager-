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

func Connect(collection string) bool {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))

	if err != nil {
		Loggers.LogError.Println("Could not connect to database")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		Loggers.LogError.Println("Could not ping database")
		return false
	} else {
		Loggers.LogInfo.Println("Connected to database")
	}
	col = client.Database("golocker").Collection(collection)
	return true
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

// // Encrypt entries denoted by 'include_keys' variadict. If 'include_keys' is empty, default to above
// // function's functionality (encrypt / decrypt anything that is labeled with 'password' or contains 'private')
// // TODO replace above functions with these
// func EncryptEntry(key []byte, entry map[string]string, include_keys ...string) map[string]string {

// }

// // Decrypt entries denoted by 'include_keys' variadict.
// func DecryptEntry(key []byte, entry map[string]string, include_keys ...string) map[string]string {

// }

func WriteEntry(entry map[string]string) bool {
	to_insert := bson.M{}

	for k, v := range entry {
		to_insert[k] = v
	}
	to_insert["date"] = time.Now().String()

	result, err := col.InsertOne(ctx, to_insert)

	if err != nil {
		Loggers.LogError.Println("Could not write entry")
		return false
	} else {
		Loggers.LogInfo.Println("Wrote entry", result)
		return true
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
func UpdateEntry(entryTitle string, updateField string, updateEntry string) bool {
	filter := bson.D{{Key: "title", Value: entryTitle}}
	var input1, input2 = "", ""
	// fmt.Println("Enter the field you would like to update: ")
	// fmt.Scanln(&input1)
	// fmt.Println("Enter the new value for your chosen field: ")
	// fmt.Scanln(&input2)
	input1 = updateField
	input2 = updateEntry
	input1 = strings.TrimSpace(input1)
	input2 = strings.TrimSpace(input2)

	update := bson.D{{Key: "$set", Value: bson.D{{Key: input1, Value: input2}}}}

	_, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		Loggers.LogError.Println("Entry couldn't be updated", err)
		return false
	}
	return true
}

func DeleteEntry(entryTitle string) bool {
	_, err := col.DeleteOne(context.TODO(), bson.D{{Key: "title", Value: entryTitle}})
	if err != nil {
		Loggers.LogError.Println("Entry couldn't be deleted", err)
		return false
	} else {
		Loggers.LogInfo.Println("Entry deleted")
		return true
	}
}
<<<<<<< HEAD
func RemoveAll() bool {
	allEntries := ReadAll()

	for i,entry := range allEntries {
	
		deleteEntry :=	DeleteEntry(allEntries[i][entry["title"]])
		if !deleteEntry {
			return false
		}
=======

func RemoveAll() bool {
	_, err := col.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		Loggers.LogError.Println("Entries not deleted", err)
		return false
	} else {
		Loggers.LogInfo.Println("Entries deleted")
		return true
>>>>>>> 60c0840da43b54cbee3bfd3caeb2dc64ee7ff8ae
	}
	// allEntries := ReadAll()

<<<<<<< HEAD
	return true
}
=======
	// for i, entry := range allEntries {

	// 	deleteEntry := DeleteEntry(allEntries[i][entry["title"]])
	// 	if !deleteEntry {
	// 		return false
	// 	}
	// }
}
>>>>>>> 60c0840da43b54cbee3bfd3caeb2dc64ee7ff8ae
