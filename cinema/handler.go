package cinema

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	database string
	client   *mongo.Client
}

func CreateApp(database string, client *mongo.Client) *App {
	app := App{
		database: database,
		client:   client,
	}
	return &app
}

func (app *App) Find(collection string, query bson.E) map[string]interface{} {
	db := app.client.Database(app.database).Collection(collection)
	filter := bson.D{query}
	cursor, err := db.Find(context.TODO(), filter)
	check(err)
	var result []bson.M
	err = cursor.All(context.TODO(), &result)
	check(err)
	fmt.Println("AAAA", result)
	response := make(map[string]interface{})
	for _, el := range result {
		for key, value := range el {
			response[key] = value
		}
	}
	return response
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
