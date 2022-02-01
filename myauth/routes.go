package myauth

import (
	"bongo/db"
	"bongo/mixin"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)
func authMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("bongoauth")
	err, data, _ := mixin.VerifyToken(cookie)
	if !err {
		c.Status(fiber.StatusForbidden)
	}
	fmt.Println("id")
	id, _ := strconv.Atoi(data.Issuer)
	_, err2 := db.Client.User.Get(context.Background(),id)
	if err2 != nil {
		return c.SendStatus(401)
	}
	c.Locals("AuthID", id)
	fmt.Println(id)
	return c.Next()
}
func AuthRoutes(app *fiber.App) {
	nonAuth := app.Group("/api/auth")
	nonAuth.Get("/user", GetUser)
	nonAuth.Get("/seller", GetSeller)
	nonAuth.Get("/admin", GetAdmin)
	nonAuth.Get("/user/csrf", GetCSRF)
	nonAuth.Post("/user/register", UserRegister)
	nonAuth.Post("/user/login", UserLogin)
	nonAuth.Post("/user/logout", UserLogout)
	nonAuth.Post("/seller/register", SellerRequestPost)
	nonAuth.Post("/seller/login", SellerLogin)
	nonAuth.Post("/admin/login", AdminLogin)
	nonAuth.Post("/admin/register", AdminRegister)
}
