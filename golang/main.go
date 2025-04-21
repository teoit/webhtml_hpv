package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Get("/", home())
	app.Get("/about", about())
	app.Get("/products/:id<regex(^[a-zA-Z0-9]+$)>", product())

	app.Static("/static", "./public")
	
	fmt.Println("Server is running on http://localhost:3000")

	log.Fatal(app.Listen(":3000"))

}

func home() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
			"Title":  "Home page",
			"name":   "John Doe",
			"age":    30,
			"items":  []string{"item1", "item2", "item3"},
			"header": "This is the header ok",
			"footer": "This is the footer ok ok air",
		})

	}
}

func about() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("About")
	}
}

func product() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		return c.SendString("Product ID: " + id)
	}
}
