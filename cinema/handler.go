package cinema

import (
	"context"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

type App struct {
	database string
	client   *mongo.Client
	url      string
}

func CreateApp(database string) *App {
	app := App{
		database: database,
	}
	app.readConfig()
	app.dbConnect()
	return &app
}

func (app *App) dbConnect() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(app.url).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	check(err)
	app.client = client
}

func (app *App) Find(collection string, query bson.E) map[string]interface{} {
	db := app.client.Database(app.database).Collection(collection)
	filter := bson.D{query}
	cursor, err := db.Find(context.TODO(), filter)
	check(err)
	var result []bson.M
	err = cursor.All(context.TODO(), &result)
	check(err)
	//fmt.Println("AAAA", result)
	response := make(map[string]interface{})
	for _, el := range result {
		for key, value := range el {
			response[key] = value
		}
	}
	return response
}

func (app *App) readConfig() {
	pathToConfig := "config.yaml"
	yamlFile, err := ioutil.ReadFile(pathToConfig)
	check(err)
	yamlContent := make(map[string]string, 1)
	err = yaml.Unmarshal(yamlFile, &yamlContent)
	check(err)
	app.url = yamlContent["url"]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
