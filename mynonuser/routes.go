package mynonuser

import (
	"bongo/mixin"
	"bongo/model"
	"github.com/gofiber/fiber/v2"
)
func authMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("bongoauth")
	err, data, _ := mixin.VerifyToken(cookie)
	if !err {
		c.Status(fiber.StatusUnauthorized)
	}
	var user model.User
	result := model.DB.First(&user, "id = ?", data.Issuer)
	if result.Error != nil {
		return c.SendStatus(401)
	}
	c.Locals("user_id", user.ID)
	return c.Next()
}
func NonAuthRoutes(app *fiber.App) {
	nonuser := app.Group("/api/nonuser")
	user := app.Group("/api/user",authMiddleware)
	nonuser.Get("/shop/categories/all", GetShopCategories)
	// frontend user

	nonuser.Get("/all/main/categories",AllProductCategories)
	nonuser.Get("/single/categories/all/product/:categorySlug/:categoryID",AllProductByCategories)
	nonuser.Get("/all/products",AllProducts)
	nonuser.Get("/single/products/:slug/:id",SingleProducts)
	nonuser.Get("/cart/product/:productID/:variationID",GetCartProduct)

	user.Post("/cart/localstorage",CartStorageProducts)
	user.Post("/cart/product/one",CartProductOne)
	user.Get("/cart/count",GetCountCart)
	user.Get("/cart/product/all",GetCartProductAll)
	user.Delete("/cart/product/remove/:productID",CartUserRemoveProduct)
}

