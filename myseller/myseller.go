package myseller

import (
	"bongo/db"
	"bongo/ent"
	"bongo/ent/attribute"
	"bongo/ent/category"
	"bongo/ent/checkoutproduct"
	"bongo/ent/sellerproduct"
	"bongo/ent/sellerproductimage"
	"bongo/ent/sellerproductvariation"
	"bongo/ent/sellerproductvariationvalues"
	"bongo/ent/sellershop"
	"bongo/ent/shopcategory"
	"bongo/ent/user"
	"bongo/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"github.com/shopspring/decimal"
	"image"
	"image/jpeg"
	"math/rand"
	"os"
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
func AllSellerShops(c *fiber.Ctx) error {
	AllShops, err := db.Client.SellerShop.Query().Where(sellershop.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Select(sellershop.FieldID, sellershop.FieldName, sellershop.FieldSlug, sellershop.FieldActive, sellershop.FieldBanner, sellershop.FieldDeletedAt, sellershop.FieldBusinessLocation).All(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	//data := fmt.Sprintf("%v",AllShops)
	return c.JSON(AllShops)
}
func AllSellerActiveShops(c *fiber.Ctx) error {
	//var activeShops []model.SellerShop
	//model.DB.Where("active = ?", true).Find(&activeShops, "user_id = ?", c.Locals("AuthID"))
	activeShops, err := db.Client.SellerShop.Query().Where(sellershop.Active(true)).Where(sellershop.HasUserWith(user.ID(c.Locals("AuthID").(int)))).All(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(activeShops)
}

func AllSellerInActiveShops(c *fiber.Ctx) error {
	var nonActiveShops []model.SellerShop
	model.DB.Where("active = ?", false).Where("deleted_at IS NULL").Find(&nonActiveShops, "user_id = ?", c.Locals("AuthID"))
	return c.JSON(nonActiveShops)
}

func AllSellerDeleteShops(c *fiber.Ctx) error {
	var deletedShops []model.SellerShop
	model.DB.Unscoped().Not("deleted_at IS NULL").Find(&deletedShops, "user_id = ?", c.Locals("AuthID"))
	return c.JSON(deletedShops)
}

func SingleSellerShops(c *fiber.Ctx) error {
	//var SingleShop model.SellerShop
	//query := model.DB.Where("active = ?", true).Find(&SingleShop, "user_id = ?", c.Locals("AuthID"))
	//if query.Error != nil {
	//	return c.SendStatus(fiber.StatusNotFound)
	//}
	SingleShop, _ := db.Client.SellerShop.Query().Where(sellershop.Active(true)).Where(sellershop.HasUserWith(user.ID(c.Locals("AuthID").(int)))).First(context.Background())
	return c.Status(200).JSON(SingleShop)
}

func CheckShopAvailability(c *fiber.Ctx) error {
	//var count int64
	//model.DB.Model(model.SellerShop{}).Where("active = ?", true).Where("deleted_at IS NULL").Count(&count)
	//sellerID , _ := db.Client.User.Get(context.Background(),c.Locals("AuthID").(int))
	count, _ := db.Client.SellerShop.Query().Where(sellershop.Active(true)).Where(sellershop.DeletedAtIsNil()).Where(sellershop.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Count(context.Background())
	if count == 0 {
		return c.SendStatus(200)
	} else {
		return c.SendStatus(204)
	}
}
func BrandByShop(c *fiber.Ctx) error {
	//var SingleShop model.SellerShop
	//query := model.DB.Find(&SingleShop, "user_id = ?", c.Locals("AuthID"))
	//if query.Error != nil {
	//	return c.SendStatus(fiber.StatusNotFound)
	//}
	//var brand []model.Brand
	//model.DB.Find(&brand, "shop_category_id = ?", SingleShop.ID)
	brand, _ := db.Client.Brand.Query().All(context.Background())
	return c.JSON(brand)
}
func CategoryByShop(c *fiber.Ctx) error {
	//var SingleShop model.SellerShop
	//query := model.DB.Find(&SingleShop, "user_id = ?", c.Locals("AuthID"))
	//if query.Error != nil {
	//	return c.SendStatus(fiber.StatusNotFound)
	//}
	//var category []model.Category
	//model.DB.Find(&category, "shop_category_id = ?", SingleShop.ID)

	//return c.JSON(category)\
	fmt.Println(c.Params("id"))
	shopID, _ := strconv.Atoi(c.Params("id"))
	fmt.Println(shopID)
	getShop, getShopErr := db.Client.SellerShop.Query().Where(sellershop.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellershop.ID(shopID)).WithGetShopCategory().First(context.Background())
	if getShopErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	getCategory, catErr := db.Client.Category.Query().Where(category.HasShopCategoryWith(shopcategory.ID(getShop.Edges.GetShopCategory.ID))).All(context.Background())
	if catErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(getCategory)
}
func VariationData(c *fiber.Ctx) error {
	//var variation []model.Attribute
	//model.DB.Find(&variation)
	variation, _ := db.Client.Attribute.Query().All(context.Background())
	return c.JSON(variation)
}
func CheckShopSpecificAvailability(c *fiber.Ctx) error {
	return nil
}
func CreateShops(c *fiber.Ctx) error {
	type Shop struct {
		Name             string `json:"name" form:"name"`
		Slug             string `json:"slug" form:"slug"`
		ContactNumber    string `json:"contact_number" form:"contact_number"`
		Banner           string `json:"banner" form:"banner"`
		ShopCategoryID   int    `json:"shop_category_id" form:"shop_category_id"`
		BusinessLocation string `json:"business_location" form:"business_location"`
		TaxID            string `json:"tax_id" form:"tax_id"`
	}
	shop := new(Shop)
	if err := c.BodyParser(shop); err != nil {
		return err
	}
	//fmt.Println(reflect.Type(shop.ContactNumber) )
	//var count int64
	//var categoryCheck model.ShopCategory
	//err := model.DB.First(&categoryCheck, "id = ?", shop.ShopCategoryID)
	getCategory, getCategoryErr := db.Client.ShopCategory.Get(context.Background(), shop.ShopCategoryID)
	fmt.Println(getCategory)
	if getCategoryErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Shop category cannot be found.")
	}
	//model.DB.Model(&model.User{}).Where(model.User{PhoneNumber: shop.ContactNumber}).Not("id = ?", c.Locals("AuthID")).Count(&count)
	UserPhoneCount, _ := db.Client.User.Query().Where(user.PhoneNumber(shop.ContactNumber)).Where(user.IDNEQ(c.Locals("AuthID").(int))).Count(context.Background())
	if UserPhoneCount > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Phone number is already used by another user.")
	}
	file, fileError := c.FormFile("image")
	if fileError != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Image cannot be null.")
	}
	value, _ := file.Header["Content-Type"]
	if !(value[0] == "image/jpeg" || value[0] == "image/png") {
		return c.Status(422).SendString("Image must be jpeg/jpg/png format.")
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	imageName := fmt.Sprintf("%s.%s", filename, fileExt)
	for {
		//var count int64
		//model.DB.Model(&model.SellerShop{}).Where("banner = ?", imageName).Count(&count)
		ShopBannerUniqueCount, _ := db.Client.SellerShop.Query().Where(sellershop.Banner(imageName)).Count(context.Background())
		if ShopBannerUniqueCount > 0 {
			uniqueId = uuid.New()
			filename = strings.Replace(uniqueId.String(), "-", "", -1)
			fileExt = strings.Split(file.Filename, ".")[1]
			imageName = fmt.Sprintf("%s.%s", filename, fileExt)
		} else {
			break
		}
	}
	img, _ := file.Open()
	CusImage, _, errImg := image.Decode(img)
	if errImg != nil {
		return errImg
	}
	m := resize.Resize(945, 410, CusImage, resize.Lanczos3)
	out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
	fmt.Println("1")
	if errCreate != nil {
		return c.SendStatus(500)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
	fmt.Println("2")
	shop.Banner = imageName
	shop.Slug = strings.Join(strings.Split(shop.Name, " ")[:], "_")

	for {
		//var count int64
		//model.DB.Model(&model.SellerShop{}).Where("slug = ?", shop.Slug).Count(&count)
		shopSlugCount, _ := db.Client.SellerShop.Query().Where(sellershop.Slug(shop.Slug)).Count(context.Background())
		if shopSlugCount > 0 {
			shop.Slug = fmt.Sprintf("%s-%d", shop.Slug, rand.Intn(9999))
		} else {
			break
		}
	}
	//shop.UserID = c.Locals("AuthID").(int)
	GetUser, _ := db.Client.User.Get(context.Background(), c.Locals("AuthID").(int))
	//myErr := model.DB.Select("ID", "Name", "Slug", "ContactNumber", "Banner", "ShopCategoryID", "BusinessLocation", "TaxID", "UserID").Create(&shop)
	_, saveErr := db.Client.SellerShop.Create().SetName(shop.Name).SetSlug(shop.Slug).SetContactNumber(shop.ContactNumber).SetBanner(imageName).SetBusinessLocation(shop.BusinessLocation).SetTaxID(shop.TaxID).SetGetShopCategory(getCategory).SetGetShopCategory(getCategory).SetUser(GetUser).Save(context.Background())
	//SetShopCategoryID(string(rune(getCategory.ID)))
	if saveErr != nil {
		os.Remove("./public/images/" + imageName)
		return c.SendStatus(500)
	}

	return c.SendStatus(201)
}

func EditShops(c *fiber.Ctx) error {
	//shop := new(model.SellerShop)
	//model.DB.First(&shop, "id = ?", c.Params("id"))
	SellerShopID, _ := strconv.Atoi(c.Params("id"))
	shop, shopErr := db.Client.SellerShop.Get(context.Background(), SellerShopID)
	if shopErr != nil {
		return c.SendStatus(500)
	}

	file, BannerErr := c.FormFile("banner")
	shopName := c.FormValue("name")
	shopBusiness := c.FormValue("business_location")
	a := db.Client.SellerShop.UpdateOne(shop)
	if shopName != "" {
		a.SetName(shopName)
	}
	if shopBusiness != "" {
		a.SetBusinessLocation(shopBusiness)
	}
	if BannerErr == nil {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		imageName := fmt.Sprintf("%s.%s", filename, fileExt)
		if _, fileErr := os.Stat("./public/images/" + shop.Banner); fileErr == nil {
			os.Remove("./public/images/" + shop.Banner)
		}
		for {
			//var count int64
			//model.DB.Model(&model.SellerShop{}).Where("banner = ?", imageName).Count(&count)
			count, _ := db.Client.SellerShop.Query().Where(sellershop.Banner(imageName)).Count(context.Background())
			if count > 0 {
				uniqueId := uuid.New()
				filename := strings.Replace(uniqueId.String(), "-", "", -1)
				fileExt := strings.Split(file.Filename, ".")[1]
				imageName = fmt.Sprintf("%s.%s", filename, fileExt)
			} else {
				break
			}
		}
		img, _ := file.Open()
		CusImage, _, errImg := image.Decode(img)
		if errImg != nil {
			return errImg
		}
		m := resize.Resize(945, 410, CusImage, resize.Lanczos3)
		out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
		if errCreate != nil {
			return c.SendStatus(500)
		}
		defer out.Close()
		jpeg.Encode(out, m, nil)
		a.SetBanner(imageName)
	}
	//shop.Name = c.FormValue("name")
	//shop.BusinessLocation = c.FormValue("business_location")
	//model.DB.Save(&shop)

	saveShop, saveShopErr := a.Save(context.Background())

	if saveShopErr != nil {
		return c.SendStatus(200)
	} else {
		return c.Status(200).JSON(saveShop.Banner)
	}

}
func SoftDeleteShops(c *fiber.Ctx) error {
	//var Shop model.SellerShop
	//err := model.DB.Delete(&Shop, c.Params("id"))
	//if err.Error != nil {
	//	return c.SendStatus(404)
	//}
	sellerShopID, _ := strconv.Atoi(c.Params("id"))
	sellerShopObj, sellerShopObjErr := db.Client.SellerShop.Query().Where(sellershop.ID(sellerShopID)).Where(sellershop.HasUserWith(user.ID(c.Locals("AuthID").(int)))).First(context.Background())
	if sellerShopObjErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	_, err := db.Client.SellerShop.UpdateOne(sellerShopObj).SetActive(false).SetDeletedAt(time.Now()).Save(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
func RestoreShops(c *fiber.Ctx) error {
	sellerShopID, _ := strconv.Atoi(c.Params("id"))
	sellerShopObj, sellerShopObjErr := db.Client.SellerShop.Query().Where(sellershop.ID(sellerShopID)).Where(sellershop.HasUserWith(user.ID(c.Locals("AuthID").(int)))).First(context.Background())
	if sellerShopObjErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	_, err := db.Client.SellerShop.UpdateOne(sellerShopObj).SetActive(true).ClearDeletedAt().Save(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func DeleteShops(c *fiber.Ctx) error {
	sellerShopID, _ := strconv.Atoi(c.Params("id"))
	sellerShopObj, sellerShopObjErr := db.Client.SellerShop.Query().Where(sellershop.ID(sellerShopID)).Where(sellershop.HasUserWith(user.ID(c.Locals("AuthID").(int)))).First(context.Background())
	if sellerShopObjErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	banner := sellerShopObj.Banner
	err := db.Client.SellerShop.DeleteOne(sellerShopObj).Exec(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	if _, fileErr := os.Stat("./public/images/" + banner); fileErr == nil {
		os.Remove("./public/images/" + banner)
	}
	return c.SendStatus(200)
}

// product

func AllSellerProductsMin(c *fiber.Ctx) error {
	products, _ := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).WithSellerProductImages(func(q *ent.SellerProductImageQuery) {
		q.Where(sellerproductimage.Display(true))
	}).WithShop().Where(sellerproduct.DeletedAtIsNil()).Where(sellerproduct.Active(true)).Order(ent.Asc(sellerproduct.FieldUpdatedAt)).Select(sellerproduct.FieldID, sellerproduct.FieldSlug, sellerproduct.FieldName, sellerproduct.FieldUpdatedAt, sellerproduct.FieldQuantity, sellerproduct.FieldProductPrice, sellerproduct.FieldSellingPrice).All(context.Background())
	//return c.JSON(pg.Response(products, c.Request(), &[]model.SellerProduct{}))
	return c.Status(200).JSON(products)
}
func SingleProduct(c *fiber.Ctx) error {
	//var product model.SellerProduct
	//if err := model.DB.Set("gorm:auto_preload", true).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("id")).Preload("SellerProductVariation.SellerProductVariationValues.Attribute").Preload(clause.Associations).First(&product); err.Error != nil {
	//	return c.SendStatus(fiber.StatusNotFound)
	//}
	ProductID, _ := strconv.Atoi(c.Params("id"))
	product, _ := db.Client.SellerProduct.Query().Where(sellerproduct.ID(ProductID)).Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).WithSellerProductImages().WithShop().WithSellerProductVariations(func(q *ent.SellerProductVariationQuery) {
		q.WithSellerProductVariationValues(func(qVariation *ent.SellerProductVariationValuesQuery) {
			qVariation.WithAttribute()
		})
	}).WithSellerProductVariations(func(q *ent.SellerProductVariationQuery) {
		q.WithSellerProductVariationValues(func(query *ent.SellerProductVariationValuesQuery) {
			query.WithAttribute()
		})
	}).First(context.Background())
	return c.JSON(product)
	//return c.Status(200).JSON(products)
}

func AllSellerNonProductsMin(c *fiber.Ctx) error {
	//pg := paginate.New()
	//var products []model.SellerProduct
	////model.DB.Find(&products, "user_id = ?", c.Locals("AuthID"))
	//models := model.DB.Model(&model.SellerProduct{}).Preload("SellerProductImage", "display = (?)", true).Where("user_id = ?", c.Locals("AuthID")).Where("active = ?", false)
	//fmt.Println(models)
	//return c.JSON(pg.Response(models, products, &[]model.SellerProduct{}))
	//return c.Status(200).JSON(products)
	products, _ := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).WithSellerProductImages(func(q *ent.SellerProductImageQuery) {
		q.Where(sellerproductimage.Display(true))
	}).WithShop().Where(sellerproduct.DeletedAtIsNil()).Where(sellerproduct.Active(false)).Order(ent.Asc(sellerproduct.FieldUpdatedAt)).Select(sellerproduct.FieldID, sellerproduct.FieldSlug, sellerproduct.FieldName, sellerproduct.FieldUpdatedAt, sellerproduct.FieldQuantity, sellerproduct.FieldProductPrice, sellerproduct.FieldSellingPrice).All(context.Background())
	//return c.JSON(pg.Response(products, c.Request(), &[]model.SellerProduct{}))
	return c.Status(200).JSON(products)
}
func AllSellerDeletedProductsMin(c *fiber.Ctx) error {
	//pg := paginate.New()
	//var products []model.SellerProduct
	////model.DB.Find(&products, "user_id = ?", c.Locals("AuthID"))
	//models := model.DB.Model(&model.SellerProduct{}).Unscoped().Preload("SellerProductImage", "display = (?)", true).Where("user_id = ?", c.Locals("AuthID")).Not("Deleted_at = ?", nil).Where("active = ?", false)
	//fmt.Println(models)
	//return c.JSON(pg.Response(models, products, &[]model.SellerProduct{}))
	//return c.Status(200).JSON(products)
	products, _ := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).WithSellerProductImages(func(q *ent.SellerProductImageQuery) {
		q.Where(sellerproductimage.Display(true))
	}).WithShop().Where(sellerproduct.DeletedAtNotNil()).Where(sellerproduct.Active(false)).Order(ent.Asc(sellerproduct.FieldUpdatedAt)).Select(sellerproduct.FieldID, sellerproduct.FieldSlug, sellerproduct.FieldName, sellerproduct.FieldUpdatedAt, sellerproduct.FieldQuantity, sellerproduct.FieldProductPrice, sellerproduct.FieldSellingPrice).All(context.Background())
	//return c.JSON(pg.Response(products, c.Request(), &[]model.SellerProduct{}))
	return c.Status(200).JSON(products)
}

func AllInactiveSellerProducts(c *fiber.Ctx) error {
	products, _ := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).WithSellerProductImages(func(q *ent.SellerProductImageQuery) {
		q.Where(sellerproductimage.Display(true))
	}).WithShop().Where(sellerproduct.DeletedAtIsNil()).Where(sellerproduct.Active(false)).Order(ent.Asc(sellerproduct.FieldUpdatedAt)).Select(sellerproduct.FieldID, sellerproduct.FieldSlug, sellerproduct.FieldName, sellerproduct.FieldUpdatedAt, sellerproduct.FieldQuantity, sellerproduct.FieldProductPrice, sellerproduct.FieldSellingPrice).All(context.Background())
	//return c.JSON(pg.Response(products, c.Request(), &[]model.SellerProduct{}))
	return c.Status(200).JSON(products)
}
func SoftDeleteProduct(c *fiber.Ctx) error {
	//var product model.SellerProduct
	//if check := model.DB.Find(&product, "id = ?", c.Params("id")); check.Error != nil {
	//	return c.SendStatus(fiber.StatusNotFound)
	//}
	//model.DB.Delete(&product)
	productID, _ := strconv.Atoi(c.Params("id"))
	getProduct, getProductErr := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).First(context.Background())
	if getProductErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	_, updateErr := db.Client.SellerProduct.UpdateOne(getProduct).SetDeletedAt(time.Now()).SetActive(false).Save(context.Background())
	if updateErr != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
func DeleteProduct(c *fiber.Ctx) error {
	//var product model.SellerProduct
	//if check := model.DB.Unscoped().Find(&product, "id = ?", c.Params("id")); check.Error != nil {
	//	return c.SendStatus(fiber.StatusNotFound)
	//}
	//model.DB.Unscoped().Delete(&product)
	productID, _ := strconv.Atoi(c.Params("id"))
	getProduct, getProductErr := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).First(context.Background())
	if getProductErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	err := db.Client.SellerProduct.DeleteOne(getProduct).Exec(context.Background())
	if err != nil {
		return err
	}
	return c.SendStatus(200)
}
func RecoverProduct(c *fiber.Ctx) error {
	//var product model.SellerProduct
	//if check := model.DB.Unscoped().Find(&product, "id = ?", c.Params("id")).Update("Deleted_at = ?", nil); check.Error != nil {
	//	return c.SendStatus(fiber.StatusNotFound)
	//}
	//return c.SendStatus(200)
	productID, _ := strconv.Atoi(c.Params("id"))
	getProduct, getProductErr := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).First(context.Background())
	if getProductErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	_, err := db.Client.SellerProduct.UpdateOne(getProduct).ClearDeletedAt().SetActive(true).Save(context.Background())
	if err != nil {
		return err
	}
	return c.SendStatus(200)
}

func CreateProduct(c *fiber.Ctx) error {
	//var sellerShop model.SellerShop
	//err := model.DB.Where("user_id = ?", c.Locals("AuthID").(uint)).Where("id = ?", c.Params("shopID")).Find(&sellerShop)
	shopID, _ := strconv.Atoi(c.Params("shopID"))
	sellerShop, sellerShopErr := db.Client.SellerShop.Query().Where(sellershop.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellershop.ID(shopID)).First(context.Background())
	if sellerShopErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if form, err := c.MultipartForm(); err == nil {
		tx, txErr := db.Client.Tx(context.Background())
		if txErr != nil {
			return fmt.Errorf("starting a transaction: %w", err)
		}
		// basic start
		type basic struct {
			Name           string          `json:"name"`
			Slug           string          `json:"slug"`
			Brand          int             `json:"brand"`
			ProductPrice   decimal.Decimal `json:"product_price" sql:"type:decimal(10,2)"`
			SellingPrice   decimal.Decimal `json:"selling_price" sql:"type:decimal(10,2)"`
			OfferPrice     int             `json:"offer_price"`
			Quantity       int             `json:"quantity"`
			Description    string          `json:"description"`
			OfferDateStart *time.Time      `json:"offer_date_start"`
			OfferDateEnd   *time.Time      `json:"offer_date_end"`
			NextStockDate  *time.Time      `json:"next_stock_date"`
		}

		var formBasic basic
		basicRaw := form.Value["basic"]
		err := json.Unmarshal([]byte(basicRaw[0]), &formBasic)
		if err != nil {
			return err
		}

		//fmt.Println(formBasic.OfferDateEnd)

		formBasic.Slug = strings.Join(strings.Split(formBasic.Name, " ")[:], "-")

		for {
			//var count int64
			//model.DB.Model(&model.SellerProduct{}).Where("slug = ?", formBasic.Slug).Count(&count)
			count, _ := tx.SellerProduct.Query().Where(sellerproduct.Slug(formBasic.Slug)).Count(context.Background())
			if count > 0 {
				formBasic.Slug = fmt.Sprintf("%s-%d", formBasic.Slug, rand.Intn(9999))
			} else {
				break
			}
		}
		//product := model.SellerProduct{
		//	Name:            formBasic.Name,
		//	BrandID:         formBasic.Brand,
		//	Slug:            formBasic.Slug,
		//	ProductPrice:    formBasic.ProductPrice,
		//	SellingPrice:    formBasic.SellingPrice,
		//	Quantity:        formBasic.Quantity,
		//	Active:          true,
		//	Description:     formBasic.Description,
		//	OfferPrice:      formBasic.OfferPrice,
		//	OfferPriceStart: formBasic.OfferDateStart,
		//	OfferPriceEnd:   formBasic.OfferDateEnd,
		//	NextStock:       formBasic.NextStockDate,
		//	UserID:          c.Locals("AuthID").(uint),
		//	SellerShopID:    sellerShop.ID,
		//}
		// category creating
		type category struct {
			ID int `json:"id"`
		}
		var categories []category
		categoriesRaw := form.Value["category"]

		errE := json.Unmarshal([]byte(categoriesRaw[0]), &categories)
		if errE != nil {
			return errE
		}
		fmt.Println("asce2")
		var categoriesArray []*ent.Category
		for _, value := range categories {
			//var category model.Category
			//if err := model.DB.Where("id = ?", value.ID).First(&category).Error; err != nil {
			//	return err
			//}
			categoryQuery, categoryErr := db.Client.Category.Get(context.Background(), value.ID)
			if categoryErr != nil {
				return rollback(tx, fmt.Errorf("failed creating the group: %w", categoryErr))
			}

			categoriesArray = append(categoriesArray, categoryQuery)
		}
		//sellerProductCat := model.SellerProductCategory{CategoryID: value.ID, SellerProductID: product.ID}
		//tx.Create(&sellerProductCat)

		fmt.Println("asce1")
		//return c.JSON(categories)
		brand, brandErr := tx.Brand.Get(context.Background(), formBasic.Brand)
		a := tx.SellerProduct.Create().SetName(formBasic.Name).SetSlug(formBasic.Slug).SetProductPrice(formBasic.ProductPrice).SetSellingPrice(formBasic.SellingPrice).SetQuantity(formBasic.Quantity).SetActive(true).SetDescription(formBasic.Description).SetOfferPrice(formBasic.OfferPrice).SetNillableOfferPriceStart(formBasic.OfferDateStart).SetNillableOfferPriceEnd(formBasic.OfferDateEnd).SetShop(sellerShop).SetNillableNextStock(formBasic.NextStockDate).AddCategories(categoriesArray...).SetUserID(c.Locals("AuthID").(int))
		if !ent.IsNotFound(brandErr) {
			a.SetBrand(brand)
		}
		saveProduct, saveProductErr := a.Save(context.Background())
		if saveProductErr != nil {
			return rollback(tx, fmt.Errorf("failed creating the group: %w", saveProductErr))
		}
		fmt.Println("asce")
		// creating seller shop product
		//shopProduct := model.SellerShopProduct{SellerShopID: sellerShop.ID, SellerProductID: product.ID}

		// basic end

		// product image start

		// primary
		primaryImageFile, BannerErr := c.FormFile("primary_image")

		if BannerErr != nil {
			_ = tx.Rollback()
			//return rollback(tx, fmt.Errorf("failed creating the group: %w", BannerErr))
			return c.Status(fiber.StatusUnprocessableEntity).SendString("Primary Image must be added")
		} else {
			uniqueId := uuid.New()
			filename := strings.Replace(uniqueId.String(), "-", "", -1)
			fileExt := strings.Split(primaryImageFile.Filename, ".")[1]
			imageName := fmt.Sprintf("%s.%s", filename, fileExt)
			for {
				//var count int64
				//model.DB.Model(&model.SellerProductImage{}).Where("image = ?", imageName).Count(&count)
				count, _ := db.Client.SellerProductImage.Query().Where(sellerproductimage.Image(imageName)).Count(context.Background())
				if count > 0 {
					uniqueId := uuid.New()
					fileExt := strings.Split(primaryImageFile.Filename, ".")[1]
					filename := strings.Replace(uniqueId.String(), "-", "", -1)
					fmt.Println("File ext: 1 ", fileExt)
					imageName = fmt.Sprintf("%s.%s", filename, fileExt)
				} else {
					break
				}
			}

			img, _ := primaryImageFile.Open()

			CusImage, _, errImg := image.Decode(img)

			fmt.Println(strings.Split(primaryImageFile.Filename, ".")[1])
			if errImg != nil {
				return errImg
			}

			m := resize.Resize(2048, 2048, CusImage, resize.Lanczos3)
			fmt.Println("6")
			out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
			fmt.Println("7")
			if errCreate != nil {
				return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
			}

			defer out.Close()
			jpeg.Encode(out, m, nil)
			// creating primary image
			//productImage := model.SellerProductImage{SellerProductID: product.ID, Image: imageName, Display: true}
			//primaryImageErr := tx.Create(&productImage)
			_, saveImgErr := tx.SellerProductImage.Create().SetSellerProduct(saveProduct).SetImage(imageName).SetDisplay(true).Save(context.Background())
			fmt.Println("8")
			if saveImgErr != nil {
				return rollback(tx, fmt.Errorf("failed creating the group: %w", saveImgErr))
			}
			// optional more image
			optionalImageRaw := form.File["images"]
			fmt.Println(len(optionalImageRaw) > 0)
			if len(optionalImageRaw) > 0 {
				for _, file := range optionalImageRaw {
					fmt.Println(file.Filename)
					uniqueId := uuid.New()
					filename := strings.Replace(uniqueId.String(), "-", "", -1)
					fileExt := strings.Split(file.Filename, ".")[1]
					imageName := fmt.Sprintf("%s.%s", filename, fileExt)
					fmt.Println("File ext: ", fileExt)
					for {
						//var count int64
						//
						//model.DB.Model(&model.SellerProductImage{}).Where("image = ?", imageName).Count(&count)
						count, _ := db.Client.SellerProductImage.Query().Where(sellerproductimage.Image(imageName)).Count(context.Background())
						if count > 0 {
							uniqueId := uuid.New()
							fileExt := strings.Split(file.Filename, ".")[1]
							filename := strings.Replace(uniqueId.String(), "-", "", -1)
							fmt.Println("File ext: 1 ", fileExt)
							imageName = fmt.Sprintf("%s.%s", filename, fileExt)
						} else {
							break
						}
					}

					img, _ := file.Open()

					CusImage, _, errImg := image.Decode(img)
					fmt.Println("6")
					fmt.Println(strings.Split(file.Filename, ".")[1])
					if errImg != nil {
						return errImg
					}

					m := resize.Resize(2048, 2048, CusImage, resize.Lanczos3)
					out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
					if errCreate != nil {
						_ = tx.Rollback()
						return c.SendStatus(500)
					}
					defer out.Close()
					jpeg.Encode(out, m, nil)
					// creating primary image
					//productImage := model.SellerProductImage{SellerProductID: product.ID, Image: imageName}
					//optionalImageErr := tx.Create(&productImage)
					_, saveImgErr := tx.SellerProductImage.Create().SetSellerProduct(saveProduct).SetImage(imageName).Save(context.Background())
					if saveImgErr != nil {
						return rollback(tx, fmt.Errorf("failed creating the group: %w", saveImgErr))
					}
				}
			}

		}
		// product image end

		// variance
		varianceRaw := form.Value["variance"]
		fmt.Printf("variance: %T", varianceRaw)
		fmt.Println("10")
		// check if variance available or not
		if len(varianceRaw) > 0 && varianceRaw[0] != "" {
			type variance struct {
				Color            string          `json:"color"`
				ColorDescription string          `json:"color_description"`
				Size             string          `json:"size"`
				SizeDescription  string          `json:"size_description"`
				Style            string          `json:"style"`
				StyleDescription string          `json:"style_description"`
				ProductPrice     decimal.Decimal `json:"product_price" sql:"type:decimal(10,2)"`
				SellingPrice     decimal.Decimal `json:"selling_price" sql:"type:decimal(10,2)"`
				Quantity         int             `json:"quantity"`
			}
			var variances []variance
			errVariance := json.Unmarshal([]byte(varianceRaw[0]), &variances)
			fmt.Println("variance: ", variances)
			if errVariance != nil {
				_ = tx.Rollback()
				return c.Status(500).SendString("Try again")
			}
			// get variance images

			varianceImagesRaw := form.File["variance_images"]
			for index, file := range varianceImagesRaw {
				uniqueId := uuid.New()
				filename := strings.Replace(uniqueId.String(), "-", "", -1)
				fileExt := strings.Split(file.Filename, ".")[1]
				imageName := fmt.Sprintf("%s.%s", filename, fileExt)
				for {
					//var count int64
					//model.DB.Model(&model.SellerProductVariation{}).Where("image = ?", imageName).Count(&count)
					count, _ := db.Client.SellerProductVariation.Query().Where(sellerproductvariation.Image(imageName)).Count(context.Background())
					if count > 0 {
						uniqueId := uuid.New()
						fileExt := strings.Split(file.Filename, ".")[1]
						filename := strings.Replace(uniqueId.String(), "-", "", -1)
						fmt.Println("File ext: 1 ", fileExt)
						imageName = fmt.Sprintf("%s.%s", filename, fileExt)
					} else {
						break
					}
				}

				img, _ := file.Open()

				CusImage, _, errImg := image.Decode(img)
				fmt.Println("6")
				fmt.Println(strings.Split(file.Filename, ".")[1])
				if errImg != nil {
					_ = tx.Rollback()
					return errImg
				}

				m := resize.Resize(2048, 2048, CusImage, resize.Lanczos3)
				out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
				if errCreate != nil {
					_ = tx.Rollback()
					return c.SendStatus(500)
				}
				defer out.Close()
				jpeg.Encode(out, m, nil)
				// creating product variation with image
				//productVariation := model.SellerProductVariation{
				//	SellerProductID: product.ID, Image: imageName,
				//	Quantity:     variances[index].Quantity,
				//	ProductPrice: variances[index].ProductPrice,
				//	SellingPrice: variances[index].SellingPrice,
				//}
				//err := tx.Create(&productVariation)
				productVariation, productVariationErr := tx.SellerProductVariation.Create().SetSellerProduct(saveProduct).SetQuantity(variances[index].Quantity).SetProductPrice(variances[index].ProductPrice).SetSellingPrice(variances[index].SellingPrice).SetImage(imageName).Save(context.Background())
				if productVariationErr != nil {
					return rollback(tx, fmt.Errorf("failed creating the group: %w", productVariationErr))
				}

				if variances[index].Color != "" {
					//var attribute model.Attribute
					//err := model.DB.Where("name = ?", "color").First(&attribute)
					attributeQuery, attributeErr := db.Client.Attribute.Query().Where(attribute.Name("color")).First(context.Background())
					if attributeErr != nil {
						return rollback(tx, fmt.Errorf("failed creating the group: %w", attributeErr))
					}
					//value := model.SellerProductVariationValues{Name: variances[index].Color, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].ColorDescription}
					_, saveErr := tx.SellerProductVariationValues.Create().SetName(variances[index].Color).SetAttribute(attributeQuery).SetSellerProductVariation(productVariation).SetDescription(variances[index].ColorDescription).Save(context.Background())
					if saveErr != nil {
						return rollback(tx, fmt.Errorf("failed creating the group: %w", saveErr))
					}
				}
				if variances[index].Style != "" {
					//var attribute model.Attribute
					//model.DB.Where("name = ?", "style").First(&attribute)
					//value := model.SellerProductVariationValues{Name: variances[index].Style, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].StyleDescription}
					//tx.Create(&value)
					attributeQuery, attributeErr := db.Client.Attribute.Query().Where(attribute.Name("style")).First(context.Background())
					if attributeErr != nil {
						return rollback(tx, fmt.Errorf("failed creating the group: %w", attributeErr))
					}
					_, saveErr := tx.SellerProductVariationValues.Create().SetName(variances[index].Style).SetAttribute(attributeQuery).SetSellerProductVariation(productVariation).SetDescription(variances[index].StyleDescription).Save(context.Background())
					if saveErr != nil {
						return rollback(tx, fmt.Errorf("failed creating the group: %w", saveErr))
					}
				}
				if variances[index].Size != "" {
					//var attribute model.Attribute
					//model.DB.Where("name = ?", "size").First(&attribute)
					//value := model.SellerProductVariationValues{Name: variances[index].Size, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].SizeDescription}
					//tx.Create(&value)
					attributeQuery, attributeErr := db.Client.Attribute.Query().Where(attribute.Name("size")).First(context.Background())
					if attributeErr != nil {
						return rollback(tx, fmt.Errorf("failed creating the group: %w", attributeErr))
					}
					_, saveErr := tx.SellerProductVariationValues.Create().SetName(variances[index].Size).SetAttribute(attributeQuery).SetSellerProductVariation(productVariation).SetDescription(variances[index].SizeDescription).Save(context.Background())
					if saveErr != nil {
						return rollback(tx, fmt.Errorf("failed creating the group: %w", saveErr))
					}
				}

			}

		}
		errCommit := tx.Commit()
		if errCommit != nil {
			return rollback(tx, fmt.Errorf("failed creating the group: %w", errCommit))
		}
		return c.SendStatus(201)
	}

	return c.SendStatus(fiber.StatusUnprocessableEntity)
}

func EditProductImageDelete(c *fiber.Ctx) error {
	//var count int64
	//model.DB.Model(&model.SellerProduct{}).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).Count(&count)
	productID, _ := strconv.Atoi(c.Params("product_id"))
	count, _ := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).Count(context.Background())
	if count == 1 {
		//var ProductImage model.SellerProductImage
		//errs := model.DB.Where("seller_product_id = ?", c.Params("product_id")).Find(&ProductImage, "id = ?", c.Params("image_id"))
		imageID, _ := strconv.Atoi(c.Params("image_id"))
		ProductImage, errs := db.Client.SellerProductImage.Query().Where(sellerproductimage.HasSellerProductWith(sellerproduct.ID(productID))).Where(sellerproductimage.ID(imageID)).First(context.Background())
		if errs != nil {
			return c.Status(fiber.StatusNotFound).SendString("Something is wrong. Try again.")
		}
		if ProductImage.Display {
			return c.Status(fiber.StatusForbidden).SendString("You cannot change display image of product.")
		}
		if _, fileErr := os.Stat("./public/images/" + ProductImage.Image); fileErr == nil {
			err := os.Remove("./public/images/" + ProductImage.Image)
			if err != nil {
				return c.Status(fiber.StatusNotFound).SendString("Something is wrong. Try again.")
			}
		}
		deleteErr := db.Client.SellerProductImage.DeleteOne(ProductImage).Exec(context.Background())
		if deleteErr != nil {
			return c.Status(fiber.StatusNotFound).SendString("Something is wrong. Try again.")
		}
		//model.DB.Unscoped().Delete(&ProductImage)
		return c.SendStatus(200)
	} else {
		return c.Status(fiber.StatusNotFound).SendString("Something is wrong. Try again.")
	}
}
func EditProductImageDisplay(c *fiber.Ctx) error {
	//var count int64
	//fmt.Println("product: ", c.Params("product_id"))
	//fmt.Println("user: ", c.Locals("AuthID"))
	//model.DB.Model(&model.SellerProduct{}).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).Count(&count)
	productID, _ := strconv.Atoi(c.Params("product_id"))
	count, _ := db.Client.SellerProduct.Query().Where(sellerproduct.ID(productID)).Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Count(context.Background())
	if count == 1 {
		//model.DB.Model(&model.SellerProductImage{}).Where("seller_product_id = ?", c.Params("product_id")).Where("display = ?", true).Update("display", false)
		_, _ = db.Client.SellerProductImage.Update().Where(sellerproductimage.HasSellerProductWith(sellerproduct.ID(productID))).Where(sellerproductimage.Display(true)).SetDisplay(false).Save(context.Background())
		ImageID, _ := strconv.Atoi(c.Params("image_id"))
		//errs := model.DB.Model(&model.SellerProductImage{}).Where("seller_product_id = ?", c.Params("product_id")).Where("id = ?", c.Params("image_id")).Update("display", true)
		_, errs := db.Client.SellerProductImage.Update().Where(sellerproductimage.ID(ImageID)).SetDisplay(true).Save(context.Background())
		if errs != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	} else {
		return c.SendStatus(fiber.StatusNotFound)
	}

}
func AddProductImage(c *fiber.Ctx) error {
	//var product model.SellerProduct
	//err := model.DB.Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).First(&product)
	productID, _ := strconv.Atoi(c.Params("product_id"))
	product, productErr := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).First(context.Background())
	if productErr == nil {
		if display := c.FormValue("display"); display != "" {
			fmt.Println("1")
			file, BannerErr := c.FormFile("image")
			if BannerErr == nil {
				uniqueId := uuid.New()
				filename := strings.Replace(uniqueId.String(), "-", "", -1)
				fileExt := strings.Split(file.Filename, ".")[1]
				imageName := fmt.Sprintf("%s.%s", filename, fileExt)

				for {
					//var count int64
					//model.DB.Model(&model.SellerProductImage{}).Where("image = ?", imageName).Count(&count)
					count, _ := db.Client.SellerProductImage.Query().Where(sellerproductimage.Image(imageName)).Count(context.Background())
					if count > 0 {
						uniqueId := uuid.New()
						filename := strings.Replace(uniqueId.String(), "-", "", -1)
						fileExt := strings.Split(file.Filename, ".")[1]
						imageName = fmt.Sprintf("%s.%s", filename, fileExt)
					} else {
						break
					}
				}
				img, _ := file.Open()
				CusImage, _, errImg := image.Decode(img)
				if errImg != nil {
					return errImg
				}
				m := resize.Resize(2048, 2048, CusImage, resize.Lanczos3)
				out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
				if errCreate != nil {
					return c.Status(500).SendString("Server error")
				}
				defer out.Close()
				jpeg.Encode(out, m, nil)
				//model.DB.Model(&model.SellerProductImage{}).Where("seller_product_id = ?", product.ID).Where("display = ?", true).Update("display", false)
				_, _ = db.Client.SellerProductImage.Update().Where(sellerproductimage.HasSellerProductWith(sellerproduct.ID(product.ID))).Where(sellerproductimage.Display(true)).SetDisplay(false).Save(context.Background())
				//productImage := model.SellerProductImage{Image: imageName, SellerProductID: product.ID, Display: true}
				//model.DB.Create(&productImage)
				productImage, _ := db.Client.SellerProductImage.Create().SetImage(imageName).SetSellerProduct(product).SetDisplay(true).Save(context.Background())
				return c.JSON(productImage)
			}
		} else {
			if form, err := c.MultipartForm(); err == nil {
				files := form.File["images"]
				if len(files) > 0 {
					//type ProductImage struct {
					//	ID      int    `json:"id,omitempty"`
					//}
					//var images []ent.SellerProductImage
					var images []*ent.SellerProductImage
					for _, file := range files {
						uniqueId := uuid.New()
						filename := strings.Replace(uniqueId.String(), "-", "", -1)
						fileExt := strings.Split(file.Filename, ".")[1]
						imageName := fmt.Sprintf("%s.%s", filename, fileExt)
						fmt.Println("File ext: ", fileExt)
						for {
							//var count int64
							//
							//model.DB.Model(&model.SellerProductImage{}).Where("image = ?", imageName).Count(&count)
							count, _ := db.Client.SellerProductImage.Query().Where(sellerproductimage.Image(imageName)).Count(context.Background())
							if count > 0 {
								uniqueId := uuid.New()
								fileExt := strings.Split(file.Filename, ".")[1]
								filename := strings.Replace(uniqueId.String(), "-", "", -1)
								fmt.Println("File ext: 1 ", fileExt)
								imageName = fmt.Sprintf("%s.%s", filename, fileExt)
							} else {
								break
							}
						}

						img, _ := file.Open()

						CusImage, _, errImg := image.Decode(img)
						fmt.Println(strings.Split(file.Filename, ".")[1])
						if errImg != nil {
							return errImg
						}

						m := resize.Resize(2048, 2048, CusImage, resize.Lanczos3)
						out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
						if errCreate != nil {
							return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
						}
						defer out.Close()
						jpeg.Encode(out, m, nil)
						// creating primary image
						//productImage := model.SellerProductImage{SellerProductID: product.ID, Image: imageName}
						//optionalImageErr := model.DB.Create(&productImage)
						newImg, optionalImageErr := db.Client.SellerProductImage.Create().SetSellerProduct(product).SetImage(imageName).Save(context.Background())
						if optionalImageErr != nil {
							return c.Status(500).SendString("Something error. Try again!!")
						}
						//getNewImg, _ := db.Client.SellerProductImage.Get(context.Background(),newImg.ID)
						fmt.Println(newImg)
						images = append(images, newImg)
						// add getNewImg to images array (due)
					}
					return c.JSON(images)
				}
			}
			return c.SendStatus(fiber.StatusNotFound)
		}
	} else {
		return c.Status(400).SendString("something is wrong.")
	}
	return c.SendStatus(fiber.StatusNotFound)
}
func EditBasicProduct(c *fiber.Ctx) error {
	type form struct {
		Quantity     int             `json:"quantity"`
		ProductPrice decimal.Decimal `json:"product_price"`
		SellingPrice decimal.Decimal `json:"selling_price"`
		NextStock    time.Time       `json:"next_stock"`
	}
	product := new(form)

	if err := c.BodyParser(product); err != nil {
		return err
	}

	//if err := model.DB.Model(model.SellerProduct{}).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).Updates(model.SellerProduct{Quantity: product.Quantity, ProductPrice: product.ProductPrice, SellingPrice: product.SellingPrice, NextStock: product.NextStock}); err.Error != nil {
	//	return c.SendStatus(fiber.StatusForbidden)
	//}
	productID, _ := strconv.Atoi(c.Params("product_id"))
	_, err := db.Client.SellerProduct.Update().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).SetQuantity(product.Quantity).SetNextStock(product.NextStock).SetProductPrice(product.ProductPrice).SetSellingPrice(product.SellingPrice).Save(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.SendStatus(200)
}
func EditBasicOfferProduct(c *fiber.Ctx) error {
	type OfferPrice struct {
		OfferPrice      int       `json:"offer_price"`
		OfferPriceStart time.Time `json:"offer_price_start"`
		OfferPriceEnd   time.Time `json:"offer_price_end"`
	}
	productOffer := new(OfferPrice)
	if err := c.BodyParser(productOffer); err != nil {
		return err
	}

	//if err := model.DB.Model(model.SellerProduct{}).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).Updates(model.SellerProduct{OfferPrice: productOffer.OfferPrice, OfferPriceStart: productOffer.OfferPriceStart, OfferPriceEnd: productOffer.OfferPriceEnd}); err.Error != nil {
	//	return c.SendStatus(fiber.StatusForbidden)
	//}
	productID, _ := strconv.Atoi(c.Params("product_id"))
	_, err := db.Client.SellerProduct.Update().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).SetOfferPrice(productOffer.OfferPrice).SetOfferPriceStart(productOffer.OfferPriceStart).SetOfferPriceEnd(productOffer.OfferPriceEnd).Save(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.SendStatus(200)
}
func EditProductVariation(c *fiber.Ctx) error {
	//var product model.SellerProduct

	//if err := model.DB.Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).First(&product); err.Error != nil {
	//	return c.SendStatus(fiber.StatusBadRequest)
	//}
	productID, _ := strconv.Atoi(c.Params("product_id"))
	product, err := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).First(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	type VariationInfo struct {
		ProductPrice decimal.Decimal `form:"product_price" json:"product_price"`
		SellingPrice decimal.Decimal `form:"selling_price" json:"selling_price"`
		Quantity     int             `form:"quantity" json:"quantity"`
		Image        string          `json:"image"`
	}
	productVariation := new(VariationInfo)

	if err := c.BodyParser(productVariation); err != nil {
		return err
	}

	//var variation model.SellerProductVariation
	//if err := model.DB.Where("seller_product_id = ?", product.ID).Find(&variation, "id = ?", c.Params("variation_id")); err.Error != nil {
	//	return c.Status(fiber.StatusBadRequest).SendString("Something is error. Try again.")
	//}
	variationID, _ := strconv.Atoi(c.Params("variation_id"))
	variation, variationErr := db.Client.SellerProductVariation.Query().Where(sellerproductvariation.HasSellerProductWith(sellerproduct.ID(product.ID))).Where(sellerproductvariation.ID(variationID)).First(context.Background())
	if variationErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Something is error. Try again.")
	}
	file, BannerErr := c.FormFile("image")
	productVariation.Image = variation.Image

	if BannerErr == nil {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		imageName := fmt.Sprintf("%s.%s", filename, fileExt)
		if _, fileErr := os.Stat("./public/images/" + variation.Image); fileErr == nil {
			os.Remove("./public/images/" + variation.Image)
		}
		for {
			//var count int64
			//model.DB.Model(&model.SellerProductVariation{}).Where("image = ?", imageName).Count(&count)
			count, _ := db.Client.SellerProductVariation.Query().Where(sellerproductvariation.Image(imageName)).Count(context.Background())
			if count > 0 {
				uniqueId := uuid.New()
				filename := strings.Replace(uniqueId.String(), "-", "", -1)
				fileExt := strings.Split(file.Filename, ".")[1]
				imageName = fmt.Sprintf("%s.%s", filename, fileExt)
			} else {
				break
			}
		}
		img, _ := file.Open()
		CusImage, _, errImg := image.Decode(img)
		if errImg != nil {
			return errImg
		}
		m := resize.Resize(2048, 2048, CusImage, resize.Lanczos3)
		out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
		if errCreate != nil {
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}
		defer out.Close()
		jpeg.Encode(out, m, nil)
		productVariation.Image = imageName
	}
	//if err := model.DB.Model(&variation).Updates(model.SellerProductVariation{ProductPrice: productVariation.ProductPrice, SellingPrice: productVariation.SellingPrice, Quantity: productVariation.Quantity, Image: productVariation.Image}); err.Error != nil {
	//	os.Remove(fmt.Sprintf("./public/images/%s", productVariation.Image))
	//	return c.Status(fiber.StatusBadRequest).SendString("Try again.")
	//}
	_, errSave := db.Client.SellerProductVariation.UpdateOne(variation).SetProductPrice(productVariation.ProductPrice).SetSellingPrice(productVariation.SellingPrice).SetQuantity(productVariation.Quantity).SetImage(productVariation.Image).Save(context.Background())
	if errSave != nil {
		os.Remove(fmt.Sprintf("./public/images/%s", productVariation.Image))
		return c.Status(fiber.StatusBadRequest).SendString("Try again.")
	}
	return c.Status(200).SendString(productVariation.Image)
}
func AddNewProductVariation(c *fiber.Ctx) error {
	// variance
	//var product model.SellerProduct
	//err := model.DB.Where("user_id = ?", c.Locals("AuthID").(uint)).Where("id = ?", c.Params("product_id")).Find(&product)
	productID, _ := strconv.Atoi(c.Params("product_id"))
	product, productErr := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).First(context.Background())
	if productErr == nil {
		if form, err := c.MultipartForm(); err == nil {
			type variance struct {
				ProductPrice decimal.Decimal `json:"product_price"`
				SellingPrice decimal.Decimal `json:"selling_price"`
				Quantity     int             `json:"quantity"`
				Image        string          `json:"image"`
			}
			var allVariance []*ent.SellerProductVariation

			varianceRaw := form.Value["variance"]

			// check if variance available or not
			if len(varianceRaw) > 0 && varianceRaw[0] != "" {

				type variance struct {
					Color            string          `json:"color"`
					ColorDescription string          `json:"color_description"`
					Size             string          `json:"size"`
					SizeDescription  string          `json:"size_description"`
					Style            string          `json:"style"`
					StyleDescription string          `json:"style_description"`
					ProductPrice     decimal.Decimal `json:"product_price" sql:"type:decimal(10,2)"`
					SellingPrice     decimal.Decimal `json:"selling_price" sql:"type:decimal(10,2)"`
					Quantity         int             `json:"quantity"`
				}
				var variances []variance
				errVariance := json.Unmarshal([]byte(varianceRaw[0]), &variances)
				fmt.Println("variance: ", variances)
				if errVariance != nil {
					return c.Status(500).SendString("Try again")
				}
				// get variance images
				varianceImagesRaw := form.File["variance_images"]
				for index, file := range varianceImagesRaw {
					uniqueId := uuid.New()
					filename := strings.Replace(uniqueId.String(), "-", "", -1)
					fileExt := strings.Split(file.Filename, ".")[1]
					imageName := fmt.Sprintf("%s.%s", filename, fileExt)
					for {
						//var count int64
						//model.DB.Model(&model.SellerProductVariation{}).Where("image = ?", imageName).Count(&count)
						count, _ := db.Client.SellerProductVariation.Query().Where(sellerproductvariation.Image(imageName)).Count(context.Background())
						if count > 0 {
							uniqueId := uuid.New()
							fileExt := strings.Split(file.Filename, ".")[1]
							filename := strings.Replace(uniqueId.String(), "-", "", -1)
							imageName = fmt.Sprintf("%s.%s", filename, fileExt)
						} else {
							break
						}
					}

					img, _ := file.Open()

					CusImage, _, errImg := image.Decode(img)
					fmt.Println(strings.Split(file.Filename, ".")[1])
					if errImg != nil {
						return errImg
					}

					m := resize.Resize(2048, 2048, CusImage, resize.Lanczos3)
					out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
					if errCreate != nil {
						return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
					}
					defer out.Close()
					jpeg.Encode(out, m, nil)
					// creating product variation with image

					//productVariation := model.SellerProductVariation{
					//	SellerProductID: product.ID, Image: imageName,
					//	Quantity:     variances[index].Quantity,
					//	ProductPrice: variances[index].ProductPrice,
					//	SellingPrice: variances[index].SellingPrice,
					//}
					//err := model.DB.Create(&productVariation)
					productVariation, productVariationErr := db.Client.SellerProductVariation.Create().SetSellerProduct(product).SetQuantity(variances[index].Quantity).SetImage(imageName).SetProductPrice(variances[index].ProductPrice).SetSellingPrice(variances[index].SellingPrice).Save(context.Background())
					if productVariationErr != nil {
						return c.Status(500).SendString("Something is wrong.")
					}

					if variances[index].Color != "" {
						attributeQuery, attributeErr := db.Client.Attribute.Query().Where(attribute.Name("color")).First(context.Background())
						if attributeErr != nil {
							return c.SendStatus(500)
						}
						//value := model.SellerProductVariationValues{Name: variances[index].Color, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].ColorDescription}
						_, saveErr := db.Client.SellerProductVariationValues.Create().SetName(variances[index].Color).SetAttribute(attributeQuery).SetSellerProductVariation(productVariation).SetDescription(variances[index].ColorDescription).Save(context.Background())
						if saveErr != nil {
							return c.SendStatus(500)
						}
						//model.DB.Create(&value)

					}
					if variances[index].Style != "" {
						attributeQuery, attributeErr := db.Client.Attribute.Query().Where(attribute.Name("style")).First(context.Background())
						if attributeErr != nil {
							return c.SendStatus(500)
						}
						_, saveErr := db.Client.SellerProductVariationValues.Create().SetName(variances[index].Style).SetAttribute(attributeQuery).SetSellerProductVariation(productVariation).SetDescription(variances[index].StyleDescription).Save(context.Background())
						if saveErr != nil {
							return c.SendStatus(500)
						}
					}
					if variances[index].Size != "" {
						attributeQuery, attributeErr := db.Client.Attribute.Query().Where(attribute.Name("size")).First(context.Background())
						if attributeErr != nil {
							return c.SendStatus(500)
						}
						_, saveErr := db.Client.SellerProductVariationValues.Create().SetName(variances[index].Size).SetAttribute(attributeQuery).SetSellerProductVariation(productVariation).SetDescription(variances[index].SizeDescription).Save(context.Background())
						if saveErr != nil {
							return c.SendStatus(500)
						}
					}
					//model.DB.Preload("SellerProductVariationValues.Attribute").Find(&productVariation)
					productVariationGet, _ := db.Client.SellerProductVariation.Query().Where(sellerproductvariation.ID(productVariation.ID)).WithSellerProductVariationValues(func(query *ent.SellerProductVariationValuesQuery) {
						query.WithAttribute()
					}).First(context.Background())
					allVariance = append(allVariance, productVariationGet)
				}
				return c.JSON(allVariance)
			}
		}
	}
	return c.SendStatus(400)
}

func DeleteProductVariation(c *fiber.Ctx) error {
	//var product model.SellerProduct
	//
	//if err := model.DB.Where("user_id = ?", c.Locals("AuthID").(uint)).Where("id = ?", c.Params("product_id")).Find(&product); err.Error != nil {
	//	return c.SendStatus(fiber.StatusBadRequest)
	//}
	productID, _ := strconv.Atoi(c.Params("product_id"))
	_, productErr := db.Client.SellerProduct.Query().Where(sellerproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(sellerproduct.ID(productID)).First(context.Background())
	if productErr != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	//var variation model.SellerProductVariation
	//if err := model.DB.Preload("SellerProductVariationValues").Where("seller_product_id = ?", c.Params("product_id")).Find(&variation, "id = ?", c.Params("variation_id")); err.Error != nil {
	//	return c.SendStatus(fiber.StatusBadRequest)
	//}
	variationID, _ := strconv.Atoi(c.Params("variation_id"))
	variation, _ := db.Client.SellerProductVariation.Query().Where(sellerproductvariation.HasSellerProductWith(sellerproduct.ID(productID))).Where(sellerproductvariation.ID(variationID)).First(context.Background())

	tx, _ := db.Client.Tx(context.Background())
	variationErr := tx.SellerProductVariation.DeleteOne(variation).Exec(context.Background())
	if variationErr != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	//var VariationValues []model.SellerProductVariationValues
	//if err := tx.Where("seller_product_variation_id = ?", variation.ID).Delete(&VariationValues); err.Error != nil {
	//	tx.Rollback()
	//	return c.JSON(fiber.StatusForbidden)
	//}
	_, variationValError := tx.SellerProductVariationValues.Delete().Where(sellerproductvariationvalues.HasSellerProductVariationWith(sellerproductvariation.ID(variationID))).Exec(context.Background())
	if variationValError != nil {
		return rollback(tx, fmt.Errorf("failed creating the group: %w", variationValError))
	}
	//if _, fileErr := os.Stat("./public/images/" + variation.Image); fileErr == nil {
	//	err := os.Remove("./public/images/" + variation.Image)
	//	if err != nil {
	//		return rollback(tx, fmt.Errorf("failed creating the group: %w", err))
	//		//return c.JSON(fiber.StatusForbidden)
	//	}
	//}
	if _, fileErr := os.Stat("./public/images/" + variation.Image); fileErr == nil {
		err := os.Remove("./public/images/" + variation.Image)
		if err != nil {
			return err
		}
	}

	_ = tx.Commit()
	return c.SendStatus(200)
}

// order

func MyNewOrder(c *fiber.Ctx) error {
	//var sellerCheckoutProduct []model.CheckoutProduct
	//model.DB.Select([]string{"id", "checkout_id", "seller_product_id", "quantity", "selling_price", "received", "status"}).Preload("SellerProduct", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"id", "name"})
	//}).Preload("SellerProduct.SellerProductImage", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"id", "seller_product_id", "display", "image"}).Where("display = ?", true)
	//}).Preload("Checkout", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"id", "user_location_id"})
	//}).Preload("Checkout.UserLocation", func(db *gorm.DB) *gorm.DB {
	//	return db.Select([]string{"id", "area", "street", "house", "post_office", "post_code", "police_station", "city"})
	//}).Where("selling_seller_id = ?", c.Locals("AuthID")).Find(&sellerCheckoutProduct)
	sellerCheckoutProduct, sellerCheckoutProductErr := db.Client.CheckoutProduct.Query().Where(checkoutproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).WithSellerProduct(func(image *ent.SellerProductQuery) {
		image.WithSellerProductImages(func(query *ent.SellerProductImageQuery) {
			query.Where(sellerproductimage.Display(true))
		})
	}).WithCheckout(func(query *ent.CheckoutQuery) {
		query.WithLocation()
	}).All(context.Background())
	if sellerCheckoutProductErr != nil {
		if ent.IsNotFound(sellerCheckoutProductErr) {
			return c.SendStatus(fiber.StatusNoContent)
		} else {
			return c.SendStatus(500)
		}
	}
	return c.JSON(sellerCheckoutProduct)
}

func OrderStatistic(c *fiber.Ctx) error {
	//var productOrder []model.CheckoutProduct
	//model.DB.Where("selling_seller_id = ?", c.Locals("AuthID")).Find(&productOrder)
	productOrder, _ := db.Client.CheckoutProduct.Query().Where(checkoutproduct.HasSellerWith(user.ID(c.Locals("AuthID").(int)))).All(context.Background())
	type statisticData struct {
		TodayIncome decimal.Decimal `json:"today_income"`
		TotalIncome decimal.Decimal `json:"total_income"`
	}
	var statistic statisticData
	for _, product := range productOrder {
		//productTime := product.CreatedAt.Format("01-02-2006")
		//currentDate := time.Now().Format("01-02-2006")
		if product.CreatedAt.Format("01-02-2006") == time.Now().Format("01-02-2006") {
			statistic.TodayIncome = statistic.TodayIncome.Add(product.SellingPrice)
		}
		statistic.TotalIncome = statistic.TotalIncome.Add(product.SellingPrice)
	}
	return c.JSON(statistic)
}

func DetailsCheckoutProduct(c *fiber.Ctx) error {
	checkoutProductID, _ := strconv.Atoi(c.Params("id"))
	sellerCheckoutProduct, sellerCheckoutProductErr := db.Client.CheckoutProduct.Query().WithSellerProduct(func(image *ent.SellerProductQuery) {
		image.WithSellerProductImages(func(query *ent.SellerProductImageQuery) {
			query.Where(sellerproductimage.Display(true))
		})
	}).WithCheckout(func(query *ent.CheckoutQuery) {
		query.WithLocation()
	}).WithSellerProductVariation(func(query *ent.SellerProductVariationQuery) {
		query.WithSellerProductVariationValues(func(values *ent.SellerProductVariationValuesQuery) {
			values.WithAttribute()
		})
	}).Where(checkoutproduct.HasUserWith(user.ID(c.Locals("AuthID").(int)))).Where(checkoutproduct.ID(checkoutProductID)).First(context.Background())
	if sellerCheckoutProductErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(sellerCheckoutProduct)
}
func AllShopProducts(c *fiber.Ctx) error {
	shopID, _ := strconv.Atoi(c.Params("shopID"))
	shop, shopErr := db.Client.SellerShop.Query().Where(sellershop.ID(shopID)).WithSellerProducts(func(query *ent.SellerProductQuery) {
		query.WithSellerProductImages(func(image *ent.SellerProductImageQuery) {
			image.Where(sellerproductimage.Display(true))
		}).Where(sellerproduct.DeletedAtIsNil()).Where(sellerproduct.Active(true)).Order(ent.Asc(sellerproduct.FieldUpdatedAt))
	}).First(context.Background())
	if shopErr != nil {
		return c.SendStatus(500)
	}
	return c.JSON(shop.Edges.SellerProducts)
}
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}
func GetCheckoutEveryday(c *fiber.Ctx) error {
	data, _ := db.Client.CheckoutProduct.Query().Where(checkoutproduct.HasSellerWith(user.ID(c.Locals("AuthID").(int)))).All(context.Background())
	return c.JSON(data)
}
