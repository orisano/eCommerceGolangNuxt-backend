package myseller

import (
	"bongo/mixin"
	"bongo/model"
	"github.com/gofiber/fiber/v2"
)

func sellerMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("bongoauth")
	err, data, _ := mixin.VerifyToken(cookie)
	if !err {
		c.Status(fiber.StatusForbidden)
	}

	var user model.User
	result := model.DB.First(&user, "id = ?", data.Issuer)
	if result.Error != nil || !user.Seller {
		return c.SendStatus(401)
	}
	c.Locals("AuthID", user.ID)
	return c.Next()
}
func SellerRoutes(app *fiber.App) {
	seller := app.Group("/api/seller", sellerMiddleware)
	seller.Get("/shop/action", AllSellerShops)
	seller.Post("/shop/create", CreateShops)
	seller.Get("/shop/availability/check", CheckShopAvailability)
}