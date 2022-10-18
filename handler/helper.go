package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func getSession(c *fiber.Ctx) *session.Session {
	sessStore := c.Locals("session").(*session.Store)
	session, err := sessStore.Get(c)

	if err != nil {
		panic(err)
	} // if()

	return session
} // getSession()

func getFlashMessage(c *fiber.Ctx) string {
	sessStore := c.Locals("session").(*session.Store)
	session, err := sessStore.Get(c)

	if err != nil {
		panic(err)
	} // if()

	msg := session.Get("msg")

	if msg == nil {
		return ""
	} // if()

	session.Delete("msg")

	// Save session
	if err := session.Save(); err != nil {
		panic(err)
	} // if()

	return msg.(string)
} // getFlashMessage()

func setFlashMessage(c *fiber.Ctx, msg string) {
	// check if logged
	sess := getSession(c)
	sess.Set("msg", msg)

	// Save session
	if err := sess.Save(); err != nil {
		panic(err)
	} // if()
} // setFlashMessage()
