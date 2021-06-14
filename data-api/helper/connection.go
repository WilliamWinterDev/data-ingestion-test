package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WilliamWinterDev/test-metrics-api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {

	mongoHost := os.Getenv("MONGODB_HOST")

	if mongoHost == "" {
		mongoHost = "localhost:6002"
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + mongoHost)
	fmt.Println("Connecting to " + mongoHost + " mongo database")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("metrics").Collection("cpu_data")
	fmt.Println(collection)

	return collection
}

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = models.ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
