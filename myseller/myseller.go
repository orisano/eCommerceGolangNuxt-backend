package myseller

import (
	"bongo/model"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/morkid/paginate"
	"github.com/nfnt/resize"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
	"image"
	"image/jpeg"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func AllSellerActiveShops(c *fiber.Ctx) error {
	var activeShops []model.SellerShop
	model.DB.Where("active = ?", true).Find(&activeShops, "user_id = ?", c.Locals("AuthID"))
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
	var SingleShop model.SellerShop
	query := model.DB.Where("active = ?", true).Find(&SingleShop, "user_id = ?", c.Locals("AuthID"))
	if query.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(200).JSON(SingleShop)
}

func CheckShopAvailability(c *fiber.Ctx) error {
	var count int64
	model.DB.Model(model.SellerShop{}).Where("active = ?", true).Where("deleted_at IS NULL").Count(&count)
	if count == 0 {
		return c.SendStatus(200)
	} else {
		return c.SendStatus(204)
	}
}
func BrandByShop(c *fiber.Ctx) error {
	var SingleShop model.SellerShop
	query := model.DB.Find(&SingleShop, "user_id = ?", c.Locals("AuthID"))
	if query.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	var brand []model.Brand
	model.DB.Find(&brand, "shop_category_id = ?", SingleShop.ID)
	return c.JSON(brand)
}
func CategoryByShop(c *fiber.Ctx) error {
	var SingleShop model.SellerShop
	query := model.DB.Find(&SingleShop, "user_id = ?", c.Locals("AuthID"))
	if query.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	var category []model.Category
	model.DB.Find(&category, "shop_category_id = ?", SingleShop.ID)
	return c.JSON(category)
}
func VariationData(c *fiber.Ctx) error {
	var variation []model.Attribute
	model.DB.Find(&variation)
	return c.JSON(variation)
}
func CheckShopSpecificAvailability(c *fiber.Ctx) error {
	return nil
}
func CreateShops(c *fiber.Ctx) error {
	shop := new(model.SellerShop)
	if err := c.BodyParser(shop); err != nil {
		return err
	}
	var count int64
	var categoryCheck model.ShopCategory
	err := model.DB.First(&categoryCheck, "id = ?", shop.ShopCategoryID)
	if err.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"shop_category": "Shop category cannot be found.",
		})
	}
	model.DB.Model(&model.User{}).Where(model.User{PhoneNumber: shop.ContactNumber}).Not("id = ?", c.Locals("AuthID")).Count(&count)
	fmt.Println("auth id : ", c.Locals("AuthID"))
	if count > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"Phone number": "Phone number is already used by another user",
		})
	}
	file, fileError := c.FormFile("image")
	if fileError != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"image": "Image cannot be null",
		})
	}
	value, _ := file.Header["Content-Type"]
	if !(value[0] == "image/jpeg" || value[0] == "image/png") {
		return c.Status(422).JSON(fiber.Map{
			"image": "Image must be jpeg/jpg/png format.",
		})
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	imageName := fmt.Sprintf("%s.%s", filename, fileExt)
	for {
		var count int64
		model.DB.Model(&model.SellerShop{}).Where("banner = ?", imageName).Count(&count)
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
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)

	shop.Banner = imageName
	shop.Slug = strings.Join(strings.Split(shop.Name, " ")[:], "_")

	for {
		var count int64
		model.DB.Model(&model.SellerShop{}).Where("slug = ?", shop.Slug).Count(&count)
		if count > 0 {
			shop.Slug = fmt.Sprintf("%s-%d", shop.Slug, rand.Intn(9999))
		} else {
			break
		}
	}
	shop.UserID = c.Locals("AuthID").(uint)
	myErr := model.DB.Select("ID", "Name", "Slug", "ContactNumber", "Banner", "ShopCategoryID", "BusinessLocation", "TaxID", "UserID").Create(&shop)
	if myErr.Error != nil {
		return c.JSON(myErr)
	}
	websocketHost := os.Getenv("WEBSOCKET_HOST")
	link := fmt.Sprintf("%s%s%d", websocketHost, "/ws/admin/abc?id=", shop.ID)
	fmt.Println(link)
	http.Get(link)
	return c.Status(201).JSON(shop)
}

func EditShops(c *fiber.Ctx) error {
	shop := new(model.SellerShop)
	model.DB.First(&shop, "id = ?", c.Params("id"))
	if err := c.BodyParser(shop); err != nil {
		return err
	}

	file, BannerErr := c.FormFile("banner")
	if BannerErr == nil {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		imageName := fmt.Sprintf("%s.%s", filename, fileExt)
		if _, fileerr := os.Stat("./public/images/" + shop.Banner); fileerr == nil {
			os.Remove("./public/images/" + shop.Banner)
		}
		for {
			var count int64
			model.DB.Model(&model.SellerShop{}).Where("banner = ?", imageName).Count(&count)
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
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}
		defer out.Close()
		jpeg.Encode(out, m, nil)

		shop.Banner = imageName

	}
	shop.Name = c.FormValue("name")
	shop.BusinessLocation = c.FormValue("business_location")
	model.DB.Save(&shop)
	if BannerErr != nil {
		return c.SendStatus(200)
	} else {
		return c.Status(200).JSON(shop.Banner)
	}

}
func SoftDeleteShops(c *fiber.Ctx) error {
	var Shop model.SellerShop
	err := model.DB.Delete(&Shop, c.Params("id"))
	if err.Error != nil {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
func RestoreShops(c *fiber.Ctx) error {
	var Shop model.SellerShop
	err := model.DB.Model(&Shop).Unscoped().Where("id = ?", c.Params("id")).Update("deleted_at", nil)
	if err.Error != nil {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}

func DeleteShops(c *fiber.Ctx) error {
	var Shop model.SellerShop
	err := model.DB.Unscoped().Delete(&Shop, c.Params("id"))
	if err.Error != nil {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}

// product

func AllSellerProductsMin(c *fiber.Ctx) error {
	pg := paginate.New()
	//var products []model.SellerProduct
	//model.DB.Find(&products, "user_id = ?", c.Locals("AuthID"))
	models := model.DB.Model(&model.SellerProduct{}).Preload("SellerProductImage", "display = (?)", true).Where("user_id = ?", c.Locals("AuthID")).Where("active = ?", true)
	fmt.Println(models)
	return c.JSON(pg.Response(models, c.Request(), &[]model.SellerProduct{}))
	//return c.Status(200).JSON(products)
}
func SingleProduct(c *fiber.Ctx) error {
	var product model.SellerProduct
	if err := model.DB.Set("gorm:auto_preload", true).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("id")).Preload("SellerProductVariation.SellerProductVariationValues.Attribute").Preload(clause.Associations).First(&product); err.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(product)
	//return c.Status(200).JSON(products)
}

func AllSellerNonProductsMin(c *fiber.Ctx) error {
	pg := paginate.New()
	var products []model.SellerProduct
	//model.DB.Find(&products, "user_id = ?", c.Locals("AuthID"))
	models := model.DB.Model(&model.SellerProduct{}).Preload("SellerProductImage", "display = (?)", true).Where("user_id = ?", c.Locals("AuthID")).Where("active = ?", false)
	fmt.Println(models)
	return c.JSON(pg.Response(models, products, &[]model.SellerProduct{}))
	//return c.Status(200).JSON(products)
}
func AllSellerDeletedProductsMin(c *fiber.Ctx) error {
	pg := paginate.New()
	var products []model.SellerProduct
	//model.DB.Find(&products, "user_id = ?", c.Locals("AuthID"))
	models := model.DB.Model(&model.SellerProduct{}).Unscoped().Preload("SellerProductImage", "display = (?)", true).Where("user_id = ?", c.Locals("AuthID")).Not("Deleted_at = ?", nil).Where("active = ?", false)
	fmt.Println(models)
	return c.JSON(pg.Response(models, products, &[]model.SellerProduct{}))
	//return c.Status(200).JSON(products)
}

func AllInactiveSellerProducts(c *fiber.Ctx) error {
	pg := paginate.New()
	var products []model.SellerProduct
	models := model.DB.Where("Active = ?", true).Find(&products, "user_id = ?", c.Locals("AuthID"))
	return c.Status(200).JSON(pg.Response(models, c.Request(), &[]model.SellerProduct{}))
}
func SoftDeleteProduct(c *fiber.Ctx) error {
	var product model.SellerProduct
	if check := model.DB.Find(&product, "id = ?", c.Params("id")); check.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	model.DB.Delete(&product)
	return c.SendStatus(200)
}
func DeleteProduct(c *fiber.Ctx) error {
	var product model.SellerProduct
	if check := model.DB.Unscoped().Find(&product, "id = ?", c.Params("id")); check.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	model.DB.Unscoped().Delete(&product)
	return c.SendStatus(200)
}
func RecoverProduct(c *fiber.Ctx) error {
	var product model.SellerProduct
	if check := model.DB.Unscoped().Find(&product, "id = ?", c.Params("id")).Update("Deleted_at = ?", nil); check.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendStatus(200)
}

func CreateProduct(c *fiber.Ctx) error {
	var sellerShop model.SellerShop
	err := model.DB.Where("user_id = ?", c.Locals("AuthID").(uint)).Where("id = ?", c.Params("shopID")).Find(&sellerShop)
	if err.Error == nil {
		if form, err := c.MultipartForm(); err == nil {
			tx := model.DB.Begin()
			// basic start
			type basic struct {
				Name           string          `json:"name"`
				Slug           string          `json:"slug"`
				Brand          *uint           `json:"brand"`
				ProductPrice   decimal.Decimal `json:"product_price" sql:"type:decimal(10,2)"`
				SellingPrice   decimal.Decimal `json:"selling_price" sql:"type:decimal(10,2)"`
				OfferPrice     int             `json:"offer_price"`
				Quantity       int             `json:"quantity"`
				Description    string          `json:"description"`
				OfferDateStart *datatypes.Date `json:"offer_date_start"`
				OfferDateEnd   *datatypes.Date `json:"offer_date_end"`
				NextStockDate  *datatypes.Date `json:"next_stock_date"`
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
				var count int64
				model.DB.Model(&model.SellerProduct{}).Where("slug = ?", formBasic.Slug).Count(&count)
				if count > 0 {
					formBasic.Slug = fmt.Sprintf("%s-%d", formBasic.Slug, rand.Intn(9999))
				} else {
					break
				}
			}
			product := model.SellerProduct{
				Name:            formBasic.Name,
				BrandID:         formBasic.Brand,
				Slug:            formBasic.Slug,
				ProductPrice:    formBasic.ProductPrice,
				SellingPrice:    formBasic.SellingPrice,
				Quantity:        formBasic.Quantity,
				Active:          true,
				Description:     formBasic.Description,
				OfferPrice:      formBasic.OfferPrice,
				OfferPriceStart: formBasic.OfferDateStart,
				OfferPriceEnd:   formBasic.OfferDateEnd,
				NextStock:       formBasic.NextStockDate,
				UserID:          c.Locals("AuthID").(uint),
				SellerShopID:    sellerShop.ID,
			}
			if err := tx.Create(&product).Error; err != nil {
				tx.Rollback()
				return err
			}
			// creating seller shop product
			shopProduct := model.SellerShopProduct{SellerShopID: sellerShop.ID, SellerProductID: product.ID}
			if err := tx.Create(&shopProduct).Error; err != nil {
				tx.Rollback()
				return err
			}
			// category creating
			type category struct {
				ID uint `json:"id"`
			}
			var categories []category
			categoriesRaw := form.Value["category"]

			errE := json.Unmarshal([]byte(categoriesRaw[0]), &categories)
			if errE != nil {
				return errE
			}
			for _, value := range categories {
				var category model.Category
				if err := model.DB.Where("id = ?", value.ID).First(&category).Error; err != nil {
					return err
				}
				sellerProductCat := model.SellerProductCategory{CategoryID: value.ID, SellerProductID: product.ID}
				tx.Create(&sellerProductCat)
			}
			//return c.JSON(categories)
			// basic end

			// product image start

			// primary
			primaryImageFile, BannerErr := c.FormFile("primary_image")

			if BannerErr != nil {
				tx.Rollback()
				return c.Status(fiber.StatusUnprocessableEntity).SendString("Primary Image must be added")
			} else {
				uniqueId := uuid.New()
				filename := strings.Replace(uniqueId.String(), "-", "", -1)
				fileExt := strings.Split(primaryImageFile.Filename, ".")[1]
				imageName := fmt.Sprintf("%s.%s", filename, fileExt)
				fmt.Println("File ext: ", fileExt)
				for {
					var count int64
					model.DB.Model(&model.SellerProductImage{}).Where("image = ?", imageName).Count(&count)

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
				productImage := model.SellerProductImage{SellerProductID: product.ID, Image: imageName, Display: true}
				primaryImageErr := tx.Create(&productImage)
				fmt.Println("8")
				if primaryImageErr.Error != nil {
					tx.Rollback()
					return c.Status(500).SendString("Try again")
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
							var count int64

							model.DB.Model(&model.SellerProductImage{}).Where("image = ?", imageName).Count(&count)

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
							return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
						}
						defer out.Close()
						jpeg.Encode(out, m, nil)
						// creating primary image
						productImage := model.SellerProductImage{SellerProductID: product.ID, Image: imageName}
						optionalImageErr := tx.Create(&productImage)
						if optionalImageErr.Error != nil {
							tx.Rollback()
							return c.Status(500).SendString("Something is wrong. Please try again with proper image.")
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
					return c.Status(500).SendString("Try again")
				}
				// get variance images

				varianceImagesRaw := form.File["variance_images"]
				for index, file := range varianceImagesRaw {
					uniqueId := uuid.New()
					filename := strings.Replace(uniqueId.String(), "-", "", -1)
					fileExt := strings.Split(file.Filename, ".")[1]
					imageName := fmt.Sprintf("%s.%s", filename, fileExt)
					fmt.Println("File ext: ", fileExt)
					for {
						var count int64
						model.DB.Model(&model.SellerProductVariation{}).Where("image = ?", imageName).Count(&count)

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
						return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
					}
					defer out.Close()
					jpeg.Encode(out, m, nil)
					// creating product variation with image
					productVariation := model.SellerProductVariation{
						SellerProductID: product.ID, Image: imageName,
						Quantity:     variances[index].Quantity,
						ProductPrice: variances[index].ProductPrice,
						SellingPrice: variances[index].SellingPrice,
					}
					err := tx.Create(&productVariation)
					if err.Error != nil {
						tx.Rollback()
						return c.Status(500).SendString("Something is wrong.")
					}

					if variances[index].Color != "" {
						var attribute model.Attribute
						err := model.DB.Where("name = ?", "color").First(&attribute)
						if err.Error != nil {
							return err.Error
						}
						value := model.SellerProductVariationValues{Name: variances[index].Color, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].ColorDescription}
						tx.Create(&value)
					}
					if variances[index].Style != "" {
						var attribute model.Attribute
						model.DB.Where("name = ?", "style").First(&attribute)
						value := model.SellerProductVariationValues{Name: variances[index].Style, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].StyleDescription}
						tx.Create(&value)
					}
					if variances[index].Size != "" {
						var attribute model.Attribute
						model.DB.Where("name = ?", "size").First(&attribute)
						value := model.SellerProductVariationValues{Name: variances[index].Size, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].SizeDescription}
						tx.Create(&value)
					}

				}

			}
			tx.Commit()
			return c.SendStatus(200)
		}

	}

	return c.SendStatus(fiber.StatusUnprocessableEntity)
}

func EditProductImageDelete(c *fiber.Ctx) error {
	var count int64
	model.DB.Model(&model.SellerProduct{}).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).Count(&count)
	if count == 1 {
		var ProductImage model.SellerProductImage
		errs := model.DB.Where("seller_product_id = ?", c.Params("product_id")).Find(&ProductImage, "id = ?", c.Params("image_id"))
		if errs.Error != nil {
			return c.Status(fiber.StatusNotFound).SendString("Something is wrong. Try again.")
		}
		if ProductImage.Display {
			return c.Status(fiber.StatusForbidden).SendString("You cannot change display image of product.")
		}
		if _, fileerr := os.Stat("./public/images/" + ProductImage.Image); fileerr == nil {
			err := os.Remove("./public/images/" + ProductImage.Image)
			if err != nil {
				return c.Status(fiber.StatusNotFound).SendString("Something is wrong. Try again.")
			}
		}
		model.DB.Unscoped().Delete(&ProductImage)
		return c.SendStatus(200)
	} else {
		return c.Status(fiber.StatusNotFound).SendString("Something is wrong. Try again.")
	}
}
func EditProductImageDisplay(c *fiber.Ctx) error {
	var count int64
	fmt.Println("product: ", c.Params("product_id"))
	fmt.Println("user: ", c.Locals("AuthID"))
	model.DB.Model(&model.SellerProduct{}).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).Count(&count)
	fmt.Println("count: ", count)
	if count == 1 {
		model.DB.Model(&model.SellerProductImage{}).Where("seller_product_id = ?", c.Params("product_id")).Where("display = ?", true).Update("display", false)
		errs := model.DB.Model(&model.SellerProductImage{}).Where("seller_product_id = ?", c.Params("product_id")).Where("id = ?", c.Params("image_id")).Update("display", true)
		if errs.Error != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendStatus(200)
	} else {
		fmt.Println("painai")
		return c.SendStatus(fiber.StatusNotFound)
	}

}
func AddProductImage(c *fiber.Ctx) error {
	var product model.SellerProduct
	err := model.DB.Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).First(&product)
	fmt.Println("Error: ", err.Error)
	if err.Error == nil {
		if display := c.FormValue("display"); display != "" {
			fmt.Println("1")
			file, BannerErr := c.FormFile("image")
			if BannerErr == nil {
				uniqueId := uuid.New()
				filename := strings.Replace(uniqueId.String(), "-", "", -1)
				fileExt := strings.Split(file.Filename, ".")[1]
				imageName := fmt.Sprintf("%s.%s", filename, fileExt)

				for {
					var count int64
					model.DB.Model(&model.SellerProductImage{}).Where("image = ?", imageName).Count(&count)
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
				model.DB.Model(&model.SellerProductImage{}).Where("seller_product_id = ?", product.ID).Where("display = ?", true).Update("display", false)
				productImage := model.SellerProductImage{Image: imageName, SellerProductID: product.ID, Display: true}
				model.DB.Create(&productImage)
				return c.JSON(productImage)
			}
		} else {
			if form, err := c.MultipartForm(); err == nil {
				files := form.File["images"]
				if len(files) > 0 {
					var images []model.SellerProductImage
					for _, file := range files {
						fmt.Println(file.Filename)
						uniqueId := uuid.New()
						filename := strings.Replace(uniqueId.String(), "-", "", -1)
						fileExt := strings.Split(file.Filename, ".")[1]
						imageName := fmt.Sprintf("%s.%s", filename, fileExt)
						fmt.Println("File ext: ", fileExt)
						for {
							var count int64

							model.DB.Model(&model.SellerProductImage{}).Where("image = ?", imageName).Count(&count)

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
						productImage := model.SellerProductImage{SellerProductID: product.ID, Image: imageName}
						optionalImageErr := model.DB.Create(&productImage)
						if optionalImageErr.Error != nil {
							return c.Status(500).SendString("Something error. Try again!!")
						}
						images = append(images, productImage)
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
		NextStock    *datatypes.Date `json:"next_stock"`
	}
	product := new(form)

	if err := c.BodyParser(product); err != nil {
		return err
	}
	fmt.Println("product: ",product)
	if err := model.DB.Model(model.SellerProduct{}).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).Updates(model.SellerProduct{Quantity: product.Quantity, ProductPrice: product.ProductPrice, SellingPrice: product.SellingPrice, NextStock: product.NextStock}); err.Error != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.SendStatus(200)
}
func EditBasicOfferProduct(c *fiber.Ctx) error {
	type OfferPrice struct {
		OfferPrice      int             `json:"offer_price"`
		OfferPriceStart *datatypes.Date `json:"offer_price_start"`
		OfferPriceEnd   *datatypes.Date `json:"offer_price_end"`
	}
	productOffer := new(OfferPrice)
	if err := c.BodyParser(productOffer); err != nil {
		return err
	}

	if err := model.DB.Model(model.SellerProduct{}).Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).Updates(model.SellerProduct{OfferPrice: productOffer.OfferPrice, OfferPriceStart: productOffer.OfferPriceStart, OfferPriceEnd: productOffer.OfferPriceEnd}); err.Error != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.SendStatus(200)
}
func EditProductVariation(c *fiber.Ctx) error {
	var product model.SellerProduct

	if err := model.DB.Where("user_id = ?", c.Locals("AuthID")).Where("id = ?", c.Params("product_id")).First(&product); err.Error != nil {
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

	var variation model.SellerProductVariation
	if err := model.DB.Where("seller_product_id = ?", product.ID).Find(&variation, "id = ?", c.Params("variation_id")); err.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Something is error. Try again.")
	}
	file, BannerErr := c.FormFile("image")
	productVariation.Image = variation.Image

	if BannerErr == nil {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		imageName := fmt.Sprintf("%s.%s", filename, fileExt)
		if _, fileerr := os.Stat("./public/images/" + variation.Image); fileerr == nil {
			os.Remove("./public/images/" + variation.Image)
		}
		for {
			var count int64
			model.DB.Model(&model.SellerProductVariation{}).Where("image = ?", imageName).Count(&count)
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
	if err := model.DB.Model(&variation).Updates(model.SellerProductVariation{ProductPrice: productVariation.ProductPrice, SellingPrice: productVariation.SellingPrice, Quantity: productVariation.Quantity, Image: productVariation.Image}); err.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Try again.")
	}
	return c.SendStatus(200)
}
func AddNewProductVariation(c *fiber.Ctx) error {
	// variance
	var product model.SellerProduct
	err := model.DB.Where("user_id = ?", c.Locals("AuthID").(uint)).Where("id = ?", c.Params("product_id")).Find(&product)
	if err.Error == nil {
		if form, err := c.MultipartForm(); err == nil {
			var allVariance []model.SellerProductVariation
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
					fmt.Println(errVariance)
					return c.Status(500).SendString("Try again")
				}
				// get variance images
				varianceImagesRaw := form.File["variance_images"]
				for index, file := range varianceImagesRaw {
					uniqueId := uuid.New()
					filename := strings.Replace(uniqueId.String(), "-", "", -1)
					fileExt := strings.Split(file.Filename, ".")[1]
					imageName := fmt.Sprintf("%s.%s", filename, fileExt)
					fmt.Println("File ext: ", fileExt)
					for {
						var count int64
						model.DB.Model(&model.SellerProductVariation{}).Where("image = ?", imageName).Count(&count)

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
						return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
					}
					defer out.Close()
					jpeg.Encode(out, m, nil)
					// creating product variation with image
					productVariation := model.SellerProductVariation{
						SellerProductID: product.ID, Image: imageName,
						Quantity:     variances[index].Quantity,
						ProductPrice: variances[index].ProductPrice,
						SellingPrice: variances[index].SellingPrice,
					}
					err := model.DB.Create(&productVariation)
					if err.Error != nil {
						return c.Status(500).SendString("Something is wrong.")
					}

					if variances[index].Color != "" {
						var attribute model.Attribute
						err := model.DB.Where("name = ?", "color").First(&attribute)
						if err.Error != nil {
							return err.Error
						}
						value := model.SellerProductVariationValues{Name: variances[index].Color, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].ColorDescription}
						model.DB.Create(&value)
					}
					if variances[index].Style != "" {
						var attribute model.Attribute
						model.DB.Where("name = ?", "style").First(&attribute)
						value := model.SellerProductVariationValues{Name: variances[index].Style, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].StyleDescription}
						model.DB.Create(&value)
					}
					if variances[index].Size != "" {
						var attribute model.Attribute
						model.DB.Where("name = ?", "size").First(&attribute)
						value := model.SellerProductVariationValues{Name: variances[index].Size, AttributeID: attribute.ID, SellerProductVariationID: productVariation.ID, Description: variances[index].SizeDescription}
						model.DB.Create(&value)
					}
					model.DB.Preload("SellerProductVariationValues.Attribute").Find(&productVariation)
					allVariance = append(allVariance, productVariation)
				}
				return c.JSON(allVariance)
			}

		}

	}
	return c.SendStatus(400)
}

func DeleteProductVariation(c *fiber.Ctx) error {
	var product model.SellerProduct

	if err := model.DB.Where("user_id = ?", c.Locals("AuthID").(uint)).Where("id = ?", c.Params("product_id")).Find(&product); err.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var variation model.SellerProductVariation
	if err := model.DB.Preload("SellerProductVariationValues").Where("seller_product_id = ?", c.Params("product_id")).Find(&variation, "id = ?", c.Params("variation_id")); err.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	tx := model.DB.Begin()
	var VariationValues []model.SellerProductVariationValues
	if err := tx.Where("seller_product_variation_id = ?", variation.ID).Delete(&VariationValues); err.Error != nil {
		tx.Rollback()
		return c.JSON(fiber.StatusForbidden)
	}
	if err := tx.Delete(&variation); err.Error != nil {
		tx.Rollback()
		return c.JSON(fiber.StatusForbidden)
	}
	if _, fileerr := os.Stat("./public/images/" + variation.Image); fileerr == nil {
		err := os.Remove("./public/images/" + variation.Image)
		if err != nil {
			tx.Rollback()
			return c.JSON(fiber.StatusForbidden)
		}
	}
	tx.Commit()
	return c.SendStatus(200)
}
