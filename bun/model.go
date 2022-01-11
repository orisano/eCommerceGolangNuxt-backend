package bun

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"time"
)

var DB *bun.DB

//var err error

type ShopCategory struct {
	bun.BaseModel `bun:"table:shop_categories,alias:u"`
	ID            int64       `bun:"type:integer,pk,autoincrement" json:"id"`
	Name          string      `bun:"name,notnull" json:"name"`
	Slug          string      `bun:"slug,notnull" json:"slug"`
	Image         string      `bun:"image,notnull" json:"image"`
	Category      []*Category `bun:"category,rel:has-many,join:id=shop_category_id" json:"category"`
	//SellerRequest []*SellerRequest `bun:"seller_request,rel:has-many,join:id=shop_category_id" json:"seller_request"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*ShopCategory)(nil)

func (m *ShopCategory) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type Category struct {
	bun.BaseModel  `bun:"table:category,alias:u"`
	ID             int64         `bun:"type:integer,pk,autoincrement"`
	Name           string        `json:"name" bun:"name,notnull"`
	Slug           string        `json:"slug"  bun:"slug,notnull"`
	ShopCategoryID int64         `json:"shop_category_id" bun:"shop_category_id,allowzero"`
	ParentID       int64         `json:"parent_id" bun:"parent_id,allowzero"`
	Parent         *Category     `json:"parent" bun:"rel:belongs-to,join:parent_id=id"`
	Children       []*Category   `json:"children" bun:"rel:has-many,join:parent_id=id" `
	ShopCategory   *ShopCategory `json:"shop_category"  bun:"rel:belongs-to,join:shop_category_id=id"`
	CreatedAt      time.Time     `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time     `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*Category)(nil)

func (m *Category) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type Brand struct {
	bun.BaseModel `bun:"table:brand,alias:u"`
	ID            int64            `bun:"type:integer,pk,autoincrement"`
	Name          string           `json:"name" bun:"name,notnull"`
	SellerProduct []*SellerProduct `json:"seller_product" bun:"rel:has-many,join:id=brand_id"`
	CreatedAt     time.Time        `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time        `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*Brand)(nil)

func (m *Brand) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type Attribute struct {
	bun.BaseModel                `bun:"table:attribute,alias:u"`
	ID                           int64                          `bun:"type:integer,pk,autoincrement"`
	Name                         string                         `json:"name" bun:"name,notnull"`
	SellerProductVariationValues []SellerProductVariationValues `json:"seller_product_variation_values" bun:"rel:has-many,join:id=attribute_id"`
	CreatedAt                    time.Time                      `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt                    time.Time                      `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*Attribute)(nil)

func (m *Attribute) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type User struct {
	bun.BaseModel      `bun:"table:users,alias:u"`
	ID                 int64              `bun:"type:integer,pk,autoincrement"`
	Name               string             `json:"name" bun:"name,notnull"`
	PhoneNumber        string             `json:"phone_number" bun:"phone_number,notnull"`
	Password           string             `json:"-" bun:"-,notnull"`
	Admin              bool               `json:"admin" bun:"admin,default:false"`
	Staff              bool               `json:"staff" bun:"staff,default:false"`
	Seller             bool               `json:"seller" bun:"seller,default:false"`
	Active             bool               `json:"active" bun:"active,default:false"`
	AdminUserName      string             `json:"admin_user_name" gorm:"unique"`
	AdminUserToken     string             `json:"admin_user_token" gorm:"unique"`
	SellerRequest      []*SellerRequest   `json:"seller_request" bun:"rel:has-many,join:id=user_id"`
	SellerShop         []*SellerShop      `json:"seller_shop" bun:"rel:has-many,join:id=user_id"`
	SellerProduct      []*SellerProduct   `json:"seller_product" bun:"rel:has-many,join:id=user_id"`
	Cart               []*Cart            `json:"cart" bun:"rel:has-many,join:id=user_id"`
	UserLocation       []*UserLocation    `json:"user_location" bun:"rel:has-many,join:id=user_id"`
	Checkout           []*Checkout        `json:"checkout" bun:"rel:has-many,join:id=user_id"`
	CheckoutProduct    []*CheckoutProduct `json:"checkout_product" bun:"rel:has-many,join:id=user_id"`
	SellingSeller      []*CheckoutProduct `json:"selling_seller" gorm:"foreignKey:SellingSellerID" bun:"rel:has-many,join:id=selling_seller_id"`
	AdminShopActivated []*SellerShop      `gorm:"foreignKey:AdminID"`
	CreatedAt          time.Time          `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt          time.Time          `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (m *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type SellerRequest struct {
	bun.BaseModel  `bun:"table:seller_requests,alias:u"`
	ID             int64         `bun:"type:integer,pk,autoincrement"`
	SellerName     string        `json:"seller_name" bun:"seller_name"`
	ShopName       string        `json:"shop_name" bun:"shop_name"`
	ContactNumber  string        `json:"contact_number" bun:"contact_number,unique"`
	ShopLocation   string        `json:"shop_location" bun:"shop_location"`
	TaxID          string        `json:"tax_id"  bun:"tax_id"`
	Accepted       bool          `json:"accepted" bun:"accepted,default:false"`
	UserID         int64         `json:"user_id" bun:"user_id"`
	ShopCategoryID int64         `json:"shop_category_id" bun:"shop_category_id,notnull"`
	ShopCategory   *ShopCategory `json:"shop_category" bun:"rel:belongs-to,join:shop_category_id=id"`
	User           *User         `json:"user" bun:"rel:belongs-to,join:user_id=id"`
	CreatedAt      time.Time     `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time     `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*SellerRequest)(nil)

func (m *SellerRequest) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

// seller start

type SellerShop struct {
	bun.BaseModel `bun:"table:seller_shops,alias:u"`

	ID                int64                `bun:"type:integer,pk,autoincrement"`
	Name              string               `form:"name" json:"name" bun:"name"`
	Slug              string               `json:"slug" bun:"slug,unique"`
	ContactNumber     string               `form:"contact_number" json:"contact_number" bun:"contact_number,unique"`
	Banner            string               `form:"banner" json:"banner" bun:"banner,notnull"`
	ShopCategoryID    int64                `form:"shop_category_id" json:"shop_category_id" bun:"shop_category_id"`
	ShopCategory      *ShopCategory        `json:"shop_category" bun:"rel:belongs-to,join:shop_category_id=id"`
	BusinessLocation  string               `form:"business_location" json:"business_location" bun:"business_location"`
	TaxID             string               `form:"tax_id" json:"tax_id" bun:"tax_id"`
	Active            bool                 `json:"active" bun:"active,default:false"`
	UserID            int64                `json:"user_id" bun:"user_id"`
	User              *User                `json:"user" bun:"rel:belongs-to,join:user_id=id"`
	AdminID           int64                `json:"admin_id" bun:"admin_id"`
	Admin             *User                `json:"admin" bun:"rel:belongs-to,join:admin_id=id"`
	SellerProduct     []*SellerProduct     `json:"seller_product" bun:"rel:has-many,join:id=seller_shop_id"`
	SellerShopProduct []*SellerShopProduct `json:"seller_shop_product" bun:"rel:has-many,join:id=seller_shop_id"`
	CreatedAt         time.Time            `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt         time.Time            `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*SellerShop)(nil)

func (m *SellerShop) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type SellerProduct struct {
	bun.BaseModel          `bun:"table:seller_products,alias:u"`
	ID                     int64                     `bun:"type:integer,pk,autoincrement"`
	Name                   string                    `json:"name" form:"name" bun:"name,notnull"`
	Slug                   string                    `bun:"name,notnull" json:"slug"`
	SellingPrice           decimal.Decimal           `bun:"selling_price,nullzero" sql:"type:decimal(10,2)" json:"selling_price" form:"selling_price"`
	ProductPrice           decimal.Decimal           `bun:"product_price,nullzero" sql:"type:decimal(10,2)" json:"product_price" form:"product_price"`
	Quantity               int                       `json:"quantity" form:"quantity" bun:"quantity,default:0"`
	Active                 bool                      `json:"active" bun:"default:false"`
	Description            string                    `json:"description" form:"description" bun:"type:text"`
	OfferPrice             int                       `json:"offer_price" form:"offer_price" bun:"offer_price,default:0"`
	OfferPriceStart        time.Time                 `json:"offer_price_start" form:"offer_price_start" bun:"offer_price_start"`
	OfferPriceEnd          time.Time                 `json:"offer_price_end" form:"offer_price_end" bun:"offer_price_end"`
	NextStock              time.Time                 `json:"next_stock" form:"next_stock" bun:"next_stock"`
	BrandID                int64                     `json:"brand_id" form:"brand_id" bun:"brand_id,allowzero"`
	Brand                  *Brand                    `json:"brand" bun:"rel:belongs-to,join:brand_id=id"`
	UserID                 uint                      `json:"user_id" form:"user_id" bun:"user_id"`
	User                   *User                     `json:"user" form:"user" bun:"rel:belongs-to,join:user_id=id"`
	SellerShopID           uint                      `json:"seller_shop_id" form:"seller_shop_id" bun:"seller_shop_id"`
	SellerShop             *SellerShop               `json:"seller_shop" form:"seller_shop" bun:"rel:belongs-to,join:seller_shop_id=id"`
	SellerProductImage     []*SellerProductImage     `json:"product_image" form:"product_image" bun:"rel:has-many,join:id=seller_product_id"`
	SellerProductCategory  []*SellerProductCategory  `json:"seller_product_category" form:"seller_product_category" bun:"rel:has-many,join:id=seller_product_id"`
	CartProduct            []*CartProduct            `json:"cart_product" form:"cart_product" bun:"rel:has-many,join:id=seller_product_id"`
	CheckoutProduct        []*CheckoutProduct        `json:"checkout_product" bun:"rel:has-many,join:id=seller_product_id"`
	SellerProductVariation []*SellerProductVariation `json:"seller_product_variation" bun:"rel:has-many,join:id=seller_product_id"`
	SellerShopProduct      []*SellerShopProduct      `json:"seller_shop_product" bun:"rel:has-many,join:id=seller_product_id"`
	CreatedAt              time.Time                 `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt              time.Time                 `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*SellerProduct)(nil)

func (m *SellerProduct) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type SellerShopProduct struct {
	bun.BaseModel   `bun:"table:seller_shop_products,alias:u"`
	ID              int64          `bun:"type:integer,pk,autoincrement"`
	SellerProductID int64          `json:"seller_product_id" bun:"seller_product_id"`
	SellerProduct   *SellerProduct `json:"seller_product" bun:"rel:belongs-to,join:seller_product_id=id"`
	SellerShopID    int64          `json:"seller_shop_id" bun:"seller_shop_id"`
	SellerShop      *SellerShop    `json:"seller_shop" bun:"rel:belongs-to,join:seller_shop_id=id"`
	CreatedAt       time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*SellerShopProduct)(nil)

func (m *SellerShopProduct) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type SellerProductImage struct {
	bun.BaseModel   `bun:"table:seller_product_images,alias:u"`
	ID              int64     `bun:"type:integer,pk,autoincrement"`
	SellerProductID uint      `json:"seller_product_id" bun:"seller_product_id"`
	Display         bool      `json:"display" bun:"default:false"`
	Image           string    `json:"image" form:"image" bun:"image"`
	CreatedAt       time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*SellerProductImage)(nil)

func (m *SellerProductImage) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type SellerProductCategory struct {
	bun.BaseModel   `bun:"table:seller_product_categories,alias:u"`
	ID              int64     `bun:"type:integer,pk,autoincrement"`
	SellerProductID int64     `form:"seller_product_id" json:"seller_product_id" bun:"seller_product_id"`
	CategoryID      int64     `form:"category_id" json:"category_id" bun:"category_id"`
	CreatedAt       time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*SellerProductCategory)(nil)

func (m *SellerProductCategory) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type SellerProductVariation struct {
	bun.BaseModel                `bun:"table:seller_product_variations,alias:u"`
	ID                           int64                           `bun:"type:integer,pk,autoincrement"`
	ProductPrice                 decimal.Decimal                 `bun:"product_price" form:"product_price" json:"product_price" sql:"type:decimal(10,2)"`
	SellingPrice                 decimal.Decimal                 `bun:"selling_price" form:"selling_price" json:"selling_price" sql:"type:decimal(10,2)"`
	Quantity                     int                             `form:"quantity" json:"quantity" bun:"quantity,default:0"`
	SellerProductID              int64                           `bun:"seller_product_id" json:"seller_product_id" gorm:"not null"`
	SellerProduct                *SellerProduct                  `bun:"rel:belongs-to,join:seller_product_id=id" json:"seller_product"`
	Image                        string                          `json:"image" bun:"image,notnull"`
	SellerProductVariationValues []*SellerProductVariationValues `bun:"rel:has-many,join:id=seller_product_variation" json:"seller_product_variation_values"`
	CreatedAt                    time.Time                       `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt                    time.Time                       `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*SellerProductVariation)(nil)

func (m *SellerProductVariation) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type SellerProductVariationValues struct {
	bun.BaseModel            `bun:"table:seller_product_variation_values,alias:u"`
	ID                       int64                   `bun:"type:integer,pk,autoincrement"`
	Name                     string                  `form:"name" json:"name" bun:"name"`
	Description              string                  `json:"description" form:"description" bun:"description"`
	SellerProductVariationID int64                   `json:"seller_product_variation_id" bun:"seller_product_variation_id,notnull,"`
	SellerProductVariation   *SellerProductVariation `json:"seller_product_variation" bun:"rel:belongs-to,join:seller_product_variation_id=id"`
	AttributeID              int64                   `json:"attribute_id" bun:"attribute_id,notnull"`
	Attribute                *Attribute              `json:"attribute" bun:"rel:belongs-to,join:attribute_id=id"`
	CreatedAt                time.Time               `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt                time.Time               `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*SellerProductVariationValues)(nil)

func (m *SellerProductVariationValues) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

// seller end

// user start

type Cart struct {
	bun.BaseModel `bun:"table:carts,alias:u"`
	ID            int64          `bun:"type:integer,pk,autoincrement"`
	Slug          string         `json:"slug" bun:"slug,unique"`
	UserID        int64          `json:"user_id" bun:"user_id"`
	User          *User          `json:"user" bun:"rel:belongs-to,join:user_id=id"`
	CartProduct   []*CartProduct `json:"cart_product" bun:"rel:has-many,join:id=cart_id"`
	Checkout      Checkout       `json:"checkout" bun:"checkout"`
	CreatedAt     time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*Cart)(nil)

func (m *Cart) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type CartProduct struct {
	bun.BaseModel            `bun:"table:cart_products,alias:u"`
	ID                       int64                   `bun:"type:integer,pk,autoincrement"`
	CartID                   int64                   `json:"cart_id" bun:"cart_id"`
	Cart                     *Cart                   `json:"cart" bun:"rel:belongs-to,join:cart_id=id"`
	Quantity                 int                     `json:"quantity" bun:"quantity"`
	SellerProductID          int64                   `json:"seller_product_id"  bun:"seller_product_id"`
	SellerProduct            *SellerProduct          `json:"seller_product" bun:"rel:belongs-to,join:seller_product_id=id"`
	SellerProductVariationID int64                   `json:"seller_product_variation_id" bun:"seller_product_variation_id,allowzero"`
	SellerProductVariation   *SellerProductVariation `json:"seller_product_variation" bun:"rel:belongs-to,join:seller_product_variation_id=id"`
	CreatedAt                time.Time               `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt                time.Time               `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*CartProduct)(nil)

func (m *CartProduct) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type UserLocation struct {
	bun.BaseModel `bun:"table:user_locations,alias:u"`
	ID            int64       `bun:"type:integer,pk,autoincrement"`
	UserID        int64       `bun:"user_id" json:"user_id"`
	User          *User       `bun:"rel:belongs-to,join:user_id=id" json:"user"`
	Area          string      `bun:"area" json:"area"`
	Street        *string     `bun:"street" json:"street"`
	House         *string     `bun:"house" json:"house"`
	PostOffice    string      `bun:"post_office" json:"post_office"`
	PostCode      string      `bun:"post_code" json:"post_code"`
	PoliceStation string      `bun:"police_station" json:"police_station"`
	City          string      `bun:"city" json:"city"`
	Checkout      []*Checkout `bun:"rel:has-many,join:id=user_location_id" json:"checkout"`
	CreatedAt     time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*UserLocation)(nil)

func (m *UserLocation) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type Checkout struct {
	bun.BaseModel   `bun:"table:checkouts,alias:u"`
	ID              int64              `bun:"type:integer,pk,autoincrement"`
	CartID          int64              `bun:"cart_id" json:"cart_id"`
	TotalPrice      decimal.Decimal    `bun:"total_price" json:"total_price" sql:"type:decimal(10,2)"`
	UserLocationID  int64              `bun:"user_location_id" json:"user_location_id"`
	UserLocation    *UserLocation      `bun:"rel:belongs-to,join:user_location_id=id" json:"user_location"`
	Completed       bool               `bun:"completed,default:false" json:"completed"`
	UserID          int64              `bun:"user_id" json:"user_id"`
	User            *User              `bun:"rel:belongs-to,join:user_id=id" json:"user"`
	CheckoutProduct []*CheckoutProduct `bun:"rel:has-many,join:id=checkout_id"`
	CreatedAt       time.Time          `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time          `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*Checkout)(nil)

func (m *Checkout) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type CheckoutProduct struct {
	bun.BaseModel            `bun:"table:checkout_products,alias:u"`
	ID                       int64                   `bun:"type:integer,pk,autoincrement"`
	CheckoutID               int64                   `bun:"checkout_id" json:"checkout_id"`
	Checkout                 *Checkout               `bun:"rel:belongs-to,join:checkout_id=id" json:"checkout"`
	SellerProductID          int64                   `bun:"seller_product_id" json:"seller_product_id" gorm:"index"`
	SellerProductVariationID int64                   `bun:"seller_product_variation_id,nullzero" json:"seller_product_variation_id"`
	SellerProductVariation   *SellerProductVariation `bun:"rel:belongs-to,join:seller_product_variation_id=id"  json:"seller_product_variation"`
	Quantity                 int                     `bun:"quantity" json:"quantity"`
	SellingPrice             decimal.Decimal         `bun:"selling_price" json:"selling_price" sql:"type:decimal(10,2)"`
	OfferPrice               int                     `bun:"offer_price" json:"offer_price" sql:"type:decimal(10,2)"`
	Received                 bool                    `bun:"received,default:false" json:"received"`
	Status                   int                     `bun:"status,default:0" json:"status"`
	UserID                   int64                   `bun:"user_id" json:"user_id"`
	User                     *User                   `bun:"rel:belongs-to,join:user_id=id" json:"user"`
	SellingSellerID          int64                   `bun:"selling_seller_id" json:"selling_seller_id"`
	SellingSeller            *User                   `bun:"rel:belongs-to,join:selling_seller_id=id" json:"seller"`
	SellerProduct            *SellerProduct          `bun:"rel:belongs-to,join:seller_product_id=id" json:"seller_product"`
	CreatedAt                time.Time               `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt                time.Time               `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `json:"deleted_at" bun:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*CheckoutProduct)(nil)

func (m *CheckoutProduct) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

//var err error

func InitDatabase() {
	dsn := "postgres://postgres:@localhost:5432/bongobitan?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB = bun.NewDB(sqldb, pgdialect.New())
	if err:= DB.Ping();err != nil {
		fmt.Println("Cannot connect")
	}
	// Create users table.
	if _, err := DB.NewCreateTable().Model((*ShopCategory)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*Category)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*Brand)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*Attribute)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*User)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*SellerRequest)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*SellerShop)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*SellerProduct)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*SellerShopProduct)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*SellerProductImage)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*SellerProductCategory)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*SellerProductVariation)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*SellerProductVariationValues)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*Cart)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*CartProduct)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*UserLocation)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*Checkout)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}
	if _, err := DB.NewCreateTable().Model((*CheckoutProduct)(nil)).Exec(context.Background()); err != nil {
		fmt.Println(err)
	}

	// Drop users table.
	//_, err = db.NewDropTable().Model((*ShopCategory)(nil)).Exec(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//}
}
