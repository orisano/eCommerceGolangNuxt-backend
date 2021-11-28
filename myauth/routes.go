package myauth

import (
	"bongo/mixin"
	"bongo/model"
	"github.com/gofiber/fiber/v2"
)
func authMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("bongoauth")
	err, data, _ := mixin.VerifyToken(cookie)
	if !err {
		c.Status(fiber.StatusForbidden)
	}
	var user model.User
	result := model.DB.First(&user, "id = ?", data.Issuer)
	if result.Error != nil {
		return c.SendStatus(401)
	}
	c.Locals("AuthID", user.ID)
	return c.Next()
}
func AuthRoutes(app *fiber.App) {
	auth := app.Group("/api/auth")
	auth.Get("/user", GetUser)
	auth.Get("/seller", GetSeller)
	auth.Get("/admin", GetAdmin)
	auth.Get("/user/csrf", GetCSRF)
	auth.Post("/user/register", UserRegister)
	auth.Post("/user/login", UserLogin)
	auth.Post("/user/logout", UserLogout)
	auth.Post("/seller/register", SellerRequestPost)
	auth.Post("/seller/login", SellerLogin)
	auth.Post("/admin/login", AdminLogin)
	auth.Post("/admin/register", AdminRegister)
}
