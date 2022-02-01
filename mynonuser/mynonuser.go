package mynonuser

import (
	"bongo/db"
	"bongo/ent"
	"bongo/ent/cart"
	"bongo/ent/cartproduct"
	"bongo/ent/category"
	"bongo/ent/sellerproduct"
	"bongo/ent/sellerproductimage"
	"bongo/ent/sellerproductvariation"
	"bongo/ent/shopcategory"
	"bongo/ent/user"
	"bongo/ent/userlocation"
	"bongo/model"
	"context"
	"encoding/json"
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

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
func GetShopCategories(c *fiber.Ctx) error {
	//var ShopCategories []model.ShopCategory
	//err := model.DB.Select([]string{"id", "name"}).Find(&ShopCategories)
	//if err.Error != nil {
	//	return c.SendStatus(204)
	//}
	//fmt.Println(ShopCategories)
	ShopCategories, err := db.Client.ShopCategory.Query().Where(shopcategory.DeletedAtIsNil()).Select(shopcategory.FieldID, shopcategory.FieldName).All(context.Background())

	if err != nil {
		return err
	}
	return c.JSON(ShopCategories)
}

func AllProductCategories(c *fiber.Ctx) error {
	//var categories []model.Category
	//model.DB.Select([]string{"id", "name", "slug"}).Where("parent_id IS NULL").Find(&categories)
	categories, categoriesErr := db.Client.Category.Query().Where(category.Not(category.HasParent())).Select(category.FieldID, category.FieldSlug, category.FieldName).All(context.Background())
	if categoriesErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(categories)
}
func AllProductByCategories(c *fiber.Ctx) error {

	categoryID, _ := strconv.Atoi(c.Params("categoryID"))
	dataGet, dataGetError := db.Client.Category.Query().Where(category.ID(categoryID)).WithSellerProducts(func(product *ent.SellerProductQuery) {
		product.WithSellerProductImages(func(images *ent.SellerProductImageQuery) {
			images.Where(sellerproductimage.Display(true))
		})
	}).First(context.Background())
	if dataGetError != nil {
		if ent.IsNotFound(dataGetError) {
			return c.SendStatus(fiber.StatusNoContent)
		} else {
			return c.SendStatus(500)
		}
	}
	fmt.Println(dataGet)
	return c.Status(200).JSON(dataGet)
}

func AllProducts(c *fiber.Ctx) error {
	//var products []model.SellerProduct
	//model.DB.Select([]string{"id", "name", "slug", "selling_price", "product_price", "offer_price", "offer_price_start", "offer_price_end"}).Preload("SellerProductImage", "display = (?)", true).Find(&products)
	products, productsErr := db.Client.SellerProduct.Query().Where(sellerproduct.Active(true)).Where(sellerproduct.DeletedAtIsNil()).WithSellerProductImages(func(q *ent.SellerProductImageQuery) {
		q.Where(sellerproductimage.Display(true)).Select(sellerproductimage.FieldID, sellerproductimage.FieldDisplay, sellerproductimage.FieldImage)
	}).Select(sellerproduct.FieldID, sellerproduct.FieldSlug, sellerproduct.FieldName, sellerproduct.FieldSellingPrice, sellerproduct.FieldProductPrice, sellerproduct.FieldOfferPrice, sellerproduct.FieldOfferPriceStart, sellerproduct.FieldOfferPriceEnd).All(context.Background())
	if productsErr != nil {
		return productsErr
	}

	return c.JSON(products)
}
func SingleProducts(c *fiber.Ctx) error {
	//var product model.SellerProduct
	//model.DB.Select([]string{"id", "name", "slug", "selling_price", "product_price", "offer_price", "offer_price_start", "offer_price_end", "quantity", "next_stock", "description"}).Preload("SellerProductImage", func(db *gorm.DB) *gorm.DB { return db.Select([]string{"image", "seller_product_id"}) }).Preload("SellerProductVariation", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"id", "image", "product_price", "selling_price", "quantity", "seller_product_id"})
	//}).Preload("SellerProductVariation.SellerProductVariationValues", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"id", "name", "description", "attribute_id", "seller_product_variation_id"})
	//}).Preload("SellerProductVariation.SellerProductVariationValues.Attribute", func(db *gorm.DB) *gorm.DB { return db.Select([]string{"id", "name"}) }).Find(&product, c.Params("id"))
	productID, _ := strconv.Atoi(c.Params("id"))
	product, _ := db.Client.SellerProduct.Query().Where(sellerproduct.Active(true)).Where(sellerproduct.DeletedAtIsNil()).Where(sellerproduct.ID(productID)).Select(sellerproduct.FieldID, sellerproduct.FieldSlug, sellerproduct.FieldName, sellerproduct.FieldSellingPrice, sellerproduct.FieldProductPrice, sellerproduct.FieldOfferPrice, sellerproduct.FieldOfferPriceStart, sellerproduct.FieldOfferPriceEnd).WithSellerProductImages().WithSellerProductVariations(func(variation *ent.SellerProductVariationQuery) {
		variation.WithSellerProductVariationValues(func(val *ent.SellerProductVariationValuesQuery) {
			val.WithAttribute()
		})
	}).First(context.Background())

	return c.JSON(product)
}
func GetCountCart(c *fiber.Ctx) error {

	Cart, cartErr := db.Client.Cart.Query().WithCartProducts().Where(cart.HasUserWith(user.ID(c.Locals("user_id").(int)))).Where(cart.DeletedAtIsNil()).First(context.Background())
	if cartErr != nil {
		return c.SendString("0")
	}
	if !ent.IsNotFound(cartErr) {
		product, _ := Cart.QueryCartProducts().Count(context.Background())
		count := strconv.Itoa(product)
		return c.SendString(count)
	}
	return c.SendString("0")
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
		ProductID   int `json:"product_id"`
		VariationID int `json:"variation_id"`
		Quantity    int `json:"quantity"`
	}
	var carts []cartsData
	errE := json.Unmarshal(c.Body(), &carts)
	if errE != nil {
		return errE
	}
	User, _ := db.Client.User.Get(context.Background(), c.Locals("user_id").(int))
	tx, txErr := db.Client.Tx(context.Background())
	if txErr != nil {
		return txErr
	}
	cartVal, cartValErr := db.Client.Cart.Query().Where(cart.HasUserWith(user.ID(User.ID))).Where(cart.DeletedAtIsNil()).First(context.Background())

	if ent.IsNotFound(cartValErr) {
		//name := strings.Join(strings.Split(strings.ToLower(User.Name), " "),"-")
		cartVal, _ = tx.Cart.Create().SetSlug(fmt.Sprintf("%s-%d-%d", strings.Join(strings.Split(strings.ToLower(User.Name), " "), "-"), rand.Intn(9999), User.ID)).SetUser(User).Save(context.Background())
	}

	for _, Cart := range carts {
		fmt.Println("3")
		product, _ := tx.SellerProduct.Get(context.Background(), Cart.ProductID)

		var productVariation *ent.SellerProductVariation
		if Cart.VariationID > 0 {
			var productVariationErr error
			productVariation, productVariationErr = tx.SellerProductVariation.Query().Where(sellerproductvariation.HasSellerProductWith(sellerproduct.ID(product.ID))).Where(sellerproductvariation.ID(Cart.VariationID)).First(context.Background())
			if productVariationErr != nil {
				return rollback(tx, fmt.Errorf("failed creating the group: %w", productVariationErr))
			}
		}
		if Cart.Quantity < 1 {
			return c.Status(fiber.StatusUnprocessableEntity).SendString("1 product must be order.")
			//return c.SendStatus(fiber.StatusUnprocessableEntity)
		}
		//var cartProduct model.CartProduct
		//err := model.DB.Where("cart_id = ?", cartVal.ID).Where("seller_product_id = ?", cart.ProductID).First(&cartProduct)
		_, cartProductErr := db.Client.CartProduct.Query().Where(cartproduct.HasCartWith(cart.ID(cartVal.ID))).Where(cartproduct.HasSellerProductWith(sellerproduct.ID(Cart.ProductID))).First(context.Background())

		if ent.IsNotFound(cartProductErr) {
			_, err := tx.CartProduct.Create().SetCart(cartVal).SetQuantity(Cart.Quantity).SetSellerProductID(Cart.ProductID).SetSellerProductVariation(productVariation).Save(context.Background())
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		} else {
			_, err := tx.CartProduct.Update().Where(cartproduct.HasCartWith(cart.ID(cartVal.ID))).Where(cartproduct.HasSellerProductWith(sellerproduct.ID(Cart.ProductID))).SetSellerProductVariation(productVariation).SetQuantity(Cart.Quantity).Save(context.Background())
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}
		//tx.CartProduct.Create()
	}
	tx.Commit()
	return c.SendStatus(200)
}

func CartProductOne(c *fiber.Ctx) error {
	type cartData struct {
		ProductID   int `json:"product_id"`
		VariationID int `json:"variation_id"`
		Quantity    int `json:"quantity"`
	}
	Cart := new(cartData)
	if err := c.BodyParser(Cart); err != nil {
		return err
	}
	//var User model.User
	//if err := model.DB.Find(&user, c.Locals("user_id")); err.Error != nil {
	//	return c.SendStatus(fiber.StatusBadRequest)
	//}
	User, UserErr := db.Client.User.Get(context.Background(), c.Locals("user_id").(int))
	if UserErr != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	cartVal, cartValErr := db.Client.Cart.Query().Where(cart.HasUserWith(user.ID(User.ID))).Where(cart.DeletedAtIsNil()).First(context.Background())

	if ent.IsNotFound(cartValErr) {
		//name := strings.Join(strings.Split(strings.ToLower(User.Name), " "),"-")
		cartVal, _ = db.Client.Cart.Create().SetSlug(fmt.Sprintf("%s-%d-%d", strings.Join(strings.Split(strings.ToLower(User.Name), " "), "-"), rand.Intn(9999), User.ID)).SetUser(User).Save(context.Background())
	}
	Product, ProductErr := db.Client.SellerProduct.Get(context.Background(), Cart.ProductID)
	if ProductErr != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	//var productVariation model.SellerProductVariation
	var productVariation *ent.SellerProductVariation
	if Cart.VariationID > 0 {
		var productVariationErr error
		productVariation, productVariationErr = db.Client.SellerProductVariation.Query().Where(sellerproductvariation.HasSellerProductWith(sellerproduct.ID(Product.ID))).Where(sellerproductvariation.ID(Cart.VariationID)).First(context.Background())
		if productVariationErr != nil {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}
	}
	if Cart.Quantity < 1 {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	_, cartProductErr := db.Client.CartProduct.Query().Where(cartproduct.HasCartWith(cart.ID(cartVal.ID))).Where(cartproduct.HasSellerProductWith(sellerproduct.ID(Cart.ProductID))).First(context.Background())

	if ent.IsNotFound(cartProductErr) {
		a := db.Client.CartProduct.Create().SetCart(cartVal).SetQuantity(Cart.Quantity).SetSellerProductID(Cart.ProductID)
		if productVariation != nil {
			a.SetSellerProductVariation(productVariation)
		}
		_, err := a.Save(context.Background())
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(201)
	} else {
		a := db.Client.CartProduct.Update().Where(cartproduct.HasCartWith(cart.ID(cartVal.ID))).Where(cartproduct.HasSellerProductWith(sellerproduct.ID(Cart.ProductID))).SetQuantity(Cart.Quantity)
		if productVariation != nil {
			a.SetSellerProductVariation(productVariation)
		}
		_, err := a.Save(context.Background())
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
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
	//var cart model.Cart
	//if err := model.DB.Select([]string{"id"}).Preload("CartProduct", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"cart_id", "id", "quantity", "seller_product_id", "seller_product_variation_id"})
	//}).Preload("CartProduct.SellerProduct", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"id", "name", "slug", "selling_price", "product_price", "offer_price", "offer_price_start", "offer_price_end", "quantity", "next_stock", "description"})
	//}).Preload("CartProduct.SellerProduct.SellerProductImage", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"image", "seller_product_id"}).Where("display = ?", true)
	//}).Preload("CartProduct.SellerProductVariation", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"id", "image", "product_price", "selling_price", "quantity", "seller_product_id"})
	//}).Preload("CartProduct.SellerProductVariation.SellerProductVariationValues", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"id", "name", "description", "attribute_id", "seller_product_variation_id"})
	//}).Preload("CartProduct.SellerProductVariation.SellerProductVariationValues.Attribute", func(db *gorm.DB) *gorm.DB { return db.Select([]string{"id", "name"}) }).Where("user_id = ?", c.Locals("user_id")).First(&cart); err.Error != nil {
	//	return c.SendStatus(fiber.StatusNoContent)
	//}
	Cart, CartErr := db.Client.Cart.Query().Where(cart.HasUserWith(user.ID(c.Locals("user_id").(int)))).WithCartProducts(func(cartProduct *ent.CartProductQuery) {
		cartProduct.WithSellerProduct(func(product *ent.SellerProductQuery) {
			product.WithSellerProductImages(func(image *ent.SellerProductImageQuery) {
				image.Where(sellerproductimage.Display(true))
			})
		}).WithSellerProductVariation(func(variation *ent.SellerProductVariationQuery) {
			variation.WithSellerProductVariationValues()
		})
	}).Where(cart.DeletedAtIsNil()).First(context.Background())
	if CartErr != nil {
		if ent.IsNotFound(CartErr) {
			return c.SendStatus(fiber.StatusNotFound)
		} else {
			return c.SendStatus(500)
		}
	}
	return c.JSON(Cart)
}

func CartUserRemoveProduct(c *fiber.Ctx) error {
	//var cart model.Cart
	//if err := model.DB.Find(&cart, "user_id = ?", c.Locals("user_id")); err.Error != nil {
	//	return c.SendStatus(fiber.StatusNotFound)
	//}
	Cart, CartErr := db.Client.Cart.Query().Where(cart.HasUserWith(user.ID(c.Locals("user_id").(int)))).WithCartProducts(func(q *ent.CartProductQuery) {
		q.WithSellerProduct()
	}).First(context.Background())
	if CartErr != nil {
		if ent.IsNotFound(CartErr) {
			return c.SendStatus(fiber.StatusNoContent)
		} else {
			return c.SendStatus(500)
		}
	}
	cartProductID, _ := strconv.Atoi(c.Params("cartProductID"))
	_, err := db.Client.CartProduct.Delete().Where(cartproduct.ID(cartProductID)).Where(cartproduct.HasCartWith(cart.ID(Cart.ID))).Exec(context.Background())
	if len(Cart.Edges.CartProducts) == 1 {
		_ = db.Client.Cart.DeleteOne(Cart).Exec(context.Background())
	}
	if err != nil {
		if ent.IsNotFound(err) {
			return c.SendStatus(fiber.StatusNotFound)
		} else {
			return c.SendStatus(500)
		}
	}
	return c.SendStatus(fiber.StatusOK)
}
func getAllLocation(c *fiber.Ctx) error {
	locations, _ := db.Client.UserLocation.Query().Where(userlocation.HasGetUserWith(user.ID(c.Locals("user_id").(int)))).All(context.Background())
	return c.JSON(locations)
}
func createLocation(c *fiber.Ctx) error {
	type LocInt struct {
		Area          string `json:"area"`
		Street        string `json:"street"`
		House         string `json:"house"`
		PostOffice    string `json:"post_office"`
		PostCode      int    `json:"post_code"`
		PoliceStation string `json:"police_station"`
		City          string `json:"city"`
	}
	loc := new(LocInt)
	if err := c.BodyParser(loc); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusForbidden)
	}
	location, locationErr := db.Client.UserLocation.Create().SetGetUserID(c.Locals("user_id").(int)).SetArea(loc.Area).SetStreet(loc.Street).SetHouse(loc.House).SetPostOffice(loc.PostOffice).SetPostCode(loc.PostCode).SetPoliceStation(loc.PoliceStation).SetCity(loc.City).Save(context.Background())
	if locationErr != nil {
		fmt.Println(locationErr)
		return locationErr
	}
	return c.JSON(location)
}
func removeLocation(c *fiber.Ctx) error {
	var location model.UserLocation
	if err := model.DB.Where("user_id = ?", c.Locals("user_id")).First(&location, c.Params("locationID")); err.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	model.DB.Unscoped().Delete(&location)
	locationID, _ := strconv.Atoi(c.Params("locationID"))
	_, err := db.Client.UserLocation.Delete().Where(userlocation.HasGetUserWith(user.ID(c.Locals("user_id").(int)))).Where(userlocation.ID(locationID)).Exec(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
func OfferPrice(productPrice decimal.Decimal, variation *ent.SellerProductVariation, offerPrice int64, offerStart *time.Time, offerEnd *time.Time) (decimal.Decimal, int) {

	price := productPrice
	today := time.Now()
	offer := 0
	if variation != nil && variation.SellingPrice.GreaterThan(decimal.NewFromInt(0)) {
		price = productPrice.Add(variation.SellingPrice)
	}
	if offerPrice > 0 && offerStart != nil && offerEnd != nil {
		if (today.Equal(*offerStart) || today.After(*offerStart)) && (today.Equal(*offerEnd) || today.Before(*offerEnd)) {
			fmt.Println("offer price:,", offer)
			fmt.Println("offer start:,", offerStart)
			fmt.Println("offer end:,", offerEnd)
			offer = int(offerPrice)
			fmt.Println("offer price:", price.Sub(price.Mul(decimal.NewFromInt(offerPrice/100))))
			offerPricePercent := float64(offerPrice) / 100.0
			fmt.Println("Offer price percent:  ", offerPricePercent)
			offerPricePercentMul := price.Mul(decimal.NewFromFloat(offerPricePercent))
			fmt.Println("Offer price percent Multi: ", offerPricePercentMul)
			offerPriceGet := price.Sub(offerPricePercentMul)
			fmt.Println("Offer price get: ", offerPriceGet)
			return offerPriceGet, offer

		}
	} else {
		return price, offer
	}
	return price, offer
}

func checkoutCart(c *fiber.Ctx) error {

	Cart, err := db.Client.Cart.Query().WithCartProducts(func(q *ent.CartProductQuery) {
		q.WithSellerProduct(func(product *ent.SellerProductQuery) {
			product.WithShop().WithUser()
		}).WithSellerProductVariation()
	}).Where(cart.HasUserWith(user.ID(c.Locals("user_id").(int)))).Where(cart.DeletedAtIsNil()).First(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	User, UserErr := db.Client.User.Get(context.Background(), c.Locals("user_id").(int))
	if UserErr != nil {
		return c.SendStatus(500)
	}
	fmt.Println("2")
	tx, _ := db.Client.Tx(context.Background())

	type Obj struct {
		Location int `json:"location"`
	}
	var obj Obj
	if err := json.Unmarshal(c.Body(), &obj); err != nil {
		panic(err)
	}
	fmt.Println("3")
	location, locationError := db.Client.UserLocation.Query().Where(userlocation.HasGetUserWith(user.ID(User.ID))).Where(userlocation.ID(obj.Location)).First(context.Background())
	if locationError != nil {
		return c.Status(500).SendString("Location is not correct.")
	}
	fmt.Println("4")
	fmt.Println(Cart.ID)
	myCheckout := tx.Checkout.Create().SetCart(Cart).SetLocation(location).SetUser(User)
	TotalPrice := decimal.NewFromInt(0)
	for _, cartPro := range Cart.Edges.CartProducts {
		temp, _ := OfferPrice(cartPro.Edges.SellerProduct.SellingPrice, cartPro.Edges.SellerProductVariation, int64(cartPro.Edges.SellerProduct.OfferPrice), cartPro.Edges.SellerProduct.OfferPriceStart, cartPro.Edges.SellerProduct.OfferPriceEnd)
		fmt.Println("tenp ; ", temp)
		TotalPrice = TotalPrice.Add(temp).Mul(decimal.NewFromInt(int64(cartPro.Quantity)))
	}
	fmt.Println("5")
	Checkout, CheckoutError := myCheckout.SetTotalPrice(TotalPrice).Save(context.Background())
	fmt.Println("6")
	if CheckoutError != nil {
		return rollback(tx, CheckoutError)
	}
	fmt.Println("6.6")
	for _, cartPro := range Cart.Edges.CartProducts {
		var prodQuan int
		if cartPro.Edges.SellerProductVariation != nil {
			prodQuan = cartPro.Edges.SellerProductVariation.Quantity
		} else {
			prodQuan = cartPro.Edges.SellerProduct.Quantity
		}
		fmt.Println("quantity: ", prodQuan)
		if prodQuan <= cartPro.Quantity {
			return c.Status(422).SendString(cartPro.Edges.SellerProduct.Name + " product quantity is less then " + strconv.FormatInt(int64(cartPro.Quantity), 10))
		}
		fmt.Println("7")
		prodPrice, offer := OfferPrice(cartPro.Edges.SellerProduct.SellingPrice, cartPro.Edges.SellerProductVariation, int64(cartPro.Edges.SellerProduct.OfferPrice), cartPro.Edges.SellerProduct.OfferPriceStart, cartPro.Edges.SellerProduct.OfferPriceEnd)

		initCheckout := tx.CheckoutProduct.Create().SetSellerProduct(cartPro.Edges.SellerProduct).SetCheckout(Checkout).SetQuantity(cartPro.Quantity).SetSellingPrice(prodPrice).SetUser(User).SetSeller(cartPro.Edges.SellerProduct.Edges.User)
		fmt.Println("cartPro.Edges.SellerProductVariation: ", cartPro.Edges.SellerProductVariation)
		if cartPro.Edges.SellerProductVariation != nil {
			initCheckout.SetSellerProductVariation(cartPro.Edges.SellerProductVariation)
		}
		fmt.Println("8")

		if offer > 0 {
			initCheckout.SetOfferPrice(offer)
		}
		_, checkoutProductErr := initCheckout.Save(context.Background())
		if checkoutProductErr != nil {
			return rollback(tx, checkoutProductErr)
		}
		if cartPro.Edges.SellerProductVariation != nil {
			a := cartPro.Edges.SellerProductVariation.Quantity - cartPro.Quantity
			//tx.Model(model.SellerProductVariation{}).Where("id = ?", cartPro.SellerProductVariation.ID).Update("quantity", a)
			_, QuantityUpdate := tx.SellerProductVariation.UpdateOne(cartPro.Edges.SellerProductVariation).SetQuantity(a).Save(context.Background())
			if QuantityUpdate != nil {
				return rollback(tx, QuantityUpdate)
			}
		} else {
			a := cartPro.Edges.SellerProduct.Quantity - cartPro.Quantity
			_, QuantityUpdate := tx.SellerProduct.UpdateOne(cartPro.Edges.SellerProduct).SetQuantity(a).Save(context.Background())
			if QuantityUpdate != nil {
				return rollback(tx, QuantityUpdate)
			}
		}
	}
	_, CartDeleteErr := tx.Cart.UpdateOne(Cart).SetDeletedAt(time.Now()).Save(context.Background())
	if CartDeleteErr != nil {
		return rollback(tx, CartDeleteErr)
	}
	errCommit := tx.Commit()
	if err != nil {
		return rollback(tx, errCommit)
	}
	return c.SendStatus(200)
}
