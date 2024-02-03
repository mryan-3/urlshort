package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/urlshort/database"
	"github.com/webermarci/base62"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strings"
)

type PostUrl struct {
	Url string `json:"url"`
}

type StoredUrl struct {
	Key      string `json:"key"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}

func CreateUser(c *fiber.Ctx) error {
	var url PostUrl
	json.Unmarshal([]byte(c.Body()), &url)

	key, short_url, long_url := HashUrl(url.Url)

	storedUrl := StoredUrl{
		Key:      key,
		ShortUrl: short_url,
		LongUrl:  long_url,
	}

	collection, err := database.GetMongoDbCollection("JsonParser", "Go")
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
}
func GetUser(c *fiber.Ctx) error {
	collection, err := database.GetMongoDbCollection("JsonParser", "Go")
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

	if err != nil {
		log.Println(err)
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
	}

	json, _ := json.Marshal(results)
	return c.Send(json)

}

func HashUrl(url string) (key string, short string, original string) {
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
