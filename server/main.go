package main

import (
	"log"

	"github.com/gofiber/fiber"

	"github.com/pascaloseko/shopping_list/server/database"
	"github.com/pascaloseko/shopping_list/server/router"

	_ "github.com/lib/pq"
)

func main() { // entry point to our program

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New() // call the New() method - used to instantiate a new Fiber App

	// app.Use(middleware.Logger())

	router.SetupRoutes(app)

	app.Listen(3000) // listen/Serve the new Fiber app on port 3000

}
