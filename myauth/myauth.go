package myauth

import (
	"bongo/mixin"
	"bongo/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
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
		c.Status(fiber.StatusNoContent)
	}
	var user model.User
	model.DB.First(&user, "id = ?", data.Issuer)

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
	var user model.User
	result := model.DB.First(&user, "id = ?", data.Issuer)

	if result.Error != nil || !user.Admin {
		return c.SendStatus(204)
	}
	var roleNumber int
	if user.Admin {
		roleNumber = 10
	} else if user.Seller {
		roleNumber = 5
	}
	fmt.Println(user)
	return c.Status(200).JSON(fiber.Map{
		"id":           user.ID,
		"name":         user.Name,
		"phone_number": user.PhoneNumber,
		"role":         roleNumber,
	})
}
func GetSeller(c *fiber.Ctx) error {
	cookie := c.Cookies("bongoauth")
	err, data, _ := mixin.VerifyToken(cookie)
	if !err {
		c.Status(fiber.StatusNoContent)
	}
	var user model.User
	result := model.DB.First(&user, "id = ?", data.Issuer)

	if result.Error != nil || !user.Seller {
		return c.SendStatus(204)
	}
	return c.Status(200).JSON(fiber.Map{
		"id":           user.ID,
		"name":         user.Name,
		"phone_number": user.PhoneNumber,
		"seller":       user.Seller,
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
	var totalUser int64
	model.DB.Model(model.User{}).Where("phone_number = ?", body.PhoneNumber).Count(&totalUser)
	if totalUser > 0 {
		return c.Status(422).JSON(fiber.Map{
			"Phone number": "User is already exists!!",
		})
	}
	matched, _ := regexp.MatchString(`(^(01)[3-9]\d{8})$`, body.PhoneNumber)
	if !matched {
		return c.Status(422).JSON(fiber.Map{
			"Phone number": "Please provide a valid phone number.",
		})
	}
	hash, _ := HashPassword(body.Password) // ignore error for the sake of simplicity
	user := model.User{PhoneNumber: body.PhoneNumber, Password: hash, Name: body.Name}
	model.DB.Select("Name", "PhoneNumber", "Password").Create(&user)
	return c.SendStatus(200)
}

func UserLogin(c *fiber.Ctx) error {
	body := new(requestLogin)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Wrong entity.")
	}
	var user model.User
	err := model.DB.Where(model.User{PhoneNumber: body.PhoneNumber}).First(&user)
	if user.Admin {
		return c.Status(fiber.StatusBadRequest).SendString("Wrong credentials.")
	}
	if err.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Wrong credentials.")
	}
	match := CheckPasswordHash(body.Password, user.Password)
	if !match {
		return c.SendStatus(422)
	}
	token, tokenErr := mixin.GetToken(int(user.ID))
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
	var user model.User
	err := model.DB.Where(model.User{PhoneNumber: body.PhoneNumber}).First(&user)
	if err.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Wrong credentials.")
	}
	match := CheckPasswordHash(body.Password, user.Password)
	if !match || !user.Seller {
		return c.Status(fiber.StatusBadRequest).SendString("Wrong Credential")
	}

	token, tokenErr := mixin.GetToken(int(user.ID))
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
	var user model.User
	fmt.Println(user)
	err := model.DB.Where(model.User{PhoneNumber: body.PhoneNumber, AdminUserToken: body.AdminUserToken, AdminUserName: body.AdminUserName}).First(&user)
	if err.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Wrong credentials.")
	}
	match := CheckPasswordHash(body.Password, user.Password)
	if !match || !user.Admin {
		return c.Status(fiber.StatusBadRequest).SendString("Wrong Credential")
	}
	token, tokenErr := mixin.GetToken(int(user.ID))
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
	fmt.Println(body)
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).SendString("Body cannot parse")
	}
	matched, _ := regexp.MatchString(`(^(01)[3-9]\d{8})$`, body.PhoneNumber)
	if !matched {
		return c.Status(422).JSON(fiber.Map{
			"Phone number": "Please provide a valid phone number.",
		})
	}
	var totalUser int64
	model.DB.Model(model.User{}).Where("phone_number = ?", body.PhoneNumber).Count(&totalUser)
	if totalUser > 0 {
		return c.Status(422).JSON(fiber.Map{
			"Phone number": "User is already exists!!",
		})
	}
	totalUser = 0
	model.DB.Model(model.User{}).Where("admin_user_name = ?", body.AdminUserName).Count(&totalUser)
	if totalUser > 0 {
		return c.Status(422).JSON(fiber.Map{
			"Phone number": "Admin username is already exists!!",
		})
	}
	totalUser = 0
	model.DB.Model(model.User{}).Where("admin_user_token = ?", body.AdminUserToken).Count(&totalUser)
	if totalUser > 0 {
		return c.Status(422).JSON(fiber.Map{
			"Phone number": "Admin user token is already exists!!",
		})
	}

	hash, _ := HashPassword(body.Password) // ignore error for the sake of simplicity
	user := model.User{PhoneNumber: body.PhoneNumber, Password: hash, Name: body.Name, AdminUserName: body.AdminUserName, AdminUserToken: body.AdminUserToken, Admin: true}
	model.DB.Select("Name", "PhoneNumber", "Password", "AdminUserName", "AdminUserToken", "Admin").Create(&user)
	return c.SendStatus(200)
}
func SellerRequestPost(c *fiber.Ctx) error {
	request := new(model.SellerRequest)
	fmt.Println(request)
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse data",
		})
	}
	matched, _ := regexp.MatchString(`(^(01)[3-9]\d{8})$`, request.ContactNumber)
	if !matched {
		return c.Status(422).JSON(fiber.Map{
			"contact_number": "Contact number is not valid.",
		})
	}
	var count int64
	model.DB.Model(model.User{}).Where("phone_number = ?", request.ContactNumber).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"contact_number": "Contact number is already exists",
		})
	}
	model.DB.Model(model.SellerRequest{}).Where("contact_number = ?", request.ContactNumber).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"contact_number": "Contact number is already exists",
		})
	}
	model.DB.Select("SellerName", "ContactNumber", "ShopName", "ShopLocation", "TaxID", "ShopCategoryID").Create(&request)
	return c.SendStatus(200)
}
func UserLogout(c *fiber.Ctx) error {
	c.ClearCookie("bongoauth")
	return c.SendStatus(200)
}


