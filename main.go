package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbConnect() (*mongo.Client, context.Context) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://admin:admin@uptest.idmbk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection successful!")
	return client, ctx
}

func get(c *gin.Context) {
	client, context := dbConnect()
	cursor, _ := client.Database("cinema").Collection("cinema").Find(context, bson.D{})
	var results []bson.M
	_ = cursor.All(context, &results)
	fmt.Println(results)
	c.JSON(http.StatusOK, gin.H{"data": "DUMMY"})
}

func put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "DUMMY"})
}

func del(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "DUMMY"})
}

func main() {
	router := gin.Default()

	router.GET("/", get)
	router.POST("/", put)
	router.DELETE("/", del)

	router.Run()
}
