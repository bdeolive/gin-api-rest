package main

import (
	"github.com/bdeolive/gin-api-rest/database"
	"github.com/bdeolive/gin-api-rest/routes"
)

func main() {
	database.ConectaDB()
	routes.HandleRequests()
}
