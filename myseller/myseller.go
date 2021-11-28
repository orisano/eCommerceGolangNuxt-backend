package myseller

import (
	"bongo/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"gorm.io/gorm"
	"image"
	"image/jpeg"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func AllSellerShops(c *fiber.Ctx) error {
	type ShopTemp struct {
		gorm.Model
		ID            uint   `json:"id"`
		Name          string `json:"name"`
		Slug          string `json:"slug"`
		ContactNumber string `json:"contact_number"`
		Banner        string `json:"banner"`
		Active        bool   `json:"active"`
		UserID        uint   `json:"-"`
	}
	var activeShops []ShopTemp
	var nonActiveShops []ShopTemp
	var deletedShops []ShopTemp
	model.DB.Model(&model.SellerShop{}).Where("active = ?", true).Find(&activeShops, "user_id = ?", c.Locals("AuthID"))
	model.DB.Model(&model.SellerShop{}).Where("active = ?", false).Where("deleted_at IS NULL").Find(&nonActiveShops, "user_id = ?", c.Locals("AuthID"))
	model.DB.Model(&model.SellerShop{}).Not("deleted_at IS NULL").Find(&deletedShops, "user_id = ?", c.Locals("AuthID"))

	return c.Status(200).JSON(fiber.Map{
		"active_shops":     activeShops,
		"non_active_shops": nonActiveShops,
		"deleted_shops":    deletedShops,
	})
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
	shop.Slug = shop.Name

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
	link := fmt.Sprintf("%s%s%d",websocketHost,"/ws/admin/abc?id=",shop.ID)
	fmt.Println(link)
	http.Get(link)
	return c.Status(201).JSON(shop)
}
