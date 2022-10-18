package main

import (
	"net/http"

	"cashier-admin/handler"
	"cashier-admin/web/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/django"
)

func main() {
	// Create a new view engine
	engine := django.NewFileSystem(http.FS(template.Contents), ".twig").Debug(true)

	// create session object
	sess := session.New()

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	app.Static("/public", "./public") // 靜態的 因為背景css的東西不需要登入才能使用

	// inject session to request context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("session", sess)
		return c.Next()
	})

	// realtime paths
	app.Get( "/demo/show", func( c *fiber.Ctx ) error{
		return c.Render( "demo_newboard", fiber.Map{
			})
	} )

	app.Get("/user/login", handler.ShowUserLoginForm) // show the page for this URL (網址)
	app.Post("/user/login", handler.HandleUserLogin)
	app.Get("/user/logout", handler.HandleUserLogout)
	app.Get("/user/changePassword", handler.ShowUserChangePasswordForm)
	app.Post("/user/changePassword", handler.ChangeUserPassword)
	app.Use(handler.CheckUserAuth())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/user/dashboard")
	})

	app.Get("/user/dashboard", handler.ShowUserInfoList)
	// app.Get( "/user/order", handler.ShowUserOrderForm )
	app.Get("/user/order", handler.ShowUserOrder)
	app.Get("/user/order/:pageId", handler.ShowUserOrder1)

	app.Listen(":3020")
} // main()
