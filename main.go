package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zograf/cinema/cinema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var app *cinema.App

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
	result := app.Find("cinema", bson.E{"test", "test"})
	c.JSON(http.StatusOK, result)
}

func put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "DUMMY"})
}

func del(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "DUMMY"})
}

func main() {
	router := gin.Default()
	client := dbConnect()

	router.GET("/", get)
	router.POST("/", put)
	router.DELETE("/", del)

	app = cinema.CreateApp("cinema", client)
	router.Run()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
