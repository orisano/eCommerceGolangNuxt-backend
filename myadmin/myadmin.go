package myadmin

import (
	"bongo/db"
	"bongo/ent"
	"bongo/ent/attribute"
	"bongo/ent/category"
	"bongo/ent/sellerrequest"
	"bongo/ent/sellershop"
	"bongo/ent/shopcategory"
	"bongo/ent/user"
	"bongo/model"
	"bongo/myauth"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"strconv"
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
	shopCategories, err := db.Client.ShopCategory.Query().All(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	//var shopCategories []model.ShopCategory
	//model.DB.Find(&shopCategories)
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
	shopCategories, _ := db.Client.ShopCategory.Query().Where(shopcategory.DeletedAtIsNil()).All(context.Background())
	shopTrashCategories, _ := db.Client.ShopCategory.Query().Where(shopcategory.DeletedAtNotNil()).All(context.Background())

	//var shopCategories []model.ShopCategory
	//var shopTrashCategories []model.ShopCategory
	//model.DB.Select([]string{"ID", "Name", "Slug", "Image", "DeletedAt"}).Find(&shopCategories)
	//model.DB.Unscoped().Not("deleted_at IS NULL").Select([]string{"ID", "Name", "Slug", "Image", "DeletedAt"}).Find(&shopTrashCategories)
	fmt.Println(shopCategories)
	return c.Status(200).JSON(fiber.Map{
		"shop_categories":       shopCategories,
		"shop_categories_trash": shopTrashCategories,
	})
}
func GetActiveShopCategory(c *fiber.Ctx) error {
	//type APIShopCategory struct {
	//	ID   uint   `json:"id"`
	//	Name string `json:"name"`
	//}
	//var shopCategories []APIShopCategory
	//model.DB.Model(&model.ShopCategory{}).Find(&shopCategories)
	shopCategories, _ := db.Client.ShopCategory.Query().Where(shopcategory.DeletedAtIsNil()).All(context.Background())
	return c.JSON(shopCategories)
}

func GetTargetShopCategory(c *fiber.Ctx) error {
	//var shopCategory model.ShopCategory
	//err := model.DB.First(&shopCategory, "id = ?", c.Params("id"))
	//if err.Error != nil {
	//	return c.Status(204).SendString("Data cannot be found.")
	//}
	id, _ := strconv.Atoi(c.Params("id"))
	shopCategory, _ := db.Client.ShopCategory.Query().Where(shopcategory.DeletedAtIsNil()).Where(shopcategory.ID(id)).First(context.Background())
	return c.JSON(shopCategory)
}
func ShopCategorySoftDelete(c *fiber.Ctx) error {
	//var shopCategories []model.ShopCategory
	//model.DB.Delete(&shopCategories, c.Params("id"))
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := db.Client.ShopCategory.UpdateOneID(id).SetDeletedAt(time.Now()).Save(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendStatus(200)
}
func ShopCategorySoftRecoverDelete(c *fiber.Ctx) error {
	//var shopCategory model.ShopCategory
	//model.DB.Model(&shopCategory).Unscoped().Where("id = ?", c.Params("id")).Update("deleted_at", nil)
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := db.Client.ShopCategory.UpdateOneID(id).ClearDeletedAt().Save(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendStatus(200)
}

func ShopCategoryPermanentDelete(c *fiber.Ctx) error {
	//var shopCategories []model.ShopCategory
	//err := model.DB.Unscoped().Delete(&shopCategories, c.Params("id"))
	//if err.Error != nil {
	//	return c.Status(fiber.StatusBadRequest).SendString("Cannot delete. try again.")
	//}
	id, _ := strconv.Atoi(c.Params("id"))
	err := db.Client.ShopCategory.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendStatus(200)
}
func CreateShopCategory(c *fiber.Ctx) error {
	//category := new(model.ShopCategory)
	//if err := c.BodyParser(category); err != nil {
	//	return err
	//}
	file, err := c.FormFile("image")
	value, _ := file.Header["Content-Type"]
	if !(value[0] == "image/jpeg" || value[0] == "image/png") {
		return c.Status(422).JSON(fiber.Map{
			"image": "Image must be jpeg/jpg/png format.",
		})
	}
	if err != nil {
		log.Println("image upload error --> ", err)
		return c.Status(500).SendString("Server error. Try again.")
	}
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	imageName := fmt.Sprintf("%s.%s", filename, fileExt)
	for {
		//var count int64
		count, _ := db.Client.ShopCategory.Query().Where(shopcategory.Image(imageName)).Count(context.Background())
		//model.DB.Model(&model.ShopCategory{}).Where("image = ?", imageName).Count(&count)
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
		return c.Status(500).SendString("Server error. Try again.")
	}
	defer out.Close()
	EncodeError := jpeg.Encode(out, m, nil)
	if EncodeError != nil {
		return EncodeError
	}
	type tempStruct struct {
		Name  string `json:"name"`
		Slug  string `json:"slug"`
		Image string `json:"image"`
	}
	Category := new(tempStruct)
	if err := c.BodyParser(Category); err != nil {
		return c.Status(500).SendString("Server error. Try again.")
	}
	Category.Image = imageName
	Category.Slug = Category.Name
	for {
		count, _ := db.Client.ShopCategory.Query().Where(shopcategory.Slug(Category.Slug)).Where(shopcategory.DeletedAtIsNil()).Count(context.Background())
		//var count int64
		//model.DB.Model(&model.ShopCategory{}).Where("slug = ?", category.Slug).Count(&count)
		if count > 0 {
			Category.Slug = fmt.Sprintf("%s-%d", Category.Slug, rand.Intn(9999))
		} else {
			break
		}
	}
	save, myError := db.Client.ShopCategory.Create().SetSlug(Category.Slug).SetName(Category.Name).SetImage(Category.Image).Save(context.Background())
	if myError != nil {
		os.Remove("./public/images/" + Category.Image)
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Status(fiber.StatusCreated).JSON(save)
	//return c.SendStatus(204)
}
func UpdateShopCategory(c *fiber.Ctx) error {
	//id := c.Params("id")
	id, _ := strconv.Atoi(c.Params("id"))
	//type tempStruct struct {
	//	Name  string `json:"name"`
	//	Slug  string `json:"slug"`
	//	Image string `json:"image"`
	//}
	//var category tempStruct
	category := new(ent.ShopCategory)
	//category := new(model.ShopCategory)
	category, _ = db.Client.ShopCategory.Get(context.Background(), id)
	if err := c.BodyParser(category); err != nil {
		return err
	}
	fmt.Println("data: ", category)
	return c.JSON(category)
	//err := model.DB.First(&category, id)
	//if err.Error != nil {
	//	return c.SendStatus(fiber.StatusNotFound)
	//}

	//file, fileErr := c.FormFile("image")
	//if fileErr == nil {
	//	if _, fileerr := os.Stat("./public/images/" + category.Image); fileerr == nil {
	//		os.Remove("./public/images/" + category.Image)
	//	}
	//	uniqueId := uuid.New()
	//	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	//	fileExt := strings.Split(file.Filename, ".")[1]
	//	imageName := fmt.Sprintf("%s.%s", filename, fileExt)
	//	for {
	//		var count int64
	//		model.DB.Model(&model.ShopCategory{}).Where("image = ?", imageName).Count(&count)
	//		if count > 0 {
	//			uniqueId := uuid.New()
	//			filename := strings.Replace(uniqueId.String(), "-", "", -1)
	//			fileExt := strings.Split(file.Filename, ".")[1]
	//			imageName = fmt.Sprintf("%s.%s", filename, fileExt)
	//		} else {
	//			break
	//		}
	//	}
	//	img, _ := file.Open()
	//	CusImage, _, errImg := image.Decode(img)
	//	if errImg != nil {
	//		return errImg
	//	}
	//	m := resize.Resize(945, 410, CusImage, resize.Lanczos3)
	//	out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
	//	if errCreate != nil {
	//		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	//	}
	//	defer out.Close()
	//	jpeg.Encode(out, m, nil)
	//	category.Image = imageName
	//}
	//
	//model.DB.Save(&category)
	//
	//return c.JSON(category)
}

func GetCategory(c *fiber.Ctx) error {
	//var categories []model.Category
	//model.DB.Preload("Parent").Preload("ShopCategory").Find(&categories)
	categories, _ := db.Client.Category.Query().Where(category.DeletedAtIsNil()).WithShopCategory().WithChildren().WithParent().All(context.Background())
	return c.JSON(categories)
}
func GetDeletedCategory(c *fiber.Ctx) error {
	//var categories []model.Category
	//model.DB.Unscoped().Preload("Parent").Preload("ShopCategory").Not("deleted_at IS NULL").Find(&categories)
	//return c.JSON(categories)
	categories, _ := db.Client.Category.Query().Where(category.DeletedAtNotNil()).WithShopCategory().WithChildren().WithParent().All(context.Background())
	return c.JSON(categories)
}
func CreateCategory(c *fiber.Ctx) error {
	type temp struct {
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		ShopCategoryID int    `json:"shop_category_id"`
		ParentID       int    `json:"parent_id"`
	}
	myCategory := new(temp)
	if err := c.BodyParser(myCategory); err != nil {
		return err
	}
	myCategory.Slug = myCategory.Name
	for {
		count, _ := db.Client.Category.Query().Where(category.Slug(myCategory.Slug)).Count(context.Background())
		if count > 0 {
			myCategory.Slug = fmt.Sprintf("%s-%d", category.Slug, rand.Intn(9999))
		} else {
			break
		}
	}
	shopCategory, err := db.Client.ShopCategory.Get(context.Background(), myCategory.ShopCategoryID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	if myCategory.ParentID > 0 {
		parentID, _ := db.Client.Category.Get(context.Background(), myCategory.ParentID)

		save, saveErr := db.Client.Category.Create().SetName(myCategory.Name).SetSlug(myCategory.Slug).SetParent(parentID).SetShopCategory(shopCategory).Save(context.Background())
		if saveErr != nil {
			return c.Status(fiber.StatusForbidden).SendString("Please try again.")
		}
		data, _ := db.Client.Category.Query().Where(category.ID(save.ID)).WithShopCategory().WithChildren().WithParent().Only(context.Background())

		return c.Status(201).JSON(data)
	} else {
		save, saveErr := db.Client.Category.Create().SetName(myCategory.Name).SetSlug(myCategory.Slug).SetShopCategory(shopCategory).Save(context.Background())
		if saveErr != nil {
			return c.Status(fiber.StatusForbidden).SendString("Please try again.")
		}
		data, _ := db.Client.Category.Query().Where(category.ID(save.ID)).WithShopCategory().WithChildren().WithParent().Only(context.Background())

		return c.Status(201).JSON(data)
	}

}
func CategorySoftDelete(c *fiber.Ctx) error {
	//var categories []model.Category
	//model.DB.Delete(&categories, c.Params("id"))
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := db.Client.Category.UpdateOneID(id).SetDeletedAt(time.Now()).Save(context.Background())
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	return c.SendStatus(fiber.StatusOK)
}
func CategoryRecoverDelete(c *fiber.Ctx) error {
	//var category model.Category
	//model.DB.Model(&category).Unscoped().Where("id = ?", c.Params("id")).Update("deleted_at", nil)
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := db.Client.Category.UpdateOneID(id).ClearDeletedAt().Save(context.Background())
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	return c.SendStatus(200)
}
func CategoryPermanentDelete(c *fiber.Ctx) error {
	//var categories []model.Category
	//model.DB.Unscoped().Delete(&categories, c.Params("id"))
	id, _ := strconv.Atoi(c.Params("id"))
	err := db.Client.Category.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	return c.SendStatus(200)
}
func UpdateCategory(c *fiber.Ctx) error {
	type temp struct {
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		ShopCategoryID int    `json:"shop_category_id"`
		ParentID       int    `json:"parent_id"`
	}
	myCategory := new(temp)

	if err := c.BodyParser(myCategory); err != nil {
		return err
	}
	shopCategory, err := db.Client.ShopCategory.Get(context.Background(), myCategory.ShopCategoryID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	categoryID, _ := strconv.Atoi(c.Params("id"))
	if myCategory.ParentID > 0 {
		parentID, _ := db.Client.Category.Get(context.Background(), myCategory.ParentID)
		save, saveErr := db.Client.Category.UpdateOneID(categoryID).SetName(myCategory.Name).SetParent(parentID).SetShopCategory(shopCategory).Save(context.Background())
		data, _ := db.Client.Category.Query().Where(category.ID(save.ID)).WithShopCategory().WithChildren().WithParent().Only(context.Background())
		if saveErr != nil {
			return c.Status(fiber.StatusForbidden).SendString("Please try again.")
		}
		return c.Status(200).JSON(data)
	} else {
		save, saveErr := db.Client.Category.UpdateOneID(categoryID).SetName(myCategory.Name).SetShopCategory(shopCategory).Save(context.Background())
		if saveErr != nil {
			return c.Status(fiber.StatusForbidden).SendString("Please try again.")
		}
		data, _ := db.Client.Category.Query().Where(category.ID(save.ID)).WithShopCategory().WithChildren().WithParent().Only(context.Background())

		return c.Status(200).JSON(data)
	}
	//
	//model.DB.Save(&category)
	//model.DB.Preload("Parent").Preload("ShopCategory").First(&category, c.Params("id"))
	//return c.JSON(200)
}

// attributes start

func GetAttributes(c *fiber.Ctx) error {
	//var attributes []model.Attribute
	//model.DB.Find(&attributes)
	attributes, err := db.Client.Attribute.Query().Where(attribute.DeletedAtIsNil()).All(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(attributes)
}
func CreateAttributes(c *fiber.Ctx) error {
	type temp struct {
		Name string `json:"name"`
	}
	myAttribute := new(temp)
	if err := c.BodyParser(myAttribute); err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	save, err := db.Client.Attribute.Create().SetName(myAttribute.Name).Save(context.Background())
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	//model.DB.Create(&attribute)
	return c.Status(201).JSON(save)
}
func GetSingleAttributes(c *fiber.Ctx) error {
	//var attribute model.Attribute
	//model.DB.First(&attribute, "id = ?", c.Params("id"))
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := db.Client.Attribute.Get(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	return c.Status(201).JSON(data)
}
func EditSingleAttributes(c *fiber.Ctx) error {
	type temp struct {
		Name string `json:"name"`
	}
	myAttribute := new(temp)
	//err := model.DB.First(&attribute, c.Params("id"))
	//if err.Error != nil {
	//	return c.JSON(fiber.Map{"status": fiber.StatusNotFound})
	//}

	if err := c.BodyParser(myAttribute); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	if myAttribute.Name == "" {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Attribute name cannot be empty.")
	}
	save, err := db.Client.Attribute.UpdateOneID(id).SetName(myAttribute.Name).Save(context.Background())
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	return c.JSON(&save)
	//model.DB.Save(&attribute)

}
func PermanentDeleteSingleAttributes(c *fiber.Ctx) error {
	//var attribute model.Attribute
	//model.DB.Unscoped().Delete(&attribute, c.Params("id"))
	id, _ := strconv.Atoi(c.Params("id"))
	err := db.Client.Attribute.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	return c.SendStatus(200)
}
func AttributesSoftDelete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	//var attribute []model.Attribute
	//model.DB.Delete(&attribute, c.Params("id"))
	//return c.JSON(fiber.Map{"status": fiber.StatusOK})
	_, err := db.Client.Attribute.UpdateOneID(id).SetDeletedAt(time.Now()).Save(context.Background())
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	return c.SendStatus(200)
}
func AttributesSoftRecoverDelete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	//var attribute model.Attribute
	//model.DB.Model(&attribute).Unscoped().Where("id = ?", c.Params("id")).Update("deleted_at", nil)
	//return c.SendStatus(200)
	_, err := db.Client.Attribute.UpdateOneID(id).SetNillableDeletedAt(nil).Save(context.Background())
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Please try again.")
	}
	return c.SendStatus(200)
}

// attributes end

// seller request start

func AllSellerRequest(c *fiber.Ctx) error {
	//var SellerRequest []model.SellerRequest
	//var SellerTrashRequest []model.SellerRequest
	//model.DB.Preload("ShopCategory").Where("accepted = ?", true).Not("deleted_at IS NULL").Find(&SellerTrashRequest)
	//model.DB.Unscoped().Preload("ShopCategory").Where("accepted = ?", true).Where("deleted_at IS NULL").Find(&SellerRequest)
	SellerRequest, _ := db.Client.SellerRequest.Query().Where(sellerrequest.Accepted(true)).Where(sellerrequest.DeletedAtIsNil()).WithShopCategory().All(context.Background())
	SellerTrashRequest, _ := db.Client.SellerRequest.Query().Where(sellerrequest.Accepted(true)).Where(sellerrequest.DeletedAtNotNil()).WithShopCategory().All(context.Background())
	return c.JSON(fiber.Map{
		"completed": SellerRequest,
		"trashes":   SellerTrashRequest,
	})

}

func GetSellerRequest(c *fiber.Ctx) error {
	//var SellerRequest []model.SellerRequest
	//model.DB.Order("created_at desc").Preload("ShopCategory").Where("accepted = ?", false).Find(&SellerRequest)
	SellerRequest, err := db.Client.SellerRequest.Query().Where(sellerrequest.Accepted(false)).Where(sellerrequest.DeletedAtIsNil()).WithShopCategory().All(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}
	return c.Status(200).JSON(SellerRequest)
}
func AcceptSellerRequest(c *fiber.Ctx) error {
	//var SellerRequest model.SellerRequest
	//model.DB.First(&SellerRequest, "id = ?", c.Params("id"))
	sellerRequestID, _ := strconv.Atoi(c.Params("id"))
	SellerRequest, _ := db.Client.SellerRequest.Get(context.Background(), sellerRequestID)
	UserCount, _ := db.Client.User.Query().Where(user.PhoneNumber(SellerRequest.ContactNumber)).Count(context.Background())
	//var UserCount int64
	//model.DB.Model(model.User{}).Where("phone_number = ?", SellerRequest.ContactNumber).Count(&UserCount)
	if UserCount == 0 {
		//SellerRequest.Accepted = true
		//SellerRequest.UserID = c.Locals("AuthID").(uint)
		//model.DB.Save(&SellerRequest)
		getAdminUser, _ := db.Client.User.Get(context.Background(), c.Locals("AuthID").(int))
		_, err := db.Client.SellerRequest.UpdateOneID(SellerRequest.ID).SetAccepted(true).SetUser(getAdminUser).Save(context.Background())
		if err != nil {
			return c.Status(500).SendString("Please try again.")
		}

		hasPass, _ := myauth.HashPassword("12345678")
		//user := model.User{Name: SellerRequest.SellerName, PhoneNumber: SellerRequest.ContactNumber, Password: hasPass, Seller: true}
		//model.DB.Select("Name", "PhoneNumber", "Password", "Seller").Create(&user)
		_, err3 := db.Client.User.Create().SetName(SellerRequest.SellerName).SetPhoneNumber(SellerRequest.ContactNumber).SetPassword(hasPass).SetSeller(true).SetActive(true).Save(context.Background())
		if err3 != nil {
			return c.Status(500).SendString("Please try again.")
		}

		return c.SendStatus(201)
	} else {
		return c.Status(422).SendString("User with this phone number already available")
	}
}
func RemoveSellerRequest(c *fiber.Ctx) error {
	//var SellerRequest model.SellerRequest
	//model.DB.Delete(&SellerRequest, "id = ?", c.Params("id"))
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := db.Client.SellerRequest.UpdateOneID(id).SetDeletedAt(time.Now()).Save(context.Background())
	if err != nil {
		return c.Status(500).SendString("Please try again.")
	}
	return c.SendStatus(200)
}
func RemovePermanentSellerRequest(c *fiber.Ctx) error {
	var SellerRequest model.SellerRequest
	model.DB.Unscoped().Delete(&SellerRequest, "id = ?", c.Params("id"))
	id, _ := strconv.Atoi(c.Params("id"))
	err := db.Client.SellerRequest.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return c.Status(500).SendString("Please try again.")
	}
	return c.SendStatus(200)
}
func RecoverSellerRequest(c *fiber.Ctx) error {
	//var SellerRequest model.SellerRequest
	//model.DB.Unscoped().First(&SellerRequest, "id = ?", c.Params("id")).Update("deleted_at", nil)
	//model.DB.First(&SellerRequest, "id = ?", c.Params("id"))

	SellerRequest, err := db.Client.SellerRequest.Update().ClearDeletedAt().Save(context.Background())
	if err != nil {
		return c.Status(500).SendString("Please try again.")
	}
	return c.JSON(SellerRequest)
}

// seller request end

// seller shops start

func SellerShopsNonActivate(c *fiber.Ctx) error {
	//var sellerNonActiveShops []model.SellerShop
	//model.DB.Preload("ShopCategory").Where("active = ?", false).Find(&sellerNonActiveShops)
	sellerNonActiveShops, _ := db.Client.SellerShop.Query().WithGetShopCategory().Where(sellershop.Active(false)).Where(sellershop.DeletedAtIsNil()).All(context.Background())
	return c.Status(200).JSON(sellerNonActiveShops)
}
func SellerShopsActivate(c *fiber.Ctx) error {
	//var sellerActiveShops []model.SellerShop
	//model.DB.Preload("ShopCategory").Where("active = ?", true).Where("deleted_at IS NULL").Find(&sellerActiveShops)
	sellerActiveShops, _ := db.Client.SellerShop.Query().Where(sellershop.Active(true)).Where(sellershop.DeletedAtIsNil()).WithGetShopCategory().All(context.Background())
	return c.Status(200).JSON(sellerActiveShops)
}
func SellerShopsDeleted(c *fiber.Ctx) error {
	//var sellerActiveShops []model.SellerShop
	//model.DB.Preload("ShopCategory").Unscoped().Not("deleted_at IS NULL").Find(&sellerActiveShops)
	sellerActiveShops, _ := db.Client.SellerShop.Query().Where(sellershop.DeletedAtNotNil()).WithGetShopCategory().All(context.Background())
	return c.Status(200).JSON(sellerActiveShops)
}

func SellerShopsAll(c *fiber.Ctx) error {
	//var sellerActiveShops []model.SellerShop
	//model.DB.Preload("ShopCategory").Unscoped().Find(&sellerActiveShops)
	sellerActiveShops, _ := db.Client.SellerShop.Query().WithGetShopCategory().All(context.Background())
	return c.Status(200).JSON(sellerActiveShops)
}

func ActiveSellerShops(c *fiber.Ctx) error {
	//shop := new(model.SellerShop)
	//ShopFound := model.DB.Where("active = ?", false).First(&shop, "id = ?", c.Params("id"))
	ShopID, _ := strconv.Atoi(c.Params("id"))
	Shop, shopErr := db.Client.SellerShop.Query().Where(sellershop.Active(false)).Where(sellershop.ID(ShopID)).First(context.Background())
	if shopErr != nil {
		return c.SendStatus(404)
	}
	AdminUser, adminErr := db.Client.User.Get(context.Background(), c.Locals("AuthID").(int))
	if adminErr != nil {
		return c.SendStatus(404)
	}
	_, err := db.Client.SellerShop.UpdateOne(Shop).SetActive(true).SetAdmin(AdminUser).Save(context.Background())
	//shop.Active = true
	//shop.AdminID = c.Locals("AuthID").(uint)
	//ShopSave := model.DB.Save(&shop)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
func SoftDeleteSellerShops(c *fiber.Ctx) error {
	//shop := new(model.SellerShop)
	//model.DB.Find(&shop, "id = ?", c.Params("id")).Updates(model.SellerShop{Active: false, AdminID: c.Locals("AuthID").(uint)})
	//model.DB.Delete(&shop)
	AdminUser, adminErr := db.Client.User.Get(context.Background(), c.Locals("AuthID").(int))
	if adminErr != nil {
		return c.SendStatus(404)
	}
	ShopID, _ := strconv.Atoi(c.Params("id"))
	_, err := db.Client.SellerShop.UpdateOneID(ShopID).SetDeletedAt(time.Now()).SetAdmin(AdminUser).SetActive(false).Save(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
func PermanentDeleteSellerShops(c *fiber.Ctx) error {
	//shop := new(model.SellerShop)
	//model.DB.Unscoped().Where("id = ?", c.Params("id")).Find(&shop)
	//model.DB.Unscoped().Delete(&shop)
	ShopID, _ := strconv.Atoi(c.Params("id"))
	err := db.Client.SellerShop.DeleteOneID(ShopID).Exec(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
func RecoverDeleteSellerShops(c *fiber.Ctx) error {
	//shop := new(model.SellerShop)
	//model.DB.Model(&shop).Unscoped().Where("id = ?", c.Params("id")).Updates(map[string]interface{}{"deleted_at": nil, "admin_id": c.Locals("AuthID").(uint)})
	ShopID, _ := strconv.Atoi(c.Params("id"))
	_, err := db.Client.SellerShop.UpdateOneID(ShopID).ClearDeletedAt().SetActive(true).Save(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

// seller shops end

// product brand start

func BrandAll(c *fiber.Ctx) error {
	//var brands []model.Brand
	//model.DB.Find(&brands)
	brands, err := db.Client.Brand.Query().All(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(brands)
}

func BrandCreate(c *fiber.Ctx) error {
	type Brand struct {
		Name string `json:"name" gorm:"not null"`
	}
	brandInstance := new(Brand)
	if err := c.BodyParser(brandInstance); err != nil {
		return err
	}
	getBrand, err := db.Client.Brand.Create().SetName(brandInstance.Name).Save(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	//model.DB.Create(&brand)
	return c.Status(200).JSON(getBrand)

}

func BrandEdit(c *fiber.Ctx) error {
	type Brand struct {
		Name string `json:"name"`
	}
	brandInstance := new(Brand)
	if err := c.BodyParser(brandInstance); err != nil {
		return err
	}
	if brandInstance.Name != "" {
		fmt.Println("1")
		brandID, _ := strconv.Atoi(c.Params("id"))
		brandOj, brandOjErr := db.Client.Brand.Get(context.Background(), brandID)
		fmt.Println("1")
		if brandOjErr != nil {
			return c.SendStatus(500)
		}
		saveBrand, err := db.Client.Brand.UpdateOne(brandOj).SetName(brandInstance.Name).Save(context.Background())
		fmt.Println("1")
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(saveBrand)
	} else {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Name cannot be null.")
	}
	//err := model.DB.First(&brand, c.Params("id"))
	//if err.Error != nil {
	//	return c.JSON(fiber.Map{"status": fiber.StatusNotFound})
	//}
	//if err := c.BodyParser(brand); err != nil {
	//	return err
	//}
	//model.DB.Save(&brand)

}
func BrandSoftDelete(c *fiber.Ctx) error {
	//var brands model.Brand
	//model.DB.Delete(&brands, c.Params("id"))
	brandID, _ := strconv.Atoi(c.Params("id"))
	_, err := db.Client.Brand.UpdateOneID(brandID).SetDeletedAt(time.Now()).Save(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
func BrandDelete(c *fiber.Ctx) error {
	//var brands model.Brand
	//model.DB.Unscoped().Delete(&brands, c.Params("id"))
	brandID, _ := strconv.Atoi(c.Params("id"))
	err := db.Client.Brand.DeleteOneID(brandID).Exec(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func BrandRecoverDelete(c *fiber.Ctx) error {
	//var brands []model.Brand
	//model.DB.Unscoped().First(&brands, "id = ?", c.Params("id")).Update("deleted_at", nil)
	brandID, _ := strconv.Atoi(c.Params("id"))
	_, err := db.Client.Brand.UpdateOneID(brandID).ClearDeletedAt().Save(context.Background())
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

// product brand end
