package model

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
var err error

type ShopCategory struct {
	gorm.Model
	ID            uint            `json:"id" gorm:"primaryKey;index;<-:create"`
	Name          string          `json:"name" form:"name"`
	Slug          string          `json:"slug" gorm:"unique;<-:create" form:"slug"`
	Image         string          `json:"image" form:"image"`
	Category      []Category      `json:"categories" form:"categories"`
	SellerRequest []SellerRequest `json:"seller_request" form:"seller_request"`
}

type Category struct {
	gorm.Model
	ID             uint         `json:"id" gorm:"primaryKey;index;<-:create"`
	Name           string       `json:"name" gorm:"not null"`
	Slug           string       `json:"slug"  gorm:"not null;unique;<-:create"`
	ShopCategoryID uint         `gorm:"index;not null" json:"shop_category_id"`
	ParentID       *uint        `gorm:"index" json:"parent_id"`
	Parent         *Category    `json:"parent"`
	Children       []Category   `gorm:"foreignkey:ParentID"`
	ShopCategory   ShopCategory `json:"shop_category"`
}
type Brand struct {
	gorm.Model
	ID            uint            `json:"id" gorm:"primaryKey;index;<-:create"`
	Name          string          `json:"name" gorm:"not null"`
	SellerProduct []SellerProduct `json:"seller_product"`
}
type Attribute struct {
	gorm.Model
	ID                           uint                           `json:"id" gorm:"primaryKey;index;<-:create"`
	Name                         string                         `json:"name"`
	SellerProductVariationValues []SellerProductVariationValues `json:"seller_product_variation_values"`
}

type User struct {
	gorm.Model
	ID                 uint              `json:"id" gorm:"primaryKey;index;<-:create"`
	Name               string            `json:"name" gorm:"type:varchar(150);not null"`
	PhoneNumber        string            `json:"phone_number" gorm:"type:varchar(11);unique;not null"`
	Password           string            `json:"-" gorm:"not null"`
	Admin              bool              `json:"admin" gorm:"default:false"`
	Staff              bool              `json:"staff" gorm:"default:false"`
	Seller             bool              `json:"seller" gorm:"default:false"`
	Active             bool              `json:"active" gorm:"default:false"`
	AdminUserName      string            `json:"admin_user_name" gorm:"unique"`
	AdminUserToken     string            `json:"admin_user_token" gorm:"unique"`
	SellerRequest      []SellerRequest   `json:"seller_request"`
	SellerShop         []SellerShop      `json:"seller_shop"`
	SellerProduct      []SellerProduct   `json:"seller_product"`
	Cart               []Cart            `json:"cart"`
	UserLocation       []UserLocation    `json:"user_location"`
	Checkout           []Checkout        `json:"checkout"`
	CheckoutProduct    []CheckoutProduct `json:"checkout_product"`
	SellingSeller      []CheckoutProduct `json:"selling_seller" gorm:"foreignKey:SellingSellerID"`
	AdminShopActivated []SellerShop      `gorm:"foreignKey:AdminID"`
}

type SellerRequest struct {
	gorm.Model
	ID             uint         `json:"id" gorm:"primaryKey;index;<-:create"`
	SellerName     string       `json:"seller_name" gorm:""`
	ShopName       string       `json:"shop_name" gorm:""`
	ContactNumber  string       `json:"contact_number" gorm:"type:varchar(11);unique"`
	ShopLocation   string       `json:"shop_location"`
	TaxID          string       `json:"tax_id"`
	Accepted       bool         `json:"accepted" gorm:"default:false"`
	UserID         uint         `json:"user_id" gorm:"index"`
	ShopCategoryID uint         `json:"shop_category_id" gorm:"index;not null"`
	ShopCategory   ShopCategory `json:"shop_category"`
	User           User         `json:"user"`
}

// seller start

type SellerShop struct {
	gorm.Model
	ID                uint                `json:"id" gorm:"primaryKey;index;<-:create"`
	Name              string              `form:"name" json:"name" gorm:"type:varchar(150)"`
	Slug              string              `json:"slug" gorm:"index,<-:create"`
	ContactNumber     string              `form:"contact_number" json:"contact_number" gorm:"type:varchar(11)"`
	Banner            string              `form:"banner" json:"banner"`
	ShopCategoryID    int                 `form:"shop_category_id" json:"shop_category_id" gorm:"index"`
	ShopCategory      ShopCategory        `json:"shop_category"`
	BusinessLocation  string              `form:"business_location" json:"business_location"`
	TaxID             string              `form:"tax_id" json:"tax_id"`
	Active            bool                `json:"active" gorm:"default:false"`
	UserID            uint                `json:"user_id"`
	User              User                `json:"user"`
	AdminID           uint                `json:"admin_id"`
	Admin             User                `json:"admin" gorm:"foreignkey:AdminID"`
	SellerProduct     []SellerProduct     `json:"seller_product"`
	SellerShopProduct []SellerShopProduct `json:"seller_shop_product"`
}
type SellerProduct struct {
	gorm.Model
	ID              uint            `json:"id" gorm:"primaryKey;<-:create;"`
	Name            string          `json:"name" form:"name" `
	Slug            string          `gorm:"unique;<-:create"`
	SellingPrice    decimal.Decimal `sql:"type:decimal(10,2)" json:"selling_price" form:"selling_price"`
	ProductPrice    decimal.Decimal `sql:"type:decimal(10,2)" json:"product_price" form:"product_price"`
	Quantity        int             `json:"quantity" form:"quantity" gorm:"default:0"`
	Active          bool            `json:"active" gorm:"default:false"`
	Description     string          `json:"description" form:"description" gorm:"type:text"`
	OfferPrice      int             `json:"offer_price" form:"offer_price"`
	OfferPriceStart time.Time       `json:"offer_price_start" form:"offer_price_start"`
	OfferPriceEnd   time.Time       `json:"offer_price_end" form:"offer_price_end"`
	NextStock       time.Time       `json:"next_stock" form:"next_stock"`

	BrandID                *uint                     `json:"brand_id" form:"brand_id" gorm:"index"`
	Brand                  Brand                    `json:"brand"`
	UserID                 uint                     `json:"user_id" form:"user_id" gorm:"index"`
	User                   User                     `json:"user" form:"user"`
	SellerShopID           uint                     `json:"seller_shop_id" form:"seller_shop_id" gorm:"index"`
	SellerShop             SellerShop               `json:"seller_shop" form:"seller_shop"`
	SellerProductImage     []SellerProductImage     `json:"product_image" form:"product_image"`
	SellerProductCategory  []SellerProductCategory  `json:"seller_product_category" form:"seller_product_category"`
	CartProduct            []CartProduct            `json:"cart_product" form:"cart_product"`
	CheckoutProduct        []CheckoutProduct        `json:"checkout_product"`
	SellerProductVariation []SellerProductVariation `json:"seller_product_variation"`
	SellerShopProduct      []SellerShopProduct      `json:"seller_shop_product"`
}
type SellerShopProduct struct {
	gorm.Model
	ID              uint          `json:"id" gorm:"primaryKey;index;<-:create"`
	SellerProductID uint          `json:"seller_product_id" gorm:"index"`
	SellerProduct   SellerProduct `json:"seller_product"`
	SellerShopID    uint          `json:"seller_shop_id" gorm:"index"`
	SellerShop      SellerShop    `json:"seller_shop"`
}

type SellerProductImage struct {
	gorm.Model
	ID              uint   `json:"id" gorm:"primaryKey;index;<-:create"`
	SellerProductID uint   `json:"seller_product_id"`
	Display         bool   `json:"display" gorm:"default:false"`
	Image           string `json:"image" form:"image"`
}
type SellerProductCategory struct {
	gorm.Model
	ID              uint `json:"id" gorm:"primaryKey;index;<-:create"`
	SellerProductID uint `form:"seller_product_id" json:"seller_product_id" gorm:"index"`
	CategoryID      uint `form:"category_id" json:"category_id" gorm:"index"`
}
type SellerProductVariation struct {
	gorm.Model
	ID                           uint                           `json:"id" gorm:"primaryKey;index;<-:create"`
	ProductPrice                 decimal.Decimal                ` form:"product_price" json:"product_price" sql:"type:decimal(10,2)"`
	SellingPrice                 decimal.Decimal                `form:"selling_price" json:"selling_price" sql:"type:decimal(10,2)"`
	Quantity                     int                            `form:"quantity" json:"quantity" gorm:"default:0"`
	SellerProductID              uint                           `json:"seller_product_id" gorm:"not null"`
	SellerProduct                SellerProduct                  `json:"seller_product"`
	Image                        string                         `json:"image" gorm:"not null"`
	SellerProductVariationValues []SellerProductVariationValues `json:"seller_product_variation_values"`
}
type SellerProductVariationValues struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey;index;<-:create"`
	Name        string `form:"name" json:"name"`
	Description string `json:"description" form:"description"`

	SellerProductVariationID uint                   `json:"seller_product_variation_id" gorm:"not null"`
	SellerProductVariation   SellerProductVariation `json:"seller_product_variation"`

	AttributeID uint      `json:"attribute_id" gorm:"not null"`
	Attribute   Attribute `json:"attribute"`
}

// seller end

// user start

type Cart struct {
	gorm.Model
	ID          uint          `json:"id" gorm:"primaryKey;index;<-:create"`
	Slug        string        `json:"slug" gorm:"<-:create"`
	UserID      uint          `json:"user_id" gorm:"index"`
	User        User          `json:"user"`
	CartProduct []CartProduct `json:"cart_product"`
	Checkout    Checkout      `json:"checkout"`
}
type CartProduct struct {
	gorm.Model
	ID                       uint                   `json:"id" gorm:"primaryKey;index;<-:create"`
	CartID                   uint                   `json:"cart_id" gorm:"index"`
	Cart                     Cart                   `json:"cart"`
	SellerProductID          uint                   `json:"seller_product_id" gorm:"index"`
	SellerProduct            SellerProduct          `json:"seller_product"`
	SellerProductVariationID uint                   `json:"seller_product_variation_id" gorm:"index"`
	SellerProductVariation   SellerProductVariation `json:"seller_product_variation"`
}
type UserLocation struct {
	gorm.Model
	ID            uint   `json:"id" gorm:"primaryKey;index;<-:create"`
	UserID        uint   `json:"user_id" gorm:"index"`
	User          User   `json:"user"`
	Area          string `json:"area"`
	Street        string `json:"street"`
	House         string `json:"house"`
	PostOffice    string `json:"post_office"`
	PostCode      string `json:"post_code"`
	PoliceStation string `json:"police_station"`
	City          string `json:"city"`
	Checkout      []Checkout
}

type Checkout struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey;index;<-:create"`
	Slug   string `json:"slug" gorm:"<-:create"`
	CartID int    `json:"cart_id" gorm:"index"`
	//Cart           Cart         `json:"cart"`
	TotalPrice      float32      `json:"total_price" sql:"type:decimal(10,2)"`
	UserLocationID  uint         `json:"user_location_id" gorm:"index"`
	UserLocation    UserLocation `json:"user_location"`
	Completed       bool         `json:"completed" gorm:"default:false"`
	UserID          uint         `json:"user_id" gorm:"index"`
	User            User         `json:"user"`
	CheckoutProduct []CheckoutProduct
}
type CheckoutProduct struct {
	gorm.Model
	ID              uint   `json:"id" gorm:"primaryKey;index;<-:create"`
	Slug            string `json:"slug" gorm:"<-:create"`
	CheckoutID      uint   `json:"checkout_id" gorm:"index"`
	Checkout        Checkout
	SellerProductID uint          `json:"seller_product_id" gorm:"index"`
	SellerProduct   SellerProduct `json:"seller_product"`
	Quantity        int           `json:"quantity"`
	SellingPrice    float32       `json:"selling_price" sql:"type:decimal(10,2)"`
	OfferPrice      float32       `json:"offer_price" sql:"type:decimal(10,2)"`
	Received        bool          `json:"received" gorm:"default:false"`
	Status          int           `json:"status" gorm:"default:0"`
	UserID          uint          `json:"user_id" gorm:"index"`
	User            User          `json:"user"`
	SellingSellerID uint          `json:"selling_seller_id" gorm:"index"`
	SellingSeller   User          `json:"seller" gorm:"foreignKey:SellingSellerID"`
}

// user end

// InitDatabase database connection
func InitDatabase() {
	//DB, err = gorm.Open(sqlite.Open("bongobitan.db"), &gorm.Config{})
	//dsn := "root:@tcp(localhost)/bongobitan?charset=utf8mb4&parseTime=True&loc=Local"
	//DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//DB, err = gorm.Open(postgres.New(postgres.Config{
	//	DSN:                  "user=postgres password=123456 dbname=bongobitan port=5432 sslmode=disable TimeZone=Asia/Dhaka",
	//	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	//}), &gorm.Config{})
	dsn := "user=postgres password=123456 dbname=bongobitan port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	err := DB.AutoMigrate(
		ShopCategory{},
		Category{},
		Brand{},
		Attribute{},

		User{},
		SellerRequest{},
		SellerShop{},
		SellerProduct{},
		SellerShopProduct{},
		SellerProductImage{},
		SellerProductCategory{},
		SellerProductVariation{},
		SellerProductVariationValues{},

		Cart{},
		CartProduct{},
		UserLocation{},
		Checkout{},
		CheckoutProduct{},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

}
