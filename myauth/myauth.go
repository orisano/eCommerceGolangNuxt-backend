package myauth

import (
	"bongo/db"
	"bongo/ent/sellerrequest"
	"bongo/ent/user"
	"bongo/mixin"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strconv"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}

type requestRegister struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type requestLogin struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("bongoauth")
	err, data, _ := mixin.VerifyToken(cookie)
	if !err {
		return c.SendStatus(fiber.StatusNoContent)
	}
	id, _ := strconv.Atoi(data.Issuer)
	user, userErr := db.Client.User.Get(context.Background(), id)
	if userErr != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}
	//var user model.User
	//model.DB.First(&user, "id = ?", data.Issuer)

	return c.Status(200).JSON(fiber.Map{
		"id":           user.ID,
		"name":         user.Name,
		"phone_number": user.PhoneNumber,
	})
}
func GetAdmin(c *fiber.Ctx) error {
	cookie := c.Cookies("bongoauth")
	err, data, _ := mixin.VerifyToken(cookie)
	if !err {
		c.Status(fiber.StatusNoContent)
	}
	id, _ := strconv.Atoi(data.Issuer)
	Admin, adErr := db.Client.User.Get(context.Background(), id)

	if adErr != nil || !Admin.Admin {
		return c.SendStatus(204)
	}
	var roleNumber int
	if Admin.Admin {
		roleNumber = 10
	} else if Admin.Seller {
		roleNumber = 5
	}
	return c.Status(200).JSON(fiber.Map{
		"id":           Admin.ID,
		"name":         Admin.Name,
		"phone_number": Admin.PhoneNumber,
		"role":         roleNumber,
	})
}
func GetSeller(c *fiber.Ctx) error {
	cookie := c.Cookies("bongoauth")
	err, data, _ := mixin.VerifyToken(cookie)
	if !err {
		c.Status(fiber.StatusNoContent)
	}
	id, _ := strconv.Atoi(data.Issuer)
	myUser, adErr := db.Client.User.Get(context.Background(), id)

	if adErr != nil || !myUser.Seller {
		return c.SendStatus(204)
	}
	return c.Status(200).JSON(fiber.Map{
		"id":           myUser.ID,
		"name":         myUser.Name,
		"phone_number": myUser.PhoneNumber,
		"seller":       myUser.Seller,
	})
}

func GetCSRF(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func UserRegister(c *fiber.Ctx) error {
	body := new(requestRegister)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	totalUser, _ := db.Client.User.Query().Where(user.PhoneNumber(body.PhoneNumber)).Count(context.Background())

	if totalUser > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Phone number is already exists!!")
	}
	matched, _ := regexp.MatchString(`(^(01)[3-9]\d{8})$`, body.PhoneNumber)
	if !matched {
		return c.Status(422).SendString("Please provide a valid phone number.")
	}
	hash, _ := HashPassword(body.Password)
	saveUser, errSave := db.Client.User.Create().SetPhoneNumber(body.PhoneNumber).SetPassword(hash).SetName(body.Name).Save(context.Background())
	if errSave != nil {
		return c.SendStatus(500)
	}
	//user := model.User{PhoneNumber: body.PhoneNumber, Password: hash, Name: body.Name}
	//model.DB.Select("Name", "PhoneNumber", "Password").Create(&user)
	// auto login
	token, tokenErr := mixin.GetToken(saveUser.ID)
	if tokenErr != nil {
		return c.Status(500).SendString("Cannot purge token.")
	}
	c.Cookie(&fiber.Cookie{
		Name:     "bongoauth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return c.SendStatus(200)
}

func UserLogin(c *fiber.Ctx) error {
	body := new(requestLogin)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Wrong entity.")
	}
	//var user model.User
	//err := model.DB.Where(model.User{PhoneNumber: body.PhoneNumber}).First(&user)
	User, err := db.Client.User.Query().Where(user.PhoneNumber(body.PhoneNumber)).First(context.Background())
	//if User.Admin {
	//	return c.Status(fiber.StatusBadRequest).SendString("Wrong credentials.")
	//}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Wrong credentials.")
	}
	pass := User.Password
	match := CheckPasswordHash(body.Password, pass)

	if !match {
		return c.Status(fiber.StatusBadRequest).SendString("Wrong credentials.")
	}

	token, tokenErr := mixin.GetToken(User.ID)
	if tokenErr != nil {
		return c.Status(500).SendString("Cannot purge token.")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "bongoauth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return c.SendStatus(200)
}
func SellerLogin(c *fiber.Ctx) error {
	body := new(requestLogin)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Wrong credentials.")
	}
	//var user model.User
	//err := model.DB.Where(model.User{PhoneNumber: body.PhoneNumber}).First(&user)
	User, err := db.Client.User.Query().Where(user.PhoneNumber(body.PhoneNumber)).First(context.Background())
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Wrong credentials.")
	}
	match := CheckPasswordHash(body.Password, User.Password)

	if !match || !User.Seller {
		return c.Status(fiber.StatusBadRequest).SendString("Wrong Credential")
	}

	token, tokenErr := mixin.GetToken(User.ID)
	if tokenErr != nil {
		return c.Status(500).SendString("Cannot purge token.")
	}
	c.Cookie(&fiber.Cookie{
		Name:     "bongoauth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return c.SendStatus(200)
}

func AdminLogin(c *fiber.Ctx) error {
	type AdminLoginRequest struct {
		PhoneNumber    string `json:"phone_number"`
		Password       string `json:"password"`
		AdminUserName  string `json:"admin_user_name"`
		AdminUserToken string `json:"admin_user_token"`
	}
	body := new(AdminLoginRequest)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Wrong credentials.")
	}
	//var user model.User
	//fmt.Println(user)
	//err := model.DB.Where(model.User{PhoneNumber: body.PhoneNumber, AdminUserToken: body.AdminUserToken, AdminUserName: body.AdminUserName}).First(&user)
	userGet, err := db.Client.User.Query().Where(user.PhoneNumber(body.PhoneNumber)).Where(user.AdminUserName(body.AdminUserName)).Where(user.AdminUserToken(body.AdminUserToken)).First(context.Background())

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Wrong credentials.")
	}
	match := CheckPasswordHash(body.Password, userGet.Password)
	if !match || !userGet.Admin {
		return c.Status(fiber.StatusBadRequest).SendString("Wrong Credential")
	}

	token, tokenErr := mixin.GetToken(userGet.ID)
	if tokenErr != nil {
		return c.Status(500).SendString("Cannot purge token.")
	}
	c.Cookie(&fiber.Cookie{
		Name:     "bongoauth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return c.SendStatus(200)
}
func AdminRegister(c *fiber.Ctx) error {
	type AdminRegisterRequest struct {
		Name           string `json:"name"`
		PhoneNumber    string `json:"phone_number"`
		Password       string `json:"password"`
		AdminUserName  string `json:"admin_user_name"`
		AdminUserToken string `json:"admin_user_token"`
	}
	body := new(AdminRegisterRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).SendString("Body cannot parse")
	}
	matched, _ := regexp.MatchString(`(^(01)[3-9]\d{8})$`, body.PhoneNumber)
	if !matched {
		return c.Status(422).JSON(fiber.Map{
			"Phone number": "Please provide a valid phone number.",
		})
	}
	//var totalUser int64
	//model.DB.Model(model.User{}).Where("phone_number = ?", body.PhoneNumber).Count(&totalUser)
	totalUser, _ := db.Client.User.Query().Where(user.PhoneNumber(body.PhoneNumber)).Count(context.Background())
	if totalUser > 0 {
		return c.Status(422).SendString("User is already exists!!")
	}
	//totalUser = 0
	//model.DB.Model(model.User{}).Where("admin_user_name = ?", body.AdminUserName).Count(&totalUser)
	totalAdminUser, _ := db.Client.User.Query().Where(user.AdminUserName(body.AdminUserName)).Where(user.AdminUserToken(body.AdminUserToken)).Count(context.Background())
	if totalAdminUser > 0 {
		return c.Status(422).SendString("User is already exists!!")
	}
	hash, _ := HashPassword(body.Password)
	_, err := db.Client.User.Create().SetName(body.Name).SetPassword(hash).SetPhoneNumber(body.PhoneNumber).SetAdminUserName(body.AdminUserName).SetAdmin(true).SetActive(true).SetAdminUserToken(body.AdminUserToken).Save(context.Background())
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("Something is wrong. Try try again.")
	}
	return c.SendStatus(200)
}
func SellerRequestPost(c *fiber.Ctx) error {
	type tempRequest struct {
		SellerName     string `json:"seller_name"`
		ShopName       string `json:"shop_name"`
		ContactNumber  string `json:"contact_number"`
		ShopLocation   string `json:"shop_location"`
		TaxID          string `json:"tax_id"`
		ShopCategoryID int    `json:"shop_category_id"`
	}
	//body := new(model.SellerRequest)
	body := new(tempRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse data",
		})
	}
	matched, _ := regexp.MatchString(`(^(01)[3-9]\d{8})$`, body.ContactNumber)
	if !matched {
		return c.Status(422).SendString("Contact number is not valid.")
	}
	//var count int64
	//model.DB.Model(model.User{}).Where("phone_number = ?", request.ContactNumber).Count(&count)
	totalUser, _ := db.Client.User.Query().Where(user.PhoneNumber(body.ContactNumber)).Count(context.Background())

	if totalUser > 0 {
		return c.Status(422).SendString("Contact number is not valid.")

	}
	count, err := db.Client.SellerRequest.Query().Where(sellerrequest.ContactNumber(body.ContactNumber)).Count(context.Background())
	if err != nil {
		return c.Status(500).SendString("Try again.")
	}
	if count > 0 {
		return c.Status(422).SendString("Contact number is not valid.")

	}
	//model.DB.Model(model.SellerRequest{}).Where("contact_number = ?", request.ContactNumber).Count(&count)
	//if count > 0 {
	//	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
	//		"contact_number": "Contact number is already exists",
	//	})
	//}
	//model.DB.Select("SellerName", "ContactNumber", "ShopName", "ShopLocation", "TaxID", "ShopCategoryID").Create(&request)
	//shopCategoryID, _ := strconv.Atoi(body.ShopCategoryID)
	shopCategory, catErr := db.Client.ShopCategory.Get(context.Background(),body.ShopCategoryID)
	if catErr != nil {
		return c.Status(500).SendString("Try again.")
	}
	_, errSeller :=db.Client.SellerRequest.Create().SetSellerName(body.SellerName).SetContactNumber(body.ContactNumber).SetShopName(body.ShopName).SetShopLocation(body.ShopLocation).SetTaxID(body.TaxID).SetShopCategory(shopCategory).Save(context.Background())
	//hash, _ := HashPassword("123456")
	//_,err :=db.Client.User.Create().SetName(body.SellerName).SetPhoneNumber(body.ContactNumber).SetPassword(hash).SetSeller(true).Save(context.Background())
	if errSeller != nil {
		return c.Status(500).SendString("Try again.")
	}
	return c.SendStatus(200)
}
func UserLogout(c *fiber.Ctx) error {
	c.ClearCookie("bongoauth")
	return c.SendStatus(200)
}
