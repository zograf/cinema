package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zograf/cinema/cinema"
	"go.mongodb.org/mongo-driver/bson"
)

var app *cinema.App

func get(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	result := app.Find("cinema", bson.E{"test", "test"})
	c.JSON(http.StatusOK, result)
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
	app = cinema.CreateApp("cinema")

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
