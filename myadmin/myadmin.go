package myadmin

import (
	"bongo/model"
	"bongo/myauth"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func AdminLogout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "bongoauth"
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(-3 * time.Second)
	c.Cookie(cookie)
	return c.SendStatus(200)
}
func GetShopCategoryOnly(c *fiber.Ctx) error {
	var shopCategories []model.ShopCategory
	model.DB.Find(&shopCategories)
	return c.JSON(shopCategories)
}
//func GetShopCategory(c *fiber.Ctx) error {
//	var shopCategories []model.ShopCategory
//	var shopTrashCategories []model.ShopCategory
//	model.DB.Select([]string{"ID", "Name", "Slug", "Image", "DeletedAt"}).Find(&shopCategories)
//	model.DB.Unscoped().Not("deleted_at IS NULL").Select([]string{"ID", "Name", "Slug", "Image", "DeletedAt"}).Find(&shopTrashCategories)
//
//	return c.Status(200).JSON(fiber.Map{
//		"shop_categories":       shopCategories,
//		"shop_categories_trash": shopTrashCategories,
//	})
//}

func GetShopCategory(c *fiber.Ctx) error {
	var shopCategories []model.ShopCategory
	var shopTrashCategories []model.ShopCategory
	model.DB.Select([]string{"ID", "Name", "Slug", "Image", "DeletedAt"}).Find(&shopCategories)
	model.DB.Unscoped().Not("deleted_at IS NULL").Select([]string{"ID", "Name", "Slug", "Image", "DeletedAt"}).Find(&shopTrashCategories)

	return c.Status(200).JSON(fiber.Map{
		"shop_categories":       shopCategories,
		"shop_categories_trash": shopTrashCategories,
	})
}
func GetActiveShopCategory(c *fiber.Ctx) error {
	type APIShopCategory struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	var shopCategories []APIShopCategory
	model.DB.Model(&model.ShopCategory{}).Find(&shopCategories)
	return c.JSON(shopCategories)
}

func GetTargetShopCategory(c *fiber.Ctx) error {
	var shopCategory model.ShopCategory
	err := model.DB.First(&shopCategory, "id = ?", c.Params("id"))
	if err.Error != nil {
		return c.Status(204).SendString("Data cannot be found.")
	}
	return c.JSON(shopCategory)
}
func ShopCategorySoftDelete(c *fiber.Ctx) error {
	var shopCategories []model.ShopCategory
	model.DB.Delete(&shopCategories, c.Params("id"))
	return c.JSON(fiber.Map{"status": fiber.StatusOK})
}
func ShopCategorySoftRecoverDelete(c *fiber.Ctx) error {
	var shopCategory model.ShopCategory
	model.DB.Model(&shopCategory).Unscoped().Where("id = ?", c.Params("id")).Update("deleted_at", nil)
	return c.SendStatus(200)
}

func ShopCategoryPermanentDelete(c *fiber.Ctx) error {
	var shopCategories []model.ShopCategory
	err := model.DB.Unscoped().Delete(&shopCategories, c.Params("id"))
	if err.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Cannot delete. try again.")
	}
	return c.SendStatus(200)
}
func CreateShopCategory(c *fiber.Ctx) error {
	category := new(model.ShopCategory)
	if err := c.BodyParser(category); err != nil {
		return err
	}
	file, err := c.FormFile("image")
	value, _ := file.Header["Content-Type"]
	if !(value[0] == "image/jpeg" || value[0] == "image/png") {
		return c.Status(422).JSON(fiber.Map{
			"image": "Image must be jpeg/jpg/png format.",
		})
	}
	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	imageName := fmt.Sprintf("%s.%s", filename, fileExt)
	for {
		var count int64
		model.DB.Model(&model.ShopCategory{}).Where("image = ?", imageName).Count(&count)
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

	category.Image = imageName

	category.Slug = category.Name
	for {
		var count int64
		model.DB.Model(&model.ShopCategory{}).Where("slug = ?", category.Slug).Count(&count)
		if count > 0 {
			category.Slug = fmt.Sprintf("%s-%d", category.Slug, rand.Intn(9999))
		} else {
			break
		}
	}

	model.DB.Create(&category)

	return c.Status(fiber.StatusCreated).JSON(&category)
	//return c.SendStatus(204)
}
func UpdateShopCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	category := new(model.ShopCategory)
	err := model.DB.First(&category, id)
	if err.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if err := c.BodyParser(category); err != nil {
		return err
	}
	file, fileErr := c.FormFile("image")
	if fileErr == nil {
		if _, fileerr := os.Stat("./public/images/" + category.Image); fileerr == nil {
			os.Remove("./public/images/" + category.Image)
		}
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		imageName := fmt.Sprintf("%s.%s", filename, fileExt)
		for {
			var count int64
			model.DB.Model(&model.ShopCategory{}).Where("image = ?", imageName).Count(&count)
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
		category.Image = imageName
	}

	model.DB.Save(&category)

	return c.JSON(category)
}

func GetCategory(c *fiber.Ctx) error {
	var categories []model.Category
	model.DB.Preload("Parent").Preload("ShopCategory").Find(&categories)
	return c.JSON(categories)
}
func GetDeletedCategory(c *fiber.Ctx) error {
	var categories []model.Category
	model.DB.Unscoped().Preload("Parent").Preload("ShopCategory").Not("deleted_at IS NULL").Find(&categories)
	return c.JSON(categories)
}
func CreateCategory(c *fiber.Ctx) error {
	category := new(model.Category)
	if err := c.BodyParser(category); err != nil {
		return err
	}
	category.Slug = category.Name
	for {
		var count int64
		model.DB.Model(&model.ShopCategory{}).Where("slug = ?", category.Slug).Count(&count)
		if count > 0 {
			category.Slug = fmt.Sprintf("%s-%d", category.Slug, rand.Intn(9999))
		} else {
			break
		}
	}
	model.DB.Create(&category)
	model.DB.Preload("Parent").Preload("ShopCategory").First(&category, "id = ?", category.ID)
	return c.Status(201).JSON(category)
}
func CategorySoftDelete(c *fiber.Ctx) error {
	var categories []model.Category
	model.DB.Delete(&categories, c.Params("id"))
	return c.JSON(fiber.Map{"status": fiber.StatusOK})
}
func CategoryRecoverDelete(c *fiber.Ctx) error {
	var category model.Category
	model.DB.Model(&category).Unscoped().Where("id = ?", c.Params("id")).Update("deleted_at", nil)
	return c.SendStatus(200)
}
func CategoryPermanentDelete(c *fiber.Ctx) error {
	var categories []model.Category
	model.DB.Unscoped().Delete(&categories, c.Params("id"))
	return c.JSON(fiber.Map{"status": fiber.StatusOK})
}
func UpdateCategory(c *fiber.Ctx) error {
	category := new(model.Category)
	err := model.DB.First(&category, c.Params("id"))
	if err.Error != nil {
		return c.JSON(fiber.Map{"status": fiber.StatusNotFound})
	}
	if err := c.BodyParser(category); err != nil {
		return err
	}

	model.DB.Save(&category)
	model.DB.Preload("Parent").Preload("ShopCategory").First(&category, c.Params("id"))
	return c.JSON(&category)
}

// attributes start

func GetAttributes(c *fiber.Ctx) error {
	var attributes []model.Attribute
	model.DB.Find(&attributes)

	return c.Status(200).JSON(attributes)
}
func CreateAttributes(c *fiber.Ctx) error {
	attribute := new(model.Attribute)
	if err := c.BodyParser(attribute); err != nil {
		return err
	}
	model.DB.Create(&attribute)
	return c.Status(201).JSON(attribute)
}
func GetSingleAttributes(c *fiber.Ctx) error {
	var attribute model.Attribute
	model.DB.First(&attribute, "id = ?", c.Params("id"))
	return c.Status(201).JSON(attribute)
}
func EditSingleAttributes(c *fiber.Ctx) error {
	attribute := new(model.Attribute)
	err := model.DB.First(&attribute, c.Params("id"))
	if err.Error != nil {
		return c.JSON(fiber.Map{"status": fiber.StatusNotFound})
	}
	if err := c.BodyParser(attribute); err != nil {
		return err
	}

	model.DB.Save(&attribute)
	return c.JSON(&attribute)
}
func PermanentDeleteSingleAttributes(c *fiber.Ctx) error {
	var attribute model.Attribute
	model.DB.Unscoped().Delete(&attribute, c.Params("id"))
	return c.SendStatus(200)
}
func AttributesSoftDelete(c *fiber.Ctx) error {
	var attribute []model.Attribute
	model.DB.Delete(&attribute, c.Params("id"))
	return c.JSON(fiber.Map{"status": fiber.StatusOK})
}
func AttributesSoftRecoverDelete(c *fiber.Ctx) error {
	var attribute model.Attribute
	model.DB.Model(&attribute).Unscoped().Where("id = ?", c.Params("id")).Update("deleted_at", nil)
	return c.SendStatus(200)
}

// attributes end

// seller request start
func AllSellerRequest(c *fiber.Ctx) error {
	var SellerRequest []model.SellerRequest
	var SellerTrashRequest []model.SellerRequest
	model.DB.Preload("ShopCategory").Where("accepted = ?", true).Not("deleted_at IS NULL").Find(&SellerTrashRequest)
	model.DB.Unscoped().Preload("ShopCategory").Where("accepted = ?", true).Where("deleted_at IS NULL").Find(&SellerRequest)
	return c.JSON(fiber.Map{
		"completed": SellerRequest,
		"trashes":   SellerTrashRequest,
	})
}

func GetSellerRequest(c *fiber.Ctx) error {
	var SellerRequest []model.SellerRequest
	model.DB.Order("created_at desc").Preload("ShopCategory").Where("accepted = ?", false).Find(&SellerRequest)
	return c.JSON(SellerRequest)
}
func AcceptSellerRequest(c *fiber.Ctx) error {
	var SellerRequest model.SellerRequest
	model.DB.First(&SellerRequest, "id = ?", c.Params("id"))
	var UserCount int64
	model.DB.Model(model.User{}).Where("phone_number = ?", SellerRequest.ContactNumber).Count(&UserCount)
	if UserCount == 0 {
		SellerRequest.Accepted = true
		SellerRequest.UserID = c.Locals("AuthID").(uint)
		model.DB.Save(&SellerRequest)
		hasPass, _ := myauth.HashPassword("12345678")
		user := model.User{Name: SellerRequest.SellerName, PhoneNumber: SellerRequest.ContactNumber, Password: hasPass, Seller: true}
		model.DB.Select("Name", "PhoneNumber", "Password", "Seller").Create(&user)
		return c.SendStatus(201)
	} else {
		return c.Status(422).SendString("User with this phone number already available")
	}
}
func RemoveSellerRequest(c *fiber.Ctx) error {
	var SellerRequest model.SellerRequest
	model.DB.Delete(&SellerRequest, "id = ?", c.Params("id"))
	return c.SendStatus(200)
}
func RemovePermanentSellerRequest(c *fiber.Ctx) error {
	var SellerRequest model.SellerRequest
	model.DB.Unscoped().Delete(&SellerRequest, "id = ?", c.Params("id"))
	return c.SendStatus(200)
}
func RecoverSellerRequest(c *fiber.Ctx) error {
	var SellerRequest model.SellerRequest
	model.DB.Unscoped().First(&SellerRequest, "id = ?", c.Params("id")).Update("deleted_at", nil)
	model.DB.First(&SellerRequest, "id = ?", c.Params("id"))
	return c.JSON(SellerRequest)
}

// seller request end

// seller shops start

func SellerShopsNonActivate(c *fiber.Ctx) error {
	var sellerNonActiveShops []model.SellerShop
	model.DB.Preload("ShopCategory").Where("active = ?", false).Find(&sellerNonActiveShops)
	return c.Status(200).JSON(sellerNonActiveShops)
}
func SellerShopsActivate(c *fiber.Ctx) error {
	var sellerActiveShops []model.SellerShop
	model.DB.Preload("ShopCategory").Where("active = ?", true).Where("deleted_at IS NULL").Find(&sellerActiveShops)
	return c.Status(200).JSON(sellerActiveShops)
}
func SellerShopsDeleted(c *fiber.Ctx) error {
	var sellerActiveShops []model.SellerShop
	model.DB.Preload("ShopCategory").Unscoped().Not("deleted_at IS NULL").Find(&sellerActiveShops)
	return c.Status(200).JSON(sellerActiveShops)
}

func SellerShopsAll(c *fiber.Ctx) error {
	var sellerActiveShops []model.SellerShop
	model.DB.Preload("ShopCategory").Unscoped().Find(&sellerActiveShops)
	return c.Status(200).JSON(sellerActiveShops)
}

func ActiveSellerShops(c *fiber.Ctx) error {
	shop := new(model.SellerShop)
	ShopFound := model.DB.Where("active = ?", false).First(&shop, "id = ?", c.Params("id"))
	if ShopFound.Error != nil {
		return c.SendStatus(404)
	}
	shop.Active = true
	shop.AdminID = c.Locals("AuthID").(uint)
	ShopSave := model.DB.Save(&shop)
	if ShopSave.Error != nil {
		return c.SendStatus(204)
	}
	return c.SendStatus(200)
}
func SoftDeleteSellerShops(c *fiber.Ctx) error {
	shop := new(model.SellerShop)
	model.DB.Find(&shop, "id = ?", c.Params("id")).Updates(model.SellerShop{Active: false, AdminID: c.Locals("AuthID").(uint)})
	model.DB.Delete(&shop)
	return c.SendStatus(200)
}
func PermanentDeleteSellerShops(c *fiber.Ctx) error {
	shop := new(model.SellerShop)
	model.DB.Unscoped().Where("id = ?", c.Params("id")).Find(&shop)
	model.DB.Unscoped().Delete(&shop)
	return c.SendStatus(200)
}
func RecoverDeleteSellerShops(c *fiber.Ctx) error {
	shop := new(model.SellerShop)
	model.DB.Model(&shop).Unscoped().Where("id = ?", c.Params("id")).Updates(map[string]interface{}{"deleted_at": nil, "admin_id": c.Locals("AuthID").(uint)})
	return c.SendStatus(200)
}

// seller shops end

// product brand start

func BrandAll(c *fiber.Ctx) error {
	var brands []model.Brand
	model.DB.Find(&brands)
	return c.JSON(brands)
}

func BrandCreate(c *fiber.Ctx) error {
	brand := new(model.Brand)
	if err := c.BodyParser(brand); err != nil {
		return err
	}
	model.DB.Create(&brand)
	return c.Status(200).JSON(brand)

}

func BrandEdit(c *fiber.Ctx) error {
	brand := new(model.Brand)
	err := model.DB.First(&brand, c.Params("id"))
	if err.Error != nil {
		return c.JSON(fiber.Map{"status": fiber.StatusNotFound})
	}
	if err := c.BodyParser(brand); err != nil {
		return err
	}
	model.DB.Save(&brand)
	return c.JSON(&brand)
}
func BrandSoftDelete(c *fiber.Ctx) error {
	var brands model.Brand
	model.DB.Delete(&brands, c.Params("id"))
	return c.SendStatus(200)
}
func BrandDelete(c *fiber.Ctx) error {
	var brands model.Brand
	model.DB.Unscoped().Delete(&brands, c.Params("id"))
	return c.SendStatus(200)
}

func BrandRecoverDelete(c *fiber.Ctx) error {
	var brands []model.Brand
	model.DB.Unscoped().First(&brands, "id = ?", c.Params("id")).Update("deleted_at", nil)
	return c.SendStatus(200)
}

// product brand end
