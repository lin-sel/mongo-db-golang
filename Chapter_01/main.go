package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Developer Details
type Developer struct {
	Fisrtname      string `bson:"firstname,omitempty"`
	LastName       string `bson:"lastname,omitempty"`
	Profession     string `bson:"profession,omitempty"`
	Responsibility string `bson:"responsibility,omitempty"`
}

func main() {

	Alex := Developer{"Alex", "Gorgia", "Engineer", "Database Developer"}
	Jaq := Developer{"Jauqline", "Farnandiz", "Engineer", "Backend Developer"}
	Jack := Developer{"Jack", "Far", "Engineer", "Frontend Developer"}

	// Creating Instance of Url
	createOption := options.Client().ApplyURI("mongodb://localhost:27017")

	// Create Connection.
	conn, err := mongo.Connect(context.TODO(), createOption)

	if err != nil {
		log.Fatal(err)
	}

	// Connection Cheking.
	err = conn.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// Selecting Database and Collection.
	collection := conn.Database("test").Collection("Programmer")

	// Inserting one Data At a time.
	_, er := collection.InsertOne(context.TODO(), Alex)

	if er != nil {
		log.Fatal(er)
	}

	// Multiple Parameter Together.
	mul := []interface{}{Jaq, Jack}

	// Inserting Many Data Together.
	_, er = collection.InsertMany(context.TODO(), mul)

	if er != nil {
		log.Fatal(er)
	}

	// updating Data.
	filter := bson.D{
		{"firstname", "Jack"},
	}

	// updating value
	val := bson.D{
		{"$set", bson.D{
			{"responsibility", "Project Manager"},
		}},
	}

	// Command used to upadte value.
	updateResult, erro := collection.UpdateOne(context.TODO(), filter, val)

	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println(updateResult.MatchedCount, updateResult.ModifiedCount)
	}

	// Finding Document From Database.
	var result Developer

	errs := collection.FindOne(context.TODO(), bson.D{{"firstname", "Jack"}}).Decode(&result)

	if errs != nil {
		fmt.Println(errs)
	} else {
		fmt.Println(result)
	}

	// Delete Document From Collection.
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// At the end Connection will Close.
	defer func() {
		er = conn.Disconnect(context.TODO())
		fmt.Println("Disconnected Successfully")
	}()
}
