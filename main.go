package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbConnect() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://admin:admin@uptest.idmbk.mongodb.net/cinema?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	// ctx is actually irrelevant for querries
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func get(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	client := dbConnect()
	filter := bson.D{{"test", "test"}}
	cursor, err := client.Database("cinema").Collection("cinema").Find(context.TODO(), filter)
	check(err)
	var result []bson.M
	err = cursor.All(context.TODO(), &result)
	check(err)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func put(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{"data": "DUMMY"})
}

func del(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{"data": "DUMMY"})
}

func main() {
	router := gin.Default()

	router.GET("/", get)
	router.POST("/", put)
	router.DELETE("/", del)

	router.Use(cors.Default())
	router.Run()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
