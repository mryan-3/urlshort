package main

import (
	"context"
    "strings"
	"crypto/sha256"
//	"fmt"
	"log"
    "encoding/binary"
    "github.com/webermarci/base62"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type ReturnedUrl struct{
    Key string `json:"key"`
    Short_Url string `json:"short_url"`
    Full_Url string `json:"full_url"`
}

type Url struct{
    Url string `json:"url"`
}
func main(){
    initRouter()
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
        key, short_url := hashUrl(body.Url)
        log.Println(key)

        log.Println(short_url)
       return c.SendString(body.Url)
    })

    app.Listen("localhost:3001")
}

func hashUrl(url string)(key string, short string){
    h := sha256.New()
    h.Write([]byte(url))
    bs := h.Sum(nil)
    keyBytes := bs[:8]

    uintBytes := binary.LittleEndian.Uint64(keyBytes)

    encoded := base62.Encode(uintBytes)
    encoded = encoded[:6]
    keyLower := strings.ToLower(encoded)

    shortened := "https://ryan/" + keyLower

    return keyLower, shortened
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
