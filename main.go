package main

import (
	"log"

	"github.com/cfyyds/fiber-api/database"
	"github.com/cfyyds/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to my awesome API")
}

func setupRoutes(app *fiber.App) {
	//welcome endpoint
	app.Get("/api", welcome)
	//User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/user/:id", routes.UpdateUser)
	app.Delete("/api/user/:id", routes.DeleteUser)

	// Product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/product/:id", routes.UpdateProduct)
	app.Delete("/api/product/:id", routes.DeleteProduct)

	// Order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)

}

//	func main() {
//		database.ConnectDb()
//		app := fiber.New()
//		setupRoutes(app)
//		log.Fatal(app.Listen(":3000"))
//	}
//
// Vercel 期望的 `Handler` 函数
func Handler() *fiber.App {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	return app
}

func main() {
	log.Fatal(Handler().Listen(":3000"))
}
