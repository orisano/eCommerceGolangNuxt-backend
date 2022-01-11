package mynonuser

import (
	"bongo/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GetShopCategories(c *fiber.Ctx) error {
	var ShopCategories []model.ShopCategory
	err := model.DB.Select([]string{"id", "name"}).Find(&ShopCategories)
	if err.Error != nil {
		return c.SendStatus(204)
	}
	fmt.Println(ShopCategories)
	return c.JSON(ShopCategories)
}

func AllProductCategories(c *fiber.Ctx) error {
	var categories []model.Category
	model.DB.Select([]string{"id", "name", "slug"}).Where("parent_id IS NULL").Find(&categories)
	return c.JSON(categories)
}
func AllProductByCategories(c *fiber.Ctx) error {
	var products []model.SellerProduct
	var ProductIDs []int64
	if err := model.DB.Model(&model.SellerProductCategory{}).Where("category_id = ?", c.Params("categoryID")).Pluck("seller_product_id", &ProductIDs); err.Error != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}
	if len(ProductIDs) > 0 {
		model.DB.Select([]string{"id", "name", "slug", "selling_price", "product_price", "offer_price", "offer_price_start", "offer_price_end"}).Preload("SellerProductImage", "display = (?)", true).Find(&products, &ProductIDs)
		return c.JSON(products)
	}
	return c.SendStatus(fiber.StatusNoContent)

}

func AllProducts(c *fiber.Ctx) error {
	var products []model.SellerProduct
	model.DB.Select([]string{"id", "name", "slug", "selling_price", "product_price", "offer_price", "offer_price_start", "offer_price_end"}).Preload("SellerProductImage", "display = (?)", true).Find(&products)
	return c.JSON(products)
}
func SingleProducts(c *fiber.Ctx) error {
	var product model.SellerProduct
	model.DB.Select([]string{"id", "name", "slug", "selling_price", "product_price", "offer_price", "offer_price_start", "offer_price_end", "quantity", "next_stock", "description"}).Preload("SellerProductImage", func(db *gorm.DB) *gorm.DB { return db.Select([]string{"image", "seller_product_id"}) }).Preload("SellerProductVariation", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "image", "product_price", "selling_price", "quantity", "seller_product_id"})
	}).Preload("SellerProductVariation.SellerProductVariationValues", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name", "description", "attribute_id", "seller_product_variation_id"})
	}).Preload("SellerProductVariation.SellerProductVariationValues.Attribute", func(db *gorm.DB) *gorm.DB { return db.Select([]string{"id", "name"}) }).Find(&product, c.Params("id"))
	return c.JSON(product)
}
func GetCountCart(c *fiber.Ctx) error {

	var cart model.Cart
	if err := model.DB.Where("user_id = ?", c.Locals("user_id")).First(&cart); err.Error != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}
	var count int64
	if err := model.DB.Model(model.CartProduct{}).Where("cart_id = ?", cart.ID).Count(&count); err.Error != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}
	return c.JSON(count)
}

func GetCartProduct(c *fiber.Ctx) error {
	var product model.SellerProduct
	variationID := c.Params("variationID")
	if variationID == "" {
		model.DB.Select([]string{"id", "name", "slug", "selling_price", "product_price", "offer_price", "offer_price_start", "offer_price_end", "quantity", "next_stock", "description"}).Preload("SellerProductImage", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"image", "seller_product_id"}).Where("display = ?", true)
		}).Find(&product, c.Params("productID"))
		return c.JSON(product)
	}
	model.DB.Select([]string{"id", "name", "slug", "selling_price", "product_price", "offer_price", "offer_price_start", "offer_price_end", "quantity", "next_stock", "description"}).Preload("SellerProductImage", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"image", "seller_product_id"}).Where("display = ?", true)
	}).Preload("SellerProductVariation", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "image", "product_price", "selling_price", "quantity", "seller_product_id"}).Where("id = ?", c.Params("variationID"))
	}).Preload("SellerProductVariation.SellerProductVariationValues", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name", "description", "attribute_id", "seller_product_variation_id"})
	}).Preload("SellerProductVariation.SellerProductVariationValues.Attribute", func(db *gorm.DB) *gorm.DB { return db.Select([]string{"id", "name"}) }).Find(&product, c.Params("productID"))
	return c.JSON(product)
}

func CartStorageProducts(c *fiber.Ctx) error {
	type cartsData struct {
		ProductID   uint  `json:"product_id"`
		VariationID *uint `json:"variation_id"`
		Quantity    int   `json:"quantity"`
	}
	var carts []cartsData

	errE := json.Unmarshal(c.Body(), &carts)
	if errE != nil {
		return errE
	}
	var user model.User
	if err := model.DB.Find(&user, c.Locals("user_id")); err.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var cartVal model.Cart
	tx := model.DB.Begin()
	err := tx.Attrs(model.Cart{Slug: fmt.Sprintf("%s-%d-%d", strings.ToLower(user.Name), rand.Intn(9999), user.ID)}).FirstOrCreate(&cartVal, model.Cart{UserID: user.ID})
	if err.Error != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	for _, cart := range carts {
		var product model.SellerProduct
		if err := model.DB.First(&product, cart.ProductID); err.Error != nil {
			tx.Rollback()
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}

		var productVariation model.SellerProductVariation
		if cart.VariationID != nil {
			if err := model.DB.Where("seller_product_id = ?", product.ID).First(&productVariation, "id = ?", cart.VariationID); err.Error != nil {
				fmt.Println(err)
				tx.Rollback()
				return c.SendStatus(fiber.StatusUnprocessableEntity)
			}
		}
		if cart.Quantity < 1 {
			tx.Rollback()
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}
		var cartProduct model.CartProduct
		err := model.DB.Where("cart_id = ?", cartVal.ID).Where("seller_product_id = ?", cart.ProductID).First(&cartProduct)
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			cartProductCreate := model.CartProduct{CartID: cartVal.ID, Quantity: cart.Quantity, SellerProductID: cart.ProductID, SellerProductVariationID: cart.VariationID}
			tx.Create(&cartProductCreate)
		} else {
			tx.Model(&cartProduct).Updates(model.CartProduct{SellerProductVariationID: cart.VariationID, Quantity: cart.Quantity})
		}
	}
	tx.Commit()
	return c.JSON(carts)
}

func CartProductOne(c *fiber.Ctx) error {
	type cartData struct {
		ProductID   uint  `json:"product_id"`
		VariationID *uint `json:"variation_id"`
		Quantity    int   `json:"quantity"`
	}
	cart := new(cartData)
	if err := c.BodyParser(cart); err != nil {
		return err
	}
	var user model.User
	if err := model.DB.Find(&user, c.Locals("user_id")); err.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var cartVal model.Cart
	tx := model.DB.Begin()
	if err := tx.Attrs(model.Cart{Slug: fmt.Sprintf("%s-%d-%d", strings.ToLower(user.Name), rand.Intn(9999), user.ID)}).FirstOrCreate(&cartVal, model.Cart{UserID: user.ID}); err.Error != nil {
		tx.Rollback()
		return c.SendStatus(fiber.StatusForbidden)
	}

	var product model.SellerProduct
	if err := model.DB.First(&product, cart.ProductID); err.Error != nil {
		tx.Rollback()
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	var productVariation model.SellerProductVariation
	if cart.VariationID != nil {
		if err := model.DB.Where("seller_product_id = ?", product.ID).First(&productVariation, "id = ?", cart.VariationID); err.Error != nil {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}
	}
	if cart.Quantity < 1 {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	var cartProduct model.CartProduct
	err := model.DB.Where("cart_id = ?", cartVal.ID).Where("seller_product_id = ?", cart.ProductID).First(&cartProduct)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		cartProductCreate := model.CartProduct{CartID: cartVal.ID, Quantity: cart.Quantity, SellerProductID: cart.ProductID, SellerProductVariationID: cart.VariationID}
		tx.Create(&cartProductCreate)
		tx.Commit()
		return c.SendStatus(201)
	} else {
		tx.Model(&cartProduct).Updates(model.CartProduct{SellerProductVariationID: cart.VariationID, Quantity: cart.Quantity})
		tx.Commit()
		return c.SendStatus(200)
	}
}
func GetCartProductAll(c *fiber.Ctx) error {
	//cart, err := models.Carts(qm.Select("id", "slug"), qm.Load(models.CartRels.CartProducts), qm.Where("user_id = ?", c.Locals("user_id"))).One(context.Background(), conn.DB)
	//fmt.Println(cart)
	//if err != nil {
	//	return err
	//}
	//
	//return c.JSON(cart)
	var cart model.Cart
	if err := model.DB.Select([]string{"id"}).Preload("CartProduct", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"cart_id", "id", "quantity", "seller_product_id", "seller_product_variation_id"})
	}).Preload("CartProduct.SellerProduct", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name", "slug", "selling_price", "product_price", "offer_price", "offer_price_start", "offer_price_end", "quantity", "next_stock", "description"})
	}).Preload("CartProduct.SellerProduct.SellerProductImage", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"image", "seller_product_id"}).Where("display = ?", true)
	}).Preload("CartProduct.SellerProductVariation", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "image", "product_price", "selling_price", "quantity", "seller_product_id"})
	}).Preload("CartProduct.SellerProductVariation.SellerProductVariationValues", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name", "description", "attribute_id", "seller_product_variation_id"})
	}).Preload("CartProduct.SellerProductVariation.SellerProductVariationValues.Attribute", func(db *gorm.DB) *gorm.DB { return db.Select([]string{"id", "name"}) }).Where("user_id = ?", c.Locals("user_id")).First(&cart); err.Error != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}
	fmt.Println(cart)
	return c.JSON(cart)
}

func CartUserRemoveProduct(c *fiber.Ctx) error {
	var cart model.Cart
	if err := model.DB.Find(&cart, "user_id = ?", c.Locals("user_id")); err.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	var cartPro []model.CartProduct
	var cartOne model.CartProduct
	res := model.DB.Where("cart_id = ?", cart.ID).Find(&cartPro)
	if res.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	var total int64
	res.Count(&total)
	fmt.Println("count: ", total)
	if err := res.Where("seller_product_id = ?", c.Params("productID")).First(&cartOne); err.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	model.DB.Unscoped().Delete(&cartOne)
	if total < 2 {
		model.DB.Unscoped().Delete(&cart)
	}
	return c.SendStatus(fiber.StatusOK)
}
func getAllLocation(c *fiber.Ctx) error {
	var locations []model.UserLocation
	if err := model.DB.Select([]string{"id", "area", "street", "house", "post_office", "post_code", "police_station", "city"}).Find(&locations, "user_id = ?", c.Locals("user_id")); err.Error != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}
	return c.JSON(locations)
}
func createLocation(c *fiber.Ctx) error {
	loc := new(model.UserLocation)
	if err := c.BodyParser(loc); err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}
	newLoc := model.UserLocation{
		UserID:        c.Locals("user_id").(uint),
		Area:          loc.Area,
		Street:        loc.Street,
		House:         loc.House,
		PostOffice:    loc.PostOffice,
		PostCode:      loc.PostCode,
		PoliceStation: loc.PoliceStation,
		City:          loc.City,
	}
	model.DB.Create(&newLoc)
	var newData model.UserLocation
	model.DB.Select([]string{"id", "area", "street", "house", "post_office", "post_code", "police_station", "city"}).First(&newData, newLoc.ID)
	return c.JSON(newData)
}
func removeLocation(c *fiber.Ctx) error {
	var location model.UserLocation
	if err := model.DB.Where("user_id = ?", c.Locals("user_id")).First(&location, c.Params("locationID")); err.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	model.DB.Unscoped().Delete(&location)
	return c.SendStatus(200)
}
func OfferPrice(productPrice decimal.Decimal, variation model.SellerProductVariation, offerPrice int64, offerStart time.Time, offerEnd time.Time) (decimal.Decimal, int) {
	price := productPrice
	today := time.Now()
	offer := 0
	if variation.SellingPrice.GreaterThan(decimal.NewFromInt(0)) {
		price = productPrice.Add(variation.SellingPrice)
	}
	if (today.Equal(offerStart) || today.After(offerStart)) && (today.Equal(offerEnd) || today.Before(offerEnd)) {
		offer = int(offerPrice)
		fmt.Println("offer price:", price.Sub(price.Mul(decimal.NewFromInt(offerPrice/100))))
		offerPricePercent := float64(offerPrice) / 100.0
		fmt.Println("Offer price percent:  ", offerPricePercent)
		offerPricePercentMul := price.Mul(decimal.NewFromFloat(offerPricePercent))
		fmt.Println("Offer price percent Multi: ", offerPricePercentMul)
		offerPriceGet := price.Sub(offerPricePercentMul)
		fmt.Println("Offer price get: ", offerPriceGet)
		return offerPriceGet, offer
	} else {
		return price, offer
	}
}
func checkoutCart(c *fiber.Ctx) error {
	var cart model.Cart
	if err := model.DB.Preload("CartProduct.SellerProduct").Preload("CartProduct.SellerProductVariation").Where("user_id = ?", c.Locals("user_id")).First(&cart); err.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	//for _, cartPro := range cart.CartProduct {
	//	return cartPro
	//}
	tx := model.DB.Begin()
	//var price []decimal.Decimal
	var location model.UserLocation
	type Obj struct {
		Location int64 `json:"location"`
	}
	var obj Obj
	if err := json.Unmarshal(c.Body(), &obj); err != nil {
		panic(err)
	}
	if err := model.DB.Where("user_id = ?", c.Locals("user_id")).First(&location, obj.Location); err.Error != nil {
		panic(err)
	}
	checkout := model.Checkout{CartID: cart.ID, UserLocationID: location.ID, UserID: c.Locals("user_id").(uint)}
	checkout.TotalPrice = decimal.NewFromInt(0)
	for _, cartPro := range cart.CartProduct {
		temp, _ := OfferPrice(cartPro.SellerProduct.SellingPrice, cartPro.SellerProductVariation, int64(cartPro.SellerProduct.OfferPrice), cartPro.SellerProduct.OfferPriceStart, cartPro.SellerProduct.OfferPriceEnd)
		checkout.TotalPrice = checkout.TotalPrice.Add(temp).Mul(decimal.NewFromInt(int64(cartPro.Quantity)))
	}
	if err := tx.Create(&checkout); err.Error != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}
	for _, cartPro := range cart.CartProduct {
		var prodQuan int
		if cartPro.SellerProductVariationID != nil {
			prodQuan = cartPro.SellerProductVariation.Quantity
		} else {
			prodQuan = cartPro.SellerProduct.Quantity
		}
		if prodQuan <= cartPro.Quantity {
			tx.Rollback()
			return c.Status(422).SendString(cartPro.SellerProduct.Name + " product quantity is less then " + strconv.FormatInt(int64(cartPro.Quantity), 10))
		}
		prodPrice, offer := OfferPrice(cartPro.SellerProduct.SellingPrice, cartPro.SellerProductVariation, int64(cartPro.SellerProduct.OfferPrice), cartPro.SellerProduct.OfferPriceStart, cartPro.SellerProduct.OfferPriceEnd)
		checkoutProduct := model.CheckoutProduct{SellerProductVariationID: cartPro.SellerProductVariationID, SellerProductID: cartPro.SellerProductID, CheckoutID: checkout.ID, Quantity: cartPro.Quantity, SellingPrice: prodPrice, OfferPrice: offer, SellingSellerID: cartPro.SellerProduct.UserID, UserID: c.Locals("user_id").(uint)}
		tx.Create(&checkoutProduct)
		if cartPro.SellerProductVariationID != nil {
			a := cartPro.SellerProductVariation.Quantity - cartPro.Quantity
			tx.Model(model.SellerProductVariation{}).Where("id = ?", cartPro.SellerProductVariation.ID).Update("quantity", a)
		} else {
			a := cartPro.SellerProduct.Quantity - cartPro.Quantity
			tx.Model(model.SellerProduct{}).Where("id = ?", cartPro.SellerProduct.ID).Update("quantity", a)
		}
	}
	tx.Delete(&cart)
	tx.Commit()
	return c.SendStatus(200)
}
