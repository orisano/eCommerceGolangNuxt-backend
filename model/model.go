package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
var err error

type ShopCategory struct {
	gorm.Model
	ID            uint            `json:"id" gorm:"primaryKey;index,<-:create"`
	Name          string          `json:"name" form:"name"`
	Slug          string          `json:"slug" gorm:"unique;<-:create" form:"slug"`
	Image         string          `json:"image" form:"image"`
	Category      []Category      `json:"categories" form:"categories"`
	SellerRequest []SellerRequest `json:"seller_request" form:"seller_request"`
}
type ProductVariation struct {
	gorm.Model
	ID                    uint `json:"id" gorm:"primaryKey;index,<-:create"`
	Name                  string
	SellerProductVariance []SellerProductVariance
}
type Category struct {
	gorm.Model
	ID             uint         `json:"id" gorm:"primaryKey;index,<-:create"`
	Name           string       `json:"name" gorm:"not null"`
	Slug           string       `json:"slug"  gorm:"not null;unique;<-:create"`
	ShopCategoryID uint         `gorm:"index;not null" json:"shop_category_id"`
	ParentID       *uint        `gorm:"index" json:"parent_id"`
	Parent         *Category    `json:"parent"`
	Children       []Category   `gorm:"foreignkey:ParentID"`
	ShopCategory   ShopCategory `json:"shop_category"`
}
type Brand struct {
	ID             uint         `json:"id" gorm:"primaryKey;index,<-:create"`
	Name           string       `json:"name" gorm:"not null"`
	ShopCategoryID int          `json:"shop_category_id"  gorm:"not null;index"`
	ShopCategory   ShopCategory `json:"shop_category"`
}
type User struct {
	gorm.Model
	ID              uint              `json:"id" gorm:"primaryKey;index,<-:create"`
	Name            string            `json:"name" gorm:"type:varchar(150);not null"`
	PhoneNumber     string            `json:"phone_number" gorm:"type:varchar(11);unique;not null"`
	Password        string            `json:"-" gorm:"type:size(255);not null"`
	Admin           bool              `json:"admin" gorm:"default:false"`
	Staff           bool              `json:"staff" gorm:"default:false"`
	Seller          bool              `json:"seller" gorm:"default:false"`
	Active          bool              `json:"active" gorm:"default:false"`
	AdminUserName   string            `json:"admin_user_name" gorm:"unique"`
	AdminUserToken  string            `json:"admin_user_token" gorm:"unique"`
	SellerRequest   []SellerRequest   `json:"seller_request"`
	SellerShop      []SellerShop      `json:"seller_shop"`
	SellerProduct   []SellerProduct   `json:"seller_product"`
	Cart            []Cart            `json:"cart"`
	UserLocation    []UserLocation    `json:"user_location"`
	Checkout        []Checkout        `json:"checkout"`
	CheckoutProduct []CheckoutProduct `json:"checkout_product"`
	SellingSeller   []CheckoutProduct `json:"selling_seller" gorm:"foreignKey:SellingSellerID"`
	AdminShopActivated []SellerShop `gorm:"foreignKey:AdminID"`

}

type SellerRequest struct {
	gorm.Model
	ID             uint         `json:"id" gorm:"primaryKey;index,<-:create"`
	SellerName     string       `json:"seller_name" gorm:"type:varchar(200)"`
	ShopName       string       `json:"shop_name" gorm:"type:size(255);"`
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
	ID               uint         `json:"id" gorm:"primaryKey;index,<-:create"`
	Name             string       `form:"name" json:"name" gorm:"type:varchar(150)"`
	Slug             string       `json:"slug" gorm:"index,<-:create"`
	ContactNumber    string       `form:"contact_number" json:"contact_number" gorm:"type:varchar(11)"`
	Banner           string       `form:"banner" json:"banner"`
	ShopCategoryID   int          `form:"shop_category_id" json:"shop_category_id" gorm:"index"`
	ShopCategory     ShopCategory `json:"shop_category"`
	BusinessLocation string       `form:"business_location" json:"business_location"`
	TaxID            string       `form:"tax_id" json:"tax_id"`
	Active           bool         `json:"active" gorm:"default:false"`
	UserID           uint         `json:"user_id"`
	User             User         `json:"user"`
	AdminID          uint         `json:"admin_id"`
	Admin            User         `json:"admin" gorm:"foreignkey:AdminID"`
	SellerProduct    []SellerProduct
}
type SellerProduct struct {
	gorm.Model
	ID                    uint      `json:"id" gorm:"primaryKey;index,<-:create"`
	Name                  string    `json:"name" form:"name" gorm:"type:size(255)"`
	Slug                  string    `gorm:"primaryKey;index,<-:create"`
	ProductPrice          float32   `sql:"type:decimal(10,2)" json:"product_price" form:"product_price"`
	SellingPrice          float32   `sql:"type:decimal(10,2)" json:"selling_price" form:"selling_price"`
	Quantity              int       `json:"quantity" form:"quantity"`
	Active                bool      `json:"active" gorm:"default:false"`
	Description           string    `json:"description" form:"description"`
	OfferPrice            float32   `sql:"type:decimal(10,2)" json:"offer_price" form:"offer_price"`
	OfferPriceStart       time.Time `json:"offer_price_start" form:"offer_price_start"`
	OfferPriceEnd         time.Time `json:"offer_price_end" form:"offer_price_end"`
	BrandID               int       `json:"brand_id" form:"brand_id" gorm:"index"`
	Brand                 Brand
	NextStock             time.Time               `json:"next_stock" form:"next_stock"`
	UserID                int                     `json:"user_id" form:"user_id" gorm:"index"`
	User                  User                    `json:"user" form:"user"`
	SellerShopID          int                     `json:"seller_shop_id" form:"seller_shop_id" gorm:"index"`
	SellerShop            SellerShop              `json:"seller_shop" form:"seller_shop"`
	SellerProductImage    []SellerProductImage    `json:"product_image" form:"product_image"`
	SellerProductOption   []SellerProductOption   `json:"product_option" form:"product_option"`
	SellerProductVariance []SellerProductVariance `json:"seller_product_variance" form:"seller_product_variance"`
	CartProduct           []CartProduct           `json:"cart_product" form:"cart_product"`
	CheckoutProduct       []CheckoutProduct       `json:"checkout_product"`
}
type SellerProductImage struct {
	gorm.Model
	ID              uint   `json:"id" gorm:"primaryKey;index,<-:create"`
	SellerProductID uint   `json:"seller_product_id"`
	Display         bool   `json:"display" gorm:"default:false"`
	Image           string `json:"image" form:"image"`
}
type SellerProductOption struct {
	gorm.Model
	ID              uint `json:"id" gorm:"primaryKey;index,<-:create"`
	SellerProductID uint `form:"seller_product_id" json:"seller_product_id" gorm:"index"`
	CategoryID      uint `form:"category_id" json:"category_id" gorm:"index"`
}
type SellerProductVariance struct {
	gorm.Model
	ID                       uint                   `json:"id" gorm:"primaryKey;index,<-:create"`
	SellerProductID          uint                   `form:"seller_product_id" json:"seller_product_id" gorm:"index"`
	SellerProduct            SellerProduct          `form:"seller_product" json:"seller_product"`
	ProductVariationID       uint                   `form:"product_variation_id" json:"product_variation_id" gorm:"index"`
	ProductVariation         ProductVariation       `form:"product_variation" json:"product_variation"`
	Color                    string                 `form:"color" json:"color"`
	ColorDescription         string                 `form:"color_description" json:"color_description"`
	Size                     string                 `form:"size" json:"size"`
	SizeDescription          string                 `form:"size_description" json:"size_description"`
	Style                    string                 `form:"style" json:"style"`
	StyleDescription         string                 `form:"style_description" json:"style_description"`
	ProductPrice             float32                `form:"product_price" json:"product_price" sql:"type:decimal(10,2)"`
	SellingPrice             float32                `form:"selling_price" json:"selling_price" sql:"type:decimal(10,2)"`
	Quantity                 int                    `form:"quantity" json:"quantity"`
	Image                    string                 `form:"image" json:"image"`
	AdminProductAttributesID uint                   `json:"admin_product_attributes_id" form:"admin_product_attributes_id"`
	AdminProductAttributes   AdminProductAttributes `json:"admin_product_attributes"`
	CartProduct              []CartProduct          `json:"cart_product"`
}

// seller end

// user start

type Cart struct {
	gorm.Model
	ID          uint          `json:"id" gorm:"primaryKey;index,<-:create"`
	Slug        string        `json:"slug" gorm:"<-:create"`
	UserID      uint          `json:"user_id" gorm:"index"`
	User        User          `json:"user"`
	CartProduct []CartProduct `json:"cart_product"`
	Checkout    Checkout      `json:"checkout"`
}
type CartProduct struct {
	gorm.Model
	ID                      uint                  `json:"id" gorm:"primaryKey;index,<-:create"`
	CartID                  uint                  `json:"cart_id" gorm:"index"`
	Cart                    Cart                  `json:"cart"`
	SellerProductID         uint                  `json:"seller_product_id" gorm:"index"`
	SellerProduct           SellerProduct         `json:"seller_product"`
	SellerProductVarianceID uint                  `json:"seller_product_variance_id" gorm:"index"`
	SellerProductVariance   SellerProductVariance `json:"seller_product_variance"`
}
type UserLocation struct {
	gorm.Model
	ID            uint   `json:"id" gorm:"primaryKey;index,<-:create"`
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
	ID     uint   `json:"id" gorm:"primaryKey;index,<-:create"`
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
	ID              uint   `json:"id" gorm:"primaryKey;index,<-:create"`
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

// admin start

type AdminProductAttributes struct {
	gorm.Model
	ID   uint   `json:"id" gorm:"primaryKey;index,<-:create"`
	Name string `json:"name"`
}

// admin end

// InitDatabase database connection
func InitDatabase() {
	DB, err = gorm.Open(sqlite.Open("bongobitan.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB.AutoMigrate(
		ShopCategory{},
		Category{},
		ProductVariation{},
		Brand{},
		User{},
		SellerRequest{},
		SellerShop{},
		SellerProduct{},
		SellerProductImage{},
		SellerProductVariance{},
		Cart{},
		CartProduct{},
		UserLocation{},
		Checkout{},
		CheckoutProduct{},
		AdminProductAttributes{},
	)
}
