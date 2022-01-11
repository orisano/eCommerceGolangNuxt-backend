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
	seller.Get("/product/min/all", AllSellerProductsMin)
	seller.Get("/product/min/non/active/all", AllSellerNonProductsMin)
	seller.Get("/product/min/deleted/all", AllSellerDeletedProductsMin)
	seller.Get("/product/all/inactivate", AllInactiveSellerProducts)
	seller.Post("/product/create/:shopID", CreateProduct)
	seller.Delete("/product/soft/delete/:id", SoftDeleteProduct)
	seller.Delete("/product/delete/:id", DeleteProduct)
	seller.Put("/product/recover/:id", RecoverProduct)
	seller.Get("/product/:id", SingleProduct)
	seller.Delete("/product/edit/:product_id/:image_id/image/delete", EditProductImageDelete)
	seller.Patch("/product/edit/:product_id/:image_id/image/display", EditProductImageDisplay)
	seller.Post("/product/edit/add/image/:product_id", AddProductImage)
	seller.Patch("/product/edit/basic/:product_id", EditBasicProduct)
	seller.Patch("/product/edit/basic/offer/:product_id", EditBasicOfferProduct)
	seller.Patch("/product/edit/variation/:product_id/:variation_id", EditProductVariation)
	seller.Post("/product/new/variation/:product_id", AddNewProductVariation)
	seller.Delete("/product/delete/variation/:product_id/:variation_id", DeleteProductVariation)
	//	order
	seller.Get("/my/order",MyNewOrder)
	seller.Get("/order/income/statistic",OrderStatistic)
	seller.Get("/checkout/product/details/:id",DetailsCheckoutProduct)
}
