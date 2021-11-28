package mynonuser

import "github.com/gofiber/fiber/v2"

func NonAuthRoutes(app *fiber.App) {
	nonuser := app.Group("/api/nonuser")
	nonuser.Get("/shop/categories/all", GetShopCategories)
}

