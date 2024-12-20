package main

import (
	"net/http"
	"github.com/cfyyds/fiber-api/database"
	"github.com/cfyyds/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

// 初始化全局 Fiber 应用实例
var app *fiber.App

func init() {
	// 初始化数据库
	if err := database.ConnectDb(); err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// 创建 Fiber 应用实例
	app = fiber.New()

	// 设置路由
	setupRoutes(app)
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)

	// User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)  // Corrected to plural
	app.Delete("/api/users/:id", routes.DeleteUser)  // Corrected to plural

	// Product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)

	// Order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
}

// Vercel 期望的 Handler 函数
func Handler(w http.ResponseWriter, r *http.Request) {
	// 使用 ServeHTTP 使 Fiber 应用符合 http.Handler 接口
	app.Handler()(w, r)  // Fixed: pass w and r to the handler
}
