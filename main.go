package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zograf/cinema/cinema"
)

var app *cinema.App
var url string

func getMovies(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	result := app.GetMovies()
	c.JSON(http.StatusOK, result)
}

func login(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	username := c.Param("username")
	password := c.Param("password")
	result := app.Login(username, password)
	c.JSON(http.StatusOK, result)
}

func getUserData(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	idParameter := c.Param("id")
	id, err := strconv.Atoi(idParameter)
	check(err)
	result := app.GetUserData(id)
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
	flag.StringVar(&url, "url", "", "mongodb connection string")
	flag.Parse()

	router := gin.Default()
	app = cinema.CreateApp("Cinema", url)

	router.GET("/login/:username/:password", login)
	router.GET("/user/:id", getUserData)
	router.GET("/movies", getMovies)
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
