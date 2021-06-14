package helper

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/WilliamWinterDev/data-ingestion-test/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {

	mongoHost := os.Getenv("MONGODB_HOST")

	if mongoHost == "" {
		mongoHost = "localhost:6002"
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + mongoHost)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("metrics").Collection("cpu_data")

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
