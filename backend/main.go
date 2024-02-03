package main

import (
	"github.com/urlshort/database"
	"github.com/urlshort/routes"
)


func main() {
	database.InitDB()
    routes.InitRouter()
}

