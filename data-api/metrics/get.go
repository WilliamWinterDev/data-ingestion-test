package metrics

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/WilliamWinterDev/test-metrics-api/helper"
	"github.com/WilliamWinterDev/test-metrics-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Get(w http.ResponseWriter, r *http.Request) {
	collection := helper.ConnectDB()

	w.Header().Set("Content-Type", "application/json")

	// we created Metrics array
	var metrics []models.Metric
	var params = mux.Vars(r)

	filter := getFilter(params)

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var metric models.Metric

		err := cur.Decode(&metric) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		metrics = append(metrics, metric)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(metrics) // encode the result and output that in json format to the user

}

func getFilter(params map[string]string) primitive.M {
	startTime, _ := strconv.Atoi(params["startTime"])
	endTime, _ := strconv.Atoi(params["endTime"])

	if startTime > 0 && endTime > 0 {
		fmt.Println("Filtering down to specific timestamps")
		return bson.M{"timestamp": bson.M{
			"$gte": startTime,
			"$lt":  endTime,
		}}
	} else {
		fmt.Println("Returning all metrics")
		return bson.M{}
	}
}
