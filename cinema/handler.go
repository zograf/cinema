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

func (app *App) Login(username, password string) map[string]interface{} {
	result := app.find("Users", bson.D{{Key: "username", Value: username}, {Key: "password", Value: password}})
	if len(result) != 0 {
		return result[0]
	}
	return make(map[string]interface{})
}

func (app *App) GetMovies() []map[string]interface{} {
	result := app.find("Repertoire", bson.D{})
	return result
}

func (app *App) GetUserData(id int) map[string]interface{} {
	result := app.find("Users", bson.D{{Key: "id", Value: id}})
	if len(result) != 0 {
		return result[0]
	}
	return make(map[string]interface{})
}

func CreateApp(database, url string) *App {
	app := App{
		database: database,
	}

	if url == "" {
		app.readConfig()
	} else {
		app.url = url
	}

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

func (app *App) find(collection string, filter bson.D) []map[string]interface{} {
	db := app.client.Database(app.database).Collection(collection)
	cursor, err := db.Find(context.TODO(), filter)
	check(err)
	var result []bson.M
	err = cursor.All(context.TODO(), &result)
	check(err)
	var responses []map[string]interface{}

	for index, el := range result {
		responses = append(responses, make(map[string]interface{}))
		for key, value := range el {
			responses[index][key] = value
		}
	}
	return responses
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
