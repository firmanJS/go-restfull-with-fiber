package main

import (
	"log"
	"github.com/firmanjs/go-restfull-with-fiber/database"
	"github.com/firmanjs/go-restfull-with-fiber/router"
	"github.com/gofiber/fiber/v2"
  	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() { // entry point to our program

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New() // call the New() method - used to instantiate a new Fiber App

	app.Use(logger.New()

	router.SetupRoutes(app)

	app.Listen(3000) // listen/Serve the new Fiber app on port 3000

}
