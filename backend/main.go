package main

import (
	"context"
	"crypto/sha256"

	//	"fmt"
	"strings"

	//	"fmt"
	"encoding/binary"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/webermarci/base62"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostUrl struct {
	Url string `json:"url"`
}

type StoredUrl struct {
	Key      string `json:"key"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}

func main() {
	initDB()
	initRouter()
	// redirect("6JhyaShUSll00")
}

func initRouter() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/:id?", func(c *fiber.Ctx) error {
		collection, err := getMongoDbCollection("JsonParser", "Go")
		if err != nil {
			log.Println(err)
		}

		var filter bson.M = bson.M{}

		if c.Params("id") != "" {
			id := c.Params("id")

			objID, _ := primitive.ObjectIDFromHex(id)
			filter = bson.M{"_id": objID}
		}

        var results []bson.M
        cur, err := collection.Find(context.Background(), filter)
        defer cur.Close(context.Background())

        if err != nil{
            log.Println(err)
        }

        cur.All(context.Background(), &results)

        if results == nil {
            c.SendStatus(404)
        }

        json, _ := json.Marshal(results)
        return c.Send(json)

	})

	app.Post("/", func(c *fiber.Ctx) error {
		var url PostUrl
		json.Unmarshal([]byte(c.Body()), &url)

		key, short_url, long_url := hashUrl(url.Url)

		storedUrl := StoredUrl{
			Key:      key,
			ShortUrl: short_url,
			LongUrl:  long_url,
		}

		collection, err := getMongoDbCollection("JsonParser", "Go")
		if err != nil {
			log.Println(err)
		}
		res, err := collection.InsertOne(context.TODO(), storedUrl)
		log.Println(res)
		if err != nil {
			log.Println(err)
		}
		response, _ := json.Marshal(res)
		return c.Send(response)
	})

	app.Listen("localhost:3001")
}

func createUrl(c *fiber.Ctx) {

	collection, err := getMongoDbCollection("JSONParser", "Go")
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
	}
	var url PostUrl
	json.Unmarshal([]byte(c.Body()), &url)

	keyOriginal, key, short_url := hashUrl(url.Url)
	var storedUrl StoredUrl
	storedUrl.Key = key
	storedUrl.ShortUrl = short_url
	storedUrl.LongUrl = keyOriginal

	res, err := collection.InsertOne(context.TODO(), storedUrl)
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
	}
	response, _ := json.Marshal(res)
	c.Send(response)

}
func hashUrl(url string) (key string, short string, original string) {
	h := sha256.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	keyBytes := bs[:8]
	//    fmt.Println( keyBytes)
	uintBytes := binary.LittleEndian.Uint64(keyBytes)

	encoded := base62.Encode(uintBytes)
	encodedShort := encoded[:6]
	keyLower := strings.ToLower(encodedShort)

	shortened := "https://ryan/" + keyLower

	return keyLower, shortened, url
}

func initDB() (*mongo.Client, error) {
	//   var collection *mongo.Collection
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client, err

	//    collection = client.Database("NotesPool").Collection("Notes")

	// result, err := collection.InsertMany(ctx, docs)
	// fmt.Printf("Inserted document with _id: %v \n", result.InsertedIDs...)
}

func getMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := initDB()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)
	return collection, nil
}
