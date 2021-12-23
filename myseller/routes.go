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
	seller.Get("/shop/:id", SingleSellerShops)
	seller.Get("/shops/active/all", AllSellerActiveShops)
	seller.Get("/shops/inactive/all", AllSellerInActiveShops)
	seller.Get("/shops/delete/all", AllSellerDeleteShops)
	seller.Post("/shop/create", CreateShops)
	seller.Put("/shop/update/:id", EditShops)
	seller.Put("/shop/soft/delete/:id", SoftDeleteShops)
	seller.Put("/shop/restore/:id", RestoreShops)
	seller.Delete("/shop/delete/:id", DeleteShops)
	seller.Get("/shop/availability/check", CheckShopAvailability)
	seller.Get("/brand/by/shop",BrandByShop)
	seller.Get("/category/by/shop",CategoryByShop)
	seller.Get("/variation",VariationData)
	//	product
	seller.Get("/product/all", AllSellerProducts)
	seller.Get("/product/all/inactivate", AllInactiveSellerProducts)
	seller.Post("/product/create/:shopID", CreateProduct)
}
