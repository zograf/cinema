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
	result := app.Find("Test", bson.E{"test", "radi"})
	c.JSON(http.StatusOK, result)
}

func login(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	username := c.Param("username")
	password := c.Param("password")
	result := app.Find("Users", bson.E{Key: "username", Value: username})
	if result["password"] == password {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"bravo": "retarde"})
	}
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
	app = cinema.CreateApp("Cinema")

	router.GET("/", get)
	router.GET("/login/:username/:password", login)
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
