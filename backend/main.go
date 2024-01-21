package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UrlNames struct{
   Url string `json:"url"`
}


func main(){
   fmt.Println("Hello from the moon")
   app := fiber.New()


   app.Get("/", func(c *fiber.Ctx) error {
       return c.SendString("Hello from the moon")
   })

   app.Post("/api", func(c *fiber.Ctx) error {
       body := new(UrlNames)

       if err := c.BodyParser(body); err != nil {
           return err
       }

       log.Println(body.Url)

       return c.SendStatus(fiber.StatusOK)

   })



   app.Listen("localhost:3001")
}

