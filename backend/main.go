package main

import (
	"context"
    "strings"
	"crypto/sha256"
	"fmt"
	"log"
    "encoding/binary"
    "github.com/webermarci/base62"
	"github.com/gofiber/fiber/v2"
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
    hashUrl("https://www.youtube.com/watch=?jhhfsaf")
}

func initRouter(){
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello World Jones")
    } )


    app.Post("/", func(c *fiber.Ctx) error {
        body := new(Url)
        if err := c.BodyParser(body); err != nil {
            c.Status(fiber.StatusBadRequest).SendString(err.Error())
            return err
        }
       return nil
    })

    app.Listen("localhost:3001")
}

func hashUrl(url string){
    h := sha256.New()
    h.Write([]byte(url))
    bs := h.Sum(nil)
    keyBytes := bs[:8]
    fmt.Println(keyBytes)

    uintBytes := binary.LittleEndian.Uint64(keyBytes)

    encoded := base62.Encode(uintBytes)
    encoded = encoded[:6]
    key := strings.ToLower(encoded)

    short_url := "https://ryan/" + key
    fmt.Println(short_url)
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
