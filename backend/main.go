package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"

	//	"fmt"
	"encoding/binary"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/webermarci/base62"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type Url struct{
    Url string `json:"url"`
}
func main(){
    hashUrl("https://codingchallenges.fyi/challenges/challenge-url-shortener/")
   // redirect("6JhyaShUSll")
}

func initRouter(){
    app := fiber.New()
    app.Use(logger.New())

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello World Jones")
    } )


    app.Post("/", func(c *fiber.Ctx) error {
        body := new(Url)
        if err := c.BodyParser(body); err != nil {
            c.Status(fiber.StatusBadRequest).SendString(err.Error())
            return err
        }
        keyOriginal, key, short_url := hashUrl(body.Url)
        log.Println(key)
        log.Println(keyOriginal)
        log.Println(short_url)
        return c.JSON(fiber.Map{"key": key, "long_url": body.Url, "short_url": short_url})
    })

    app.Listen("localhost:3001")
}

func hashUrl(url string)(keyOriginal string, key string, short string){
    h := sha256.New()
    h.Write([]byte(url))
    bs := h.Sum(nil)
    keyBytes := bs[:8]
    fmt.Println( keyBytes)
    uintBytes := binary.LittleEndian.Uint64(keyBytes)

    encoded := base62.Encode(uintBytes)
    encodedShort := encoded[:6]
    keyLower := strings.ToLower(encodedShort)

    shortened := "https://ryan/" + keyLower

    return encoded, keyLower, shortened
}

 func redirect(originalKey  string){
     decoded, _ := base62.Decode(originalKey)

     keyBytes := make([] byte, 8)
     binary.LittleEndian.PutUint64(keyBytes, decoded)

     fmt.Printf("%T", keyBytes)

 }

func initDB() {
 //   var collection *mongo.Collection
    ctx := context.TODO()
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Println(err)
    }
    defer func (){
        if err = client.Disconnect(ctx);err != nil {
            log.Panic(err)
        }
    }()
    err = client.Ping(ctx, nil)
    if err != nil{
        log.Fatal(err)
    }

//    collection = client.Database("NotesPool").Collection("Notes")

   // result, err := collection.InsertMany(ctx, docs)
   // fmt.Printf("Inserted document with _id: %v \n", result.InsertedIDs...)
}
