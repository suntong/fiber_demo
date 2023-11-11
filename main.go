package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	//Check ENV variables.
	envChecks()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusAlreadyReported).SendString("healthy 👋!")
	})

	// == Basic Routing

	// GET /hi/john
	app.Get("/hi/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	// GET /hi/john/75
	app.Get("/hi/:name/:age", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func envChecks() {
	port, portExist := os.LookupEnv("PORT")

	if !portExist || port == "" {
		log.Fatal("PORT must be set in .env and not empty")
	}
}
