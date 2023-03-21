package main

import (
	"log"

	"github.com/gofiber/fiber/middleware"

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

	app.Use(middleware.Logger())

	router.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	if err := app.Listen(port); err != nil {
		log.Println(err)
	} // listen/Serve the new Fiber app on port 3000
}
