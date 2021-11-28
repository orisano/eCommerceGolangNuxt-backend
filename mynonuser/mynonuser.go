package mynonuser

import (
	"bongo/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetShopCategories(c *fiber.Ctx) error {
	var ShopCategories []model.ShopCategory
	err :=model.DB.Select([]string{"id", "name"}).Find(&ShopCategories)
	if err.Error != nil{
		c.SendStatus(204)
	}
	fmt.Println(ShopCategories)
	return c.JSON(ShopCategories)
}
