// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Attributes", testAttributes)
	t.Run("Brands", testBrands)
	t.Run("CartProducts", testCartProducts)
	t.Run("Carts", testCarts)
	t.Run("Categories", testCategories)
	t.Run("CheckoutProducts", testCheckoutProducts)
	t.Run("Checkouts", testCheckouts)
	t.Run("SellerProductCategories", testSellerProductCategories)
	t.Run("SellerProductImages", testSellerProductImages)
	t.Run("SellerProductVariationValues", testSellerProductVariationValues)
	t.Run("SellerProductVariations", testSellerProductVariations)
	t.Run("SellerProducts", testSellerProducts)
	t.Run("SellerRequests", testSellerRequests)
	t.Run("SellerShopProducts", testSellerShopProducts)
	t.Run("SellerShops", testSellerShops)
	t.Run("ShopCategories", testShopCategories)
	t.Run("UserLocations", testUserLocations)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Attributes", testAttributesDelete)
	t.Run("Brands", testBrandsDelete)
	t.Run("CartProducts", testCartProductsDelete)
	t.Run("Carts", testCartsDelete)
	t.Run("Categories", testCategoriesDelete)
	t.Run("CheckoutProducts", testCheckoutProductsDelete)
	t.Run("Checkouts", testCheckoutsDelete)
	t.Run("SellerProductCategories", testSellerProductCategoriesDelete)
	t.Run("SellerProductImages", testSellerProductImagesDelete)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesDelete)
	t.Run("SellerProductVariations", testSellerProductVariationsDelete)
	t.Run("SellerProducts", testSellerProductsDelete)
	t.Run("SellerRequests", testSellerRequestsDelete)
	t.Run("SellerShopProducts", testSellerShopProductsDelete)
	t.Run("SellerShops", testSellerShopsDelete)
	t.Run("ShopCategories", testShopCategoriesDelete)
	t.Run("UserLocations", testUserLocationsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Attributes", testAttributesQueryDeleteAll)
	t.Run("Brands", testBrandsQueryDeleteAll)
	t.Run("CartProducts", testCartProductsQueryDeleteAll)
	t.Run("Carts", testCartsQueryDeleteAll)
	t.Run("Categories", testCategoriesQueryDeleteAll)
	t.Run("CheckoutProducts", testCheckoutProductsQueryDeleteAll)
	t.Run("Checkouts", testCheckoutsQueryDeleteAll)
	t.Run("SellerProductCategories", testSellerProductCategoriesQueryDeleteAll)
	t.Run("SellerProductImages", testSellerProductImagesQueryDeleteAll)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesQueryDeleteAll)
	t.Run("SellerProductVariations", testSellerProductVariationsQueryDeleteAll)
	t.Run("SellerProducts", testSellerProductsQueryDeleteAll)
	t.Run("SellerRequests", testSellerRequestsQueryDeleteAll)
	t.Run("SellerShopProducts", testSellerShopProductsQueryDeleteAll)
	t.Run("SellerShops", testSellerShopsQueryDeleteAll)
	t.Run("ShopCategories", testShopCategoriesQueryDeleteAll)
	t.Run("UserLocations", testUserLocationsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Attributes", testAttributesSliceDeleteAll)
	t.Run("Brands", testBrandsSliceDeleteAll)
	t.Run("CartProducts", testCartProductsSliceDeleteAll)
	t.Run("Carts", testCartsSliceDeleteAll)
	t.Run("Categories", testCategoriesSliceDeleteAll)
	t.Run("CheckoutProducts", testCheckoutProductsSliceDeleteAll)
	t.Run("Checkouts", testCheckoutsSliceDeleteAll)
	t.Run("SellerProductCategories", testSellerProductCategoriesSliceDeleteAll)
	t.Run("SellerProductImages", testSellerProductImagesSliceDeleteAll)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesSliceDeleteAll)
	t.Run("SellerProductVariations", testSellerProductVariationsSliceDeleteAll)
	t.Run("SellerProducts", testSellerProductsSliceDeleteAll)
	t.Run("SellerRequests", testSellerRequestsSliceDeleteAll)
	t.Run("SellerShopProducts", testSellerShopProductsSliceDeleteAll)
	t.Run("SellerShops", testSellerShopsSliceDeleteAll)
	t.Run("ShopCategories", testShopCategoriesSliceDeleteAll)
	t.Run("UserLocations", testUserLocationsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Attributes", testAttributesExists)
	t.Run("Brands", testBrandsExists)
	t.Run("CartProducts", testCartProductsExists)
	t.Run("Carts", testCartsExists)
	t.Run("Categories", testCategoriesExists)
	t.Run("CheckoutProducts", testCheckoutProductsExists)
	t.Run("Checkouts", testCheckoutsExists)
	t.Run("SellerProductCategories", testSellerProductCategoriesExists)
	t.Run("SellerProductImages", testSellerProductImagesExists)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesExists)
	t.Run("SellerProductVariations", testSellerProductVariationsExists)
	t.Run("SellerProducts", testSellerProductsExists)
	t.Run("SellerRequests", testSellerRequestsExists)
	t.Run("SellerShopProducts", testSellerShopProductsExists)
	t.Run("SellerShops", testSellerShopsExists)
	t.Run("ShopCategories", testShopCategoriesExists)
	t.Run("UserLocations", testUserLocationsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Attributes", testAttributesFind)
	t.Run("Brands", testBrandsFind)
	t.Run("CartProducts", testCartProductsFind)
	t.Run("Carts", testCartsFind)
	t.Run("Categories", testCategoriesFind)
	t.Run("CheckoutProducts", testCheckoutProductsFind)
	t.Run("Checkouts", testCheckoutsFind)
	t.Run("SellerProductCategories", testSellerProductCategoriesFind)
	t.Run("SellerProductImages", testSellerProductImagesFind)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesFind)
	t.Run("SellerProductVariations", testSellerProductVariationsFind)
	t.Run("SellerProducts", testSellerProductsFind)
	t.Run("SellerRequests", testSellerRequestsFind)
	t.Run("SellerShopProducts", testSellerShopProductsFind)
	t.Run("SellerShops", testSellerShopsFind)
	t.Run("ShopCategories", testShopCategoriesFind)
	t.Run("UserLocations", testUserLocationsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Attributes", testAttributesBind)
	t.Run("Brands", testBrandsBind)
	t.Run("CartProducts", testCartProductsBind)
	t.Run("Carts", testCartsBind)
	t.Run("Categories", testCategoriesBind)
	t.Run("CheckoutProducts", testCheckoutProductsBind)
	t.Run("Checkouts", testCheckoutsBind)
	t.Run("SellerProductCategories", testSellerProductCategoriesBind)
	t.Run("SellerProductImages", testSellerProductImagesBind)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesBind)
	t.Run("SellerProductVariations", testSellerProductVariationsBind)
	t.Run("SellerProducts", testSellerProductsBind)
	t.Run("SellerRequests", testSellerRequestsBind)
	t.Run("SellerShopProducts", testSellerShopProductsBind)
	t.Run("SellerShops", testSellerShopsBind)
	t.Run("ShopCategories", testShopCategoriesBind)
	t.Run("UserLocations", testUserLocationsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Attributes", testAttributesOne)
	t.Run("Brands", testBrandsOne)
	t.Run("CartProducts", testCartProductsOne)
	t.Run("Carts", testCartsOne)
	t.Run("Categories", testCategoriesOne)
	t.Run("CheckoutProducts", testCheckoutProductsOne)
	t.Run("Checkouts", testCheckoutsOne)
	t.Run("SellerProductCategories", testSellerProductCategoriesOne)
	t.Run("SellerProductImages", testSellerProductImagesOne)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesOne)
	t.Run("SellerProductVariations", testSellerProductVariationsOne)
	t.Run("SellerProducts", testSellerProductsOne)
	t.Run("SellerRequests", testSellerRequestsOne)
	t.Run("SellerShopProducts", testSellerShopProductsOne)
	t.Run("SellerShops", testSellerShopsOne)
	t.Run("ShopCategories", testShopCategoriesOne)
	t.Run("UserLocations", testUserLocationsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Attributes", testAttributesAll)
	t.Run("Brands", testBrandsAll)
	t.Run("CartProducts", testCartProductsAll)
	t.Run("Carts", testCartsAll)
	t.Run("Categories", testCategoriesAll)
	t.Run("CheckoutProducts", testCheckoutProductsAll)
	t.Run("Checkouts", testCheckoutsAll)
	t.Run("SellerProductCategories", testSellerProductCategoriesAll)
	t.Run("SellerProductImages", testSellerProductImagesAll)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesAll)
	t.Run("SellerProductVariations", testSellerProductVariationsAll)
	t.Run("SellerProducts", testSellerProductsAll)
	t.Run("SellerRequests", testSellerRequestsAll)
	t.Run("SellerShopProducts", testSellerShopProductsAll)
	t.Run("SellerShops", testSellerShopsAll)
	t.Run("ShopCategories", testShopCategoriesAll)
	t.Run("UserLocations", testUserLocationsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Attributes", testAttributesCount)
	t.Run("Brands", testBrandsCount)
	t.Run("CartProducts", testCartProductsCount)
	t.Run("Carts", testCartsCount)
	t.Run("Categories", testCategoriesCount)
	t.Run("CheckoutProducts", testCheckoutProductsCount)
	t.Run("Checkouts", testCheckoutsCount)
	t.Run("SellerProductCategories", testSellerProductCategoriesCount)
	t.Run("SellerProductImages", testSellerProductImagesCount)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesCount)
	t.Run("SellerProductVariations", testSellerProductVariationsCount)
	t.Run("SellerProducts", testSellerProductsCount)
	t.Run("SellerRequests", testSellerRequestsCount)
	t.Run("SellerShopProducts", testSellerShopProductsCount)
	t.Run("SellerShops", testSellerShopsCount)
	t.Run("ShopCategories", testShopCategoriesCount)
	t.Run("UserLocations", testUserLocationsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Attributes", testAttributesHooks)
	t.Run("Brands", testBrandsHooks)
	t.Run("CartProducts", testCartProductsHooks)
	t.Run("Carts", testCartsHooks)
	t.Run("Categories", testCategoriesHooks)
	t.Run("CheckoutProducts", testCheckoutProductsHooks)
	t.Run("Checkouts", testCheckoutsHooks)
	t.Run("SellerProductCategories", testSellerProductCategoriesHooks)
	t.Run("SellerProductImages", testSellerProductImagesHooks)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesHooks)
	t.Run("SellerProductVariations", testSellerProductVariationsHooks)
	t.Run("SellerProducts", testSellerProductsHooks)
	t.Run("SellerRequests", testSellerRequestsHooks)
	t.Run("SellerShopProducts", testSellerShopProductsHooks)
	t.Run("SellerShops", testSellerShopsHooks)
	t.Run("ShopCategories", testShopCategoriesHooks)
	t.Run("UserLocations", testUserLocationsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Attributes", testAttributesInsert)
	t.Run("Attributes", testAttributesInsertWhitelist)
	t.Run("Brands", testBrandsInsert)
	t.Run("Brands", testBrandsInsertWhitelist)
	t.Run("CartProducts", testCartProductsInsert)
	t.Run("CartProducts", testCartProductsInsertWhitelist)
	t.Run("Carts", testCartsInsert)
	t.Run("Carts", testCartsInsertWhitelist)
	t.Run("Categories", testCategoriesInsert)
	t.Run("Categories", testCategoriesInsertWhitelist)
	t.Run("CheckoutProducts", testCheckoutProductsInsert)
	t.Run("CheckoutProducts", testCheckoutProductsInsertWhitelist)
	t.Run("Checkouts", testCheckoutsInsert)
	t.Run("Checkouts", testCheckoutsInsertWhitelist)
	t.Run("SellerProductCategories", testSellerProductCategoriesInsert)
	t.Run("SellerProductCategories", testSellerProductCategoriesInsertWhitelist)
	t.Run("SellerProductImages", testSellerProductImagesInsert)
	t.Run("SellerProductImages", testSellerProductImagesInsertWhitelist)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesInsert)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesInsertWhitelist)
	t.Run("SellerProductVariations", testSellerProductVariationsInsert)
	t.Run("SellerProductVariations", testSellerProductVariationsInsertWhitelist)
	t.Run("SellerProducts", testSellerProductsInsert)
	t.Run("SellerProducts", testSellerProductsInsertWhitelist)
	t.Run("SellerRequests", testSellerRequestsInsert)
	t.Run("SellerRequests", testSellerRequestsInsertWhitelist)
	t.Run("SellerShopProducts", testSellerShopProductsInsert)
	t.Run("SellerShopProducts", testSellerShopProductsInsertWhitelist)
	t.Run("SellerShops", testSellerShopsInsert)
	t.Run("SellerShops", testSellerShopsInsertWhitelist)
	t.Run("ShopCategories", testShopCategoriesInsert)
	t.Run("ShopCategories", testShopCategoriesInsertWhitelist)
	t.Run("UserLocations", testUserLocationsInsert)
	t.Run("UserLocations", testUserLocationsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CartProductToSellerProductVariationUsingSellerProductVariation", testCartProductToOneSellerProductVariationUsingSellerProductVariation)
	t.Run("CartProductToCartUsingCart", testCartProductToOneCartUsingCart)
	t.Run("CartProductToSellerProductUsingSellerProduct", testCartProductToOneSellerProductUsingSellerProduct)
	t.Run("CartToUserUsingUser", testCartToOneUserUsingUser)
	t.Run("CategoryToCategoryUsingParent", testCategoryToOneCategoryUsingParent)
	t.Run("CategoryToShopCategoryUsingShopCategory", testCategoryToOneShopCategoryUsingShopCategory)
	t.Run("CheckoutProductToSellerProductVariationUsingSellerProductVariation", testCheckoutProductToOneSellerProductVariationUsingSellerProductVariation)
	t.Run("CheckoutProductToCheckoutUsingCheckout", testCheckoutProductToOneCheckoutUsingCheckout)
	t.Run("CheckoutProductToSellerProductUsingSellerProduct", testCheckoutProductToOneSellerProductUsingSellerProduct)
	t.Run("CheckoutProductToUserUsingUser", testCheckoutProductToOneUserUsingUser)
	t.Run("CheckoutProductToUserUsingSellingSeller", testCheckoutProductToOneUserUsingSellingSeller)
	t.Run("CheckoutToCartUsingCart", testCheckoutToOneCartUsingCart)
	t.Run("CheckoutToUserLocationUsingUserLocation", testCheckoutToOneUserLocationUsingUserLocation)
	t.Run("CheckoutToUserUsingUser", testCheckoutToOneUserUsingUser)
	t.Run("SellerProductCategoryToSellerProductUsingSellerProduct", testSellerProductCategoryToOneSellerProductUsingSellerProduct)
	t.Run("SellerProductImageToSellerProductUsingSellerProduct", testSellerProductImageToOneSellerProductUsingSellerProduct)
	t.Run("SellerProductVariationValueToAttributeUsingAttribute", testSellerProductVariationValueToOneAttributeUsingAttribute)
	t.Run("SellerProductVariationValueToSellerProductVariationUsingSellerProductVariation", testSellerProductVariationValueToOneSellerProductVariationUsingSellerProductVariation)
	t.Run("SellerProductVariationToSellerProductUsingSellerProduct", testSellerProductVariationToOneSellerProductUsingSellerProduct)
	t.Run("SellerProductToBrandUsingBrand", testSellerProductToOneBrandUsingBrand)
	t.Run("SellerProductToSellerShopUsingSellerShop", testSellerProductToOneSellerShopUsingSellerShop)
	t.Run("SellerProductToUserUsingUser", testSellerProductToOneUserUsingUser)
	t.Run("SellerRequestToShopCategoryUsingShopCategory", testSellerRequestToOneShopCategoryUsingShopCategory)
	t.Run("SellerRequestToUserUsingUser", testSellerRequestToOneUserUsingUser)
	t.Run("SellerShopProductToSellerProductUsingSellerProduct", testSellerShopProductToOneSellerProductUsingSellerProduct)
	t.Run("SellerShopProductToSellerShopUsingSellerShop", testSellerShopProductToOneSellerShopUsingSellerShop)
	t.Run("SellerShopToShopCategoryUsingShopCategory", testSellerShopToOneShopCategoryUsingShopCategory)
	t.Run("SellerShopToUserUsingAdmin", testSellerShopToOneUserUsingAdmin)
	t.Run("SellerShopToUserUsingUser", testSellerShopToOneUserUsingUser)
	t.Run("UserLocationToUserUsingUser", testUserLocationToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AttributeToSellerProductVariationValues", testAttributeToManySellerProductVariationValues)
	t.Run("BrandToSellerProducts", testBrandToManySellerProducts)
	t.Run("CartToCartProducts", testCartToManyCartProducts)
	t.Run("CartToCheckouts", testCartToManyCheckouts)
	t.Run("CategoryToParentCategories", testCategoryToManyParentCategories)
	t.Run("CheckoutToCheckoutProducts", testCheckoutToManyCheckoutProducts)
	t.Run("SellerProductVariationToCartProducts", testSellerProductVariationToManyCartProducts)
	t.Run("SellerProductVariationToCheckoutProducts", testSellerProductVariationToManyCheckoutProducts)
	t.Run("SellerProductVariationToSellerProductVariationValues", testSellerProductVariationToManySellerProductVariationValues)
	t.Run("SellerProductToCartProducts", testSellerProductToManyCartProducts)
	t.Run("SellerProductToCheckoutProducts", testSellerProductToManyCheckoutProducts)
	t.Run("SellerProductToSellerProductCategories", testSellerProductToManySellerProductCategories)
	t.Run("SellerProductToSellerProductImages", testSellerProductToManySellerProductImages)
	t.Run("SellerProductToSellerProductVariations", testSellerProductToManySellerProductVariations)
	t.Run("SellerProductToSellerShopProducts", testSellerProductToManySellerShopProducts)
	t.Run("SellerShopToSellerProducts", testSellerShopToManySellerProducts)
	t.Run("SellerShopToSellerShopProducts", testSellerShopToManySellerShopProducts)
	t.Run("ShopCategoryToCategories", testShopCategoryToManyCategories)
	t.Run("ShopCategoryToSellerRequests", testShopCategoryToManySellerRequests)
	t.Run("ShopCategoryToSellerShops", testShopCategoryToManySellerShops)
	t.Run("UserLocationToCheckouts", testUserLocationToManyCheckouts)
	t.Run("UserToCarts", testUserToManyCarts)
	t.Run("UserToCheckoutProducts", testUserToManyCheckoutProducts)
	t.Run("UserToSellingSellerCheckoutProducts", testUserToManySellingSellerCheckoutProducts)
	t.Run("UserToCheckouts", testUserToManyCheckouts)
	t.Run("UserToSellerProducts", testUserToManySellerProducts)
	t.Run("UserToSellerRequests", testUserToManySellerRequests)
	t.Run("UserToAdminSellerShops", testUserToManyAdminSellerShops)
	t.Run("UserToSellerShops", testUserToManySellerShops)
	t.Run("UserToUserLocations", testUserToManyUserLocations)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CartProductToSellerProductVariationUsingCartProducts", testCartProductToOneSetOpSellerProductVariationUsingSellerProductVariation)
	t.Run("CartProductToCartUsingCartProducts", testCartProductToOneSetOpCartUsingCart)
	t.Run("CartProductToSellerProductUsingCartProducts", testCartProductToOneSetOpSellerProductUsingSellerProduct)
	t.Run("CartToUserUsingCarts", testCartToOneSetOpUserUsingUser)
	t.Run("CategoryToCategoryUsingParentCategories", testCategoryToOneSetOpCategoryUsingParent)
	t.Run("CategoryToShopCategoryUsingCategories", testCategoryToOneSetOpShopCategoryUsingShopCategory)
	t.Run("CheckoutProductToSellerProductVariationUsingCheckoutProducts", testCheckoutProductToOneSetOpSellerProductVariationUsingSellerProductVariation)
	t.Run("CheckoutProductToCheckoutUsingCheckoutProducts", testCheckoutProductToOneSetOpCheckoutUsingCheckout)
	t.Run("CheckoutProductToSellerProductUsingCheckoutProducts", testCheckoutProductToOneSetOpSellerProductUsingSellerProduct)
	t.Run("CheckoutProductToUserUsingCheckoutProducts", testCheckoutProductToOneSetOpUserUsingUser)
	t.Run("CheckoutProductToUserUsingSellingSellerCheckoutProducts", testCheckoutProductToOneSetOpUserUsingSellingSeller)
	t.Run("CheckoutToCartUsingCheckouts", testCheckoutToOneSetOpCartUsingCart)
	t.Run("CheckoutToUserLocationUsingCheckouts", testCheckoutToOneSetOpUserLocationUsingUserLocation)
	t.Run("CheckoutToUserUsingCheckouts", testCheckoutToOneSetOpUserUsingUser)
	t.Run("SellerProductCategoryToSellerProductUsingSellerProductCategories", testSellerProductCategoryToOneSetOpSellerProductUsingSellerProduct)
	t.Run("SellerProductImageToSellerProductUsingSellerProductImages", testSellerProductImageToOneSetOpSellerProductUsingSellerProduct)
	t.Run("SellerProductVariationValueToAttributeUsingSellerProductVariationValues", testSellerProductVariationValueToOneSetOpAttributeUsingAttribute)
	t.Run("SellerProductVariationValueToSellerProductVariationUsingSellerProductVariationValues", testSellerProductVariationValueToOneSetOpSellerProductVariationUsingSellerProductVariation)
	t.Run("SellerProductVariationToSellerProductUsingSellerProductVariations", testSellerProductVariationToOneSetOpSellerProductUsingSellerProduct)
	t.Run("SellerProductToBrandUsingSellerProducts", testSellerProductToOneSetOpBrandUsingBrand)
	t.Run("SellerProductToSellerShopUsingSellerProducts", testSellerProductToOneSetOpSellerShopUsingSellerShop)
	t.Run("SellerProductToUserUsingSellerProducts", testSellerProductToOneSetOpUserUsingUser)
	t.Run("SellerRequestToShopCategoryUsingSellerRequests", testSellerRequestToOneSetOpShopCategoryUsingShopCategory)
	t.Run("SellerRequestToUserUsingSellerRequests", testSellerRequestToOneSetOpUserUsingUser)
	t.Run("SellerShopProductToSellerProductUsingSellerShopProducts", testSellerShopProductToOneSetOpSellerProductUsingSellerProduct)
	t.Run("SellerShopProductToSellerShopUsingSellerShopProducts", testSellerShopProductToOneSetOpSellerShopUsingSellerShop)
	t.Run("SellerShopToShopCategoryUsingSellerShops", testSellerShopToOneSetOpShopCategoryUsingShopCategory)
	t.Run("SellerShopToUserUsingAdminSellerShops", testSellerShopToOneSetOpUserUsingAdmin)
	t.Run("SellerShopToUserUsingSellerShops", testSellerShopToOneSetOpUserUsingUser)
	t.Run("UserLocationToUserUsingUserLocations", testUserLocationToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("CartProductToSellerProductVariationUsingCartProducts", testCartProductToOneRemoveOpSellerProductVariationUsingSellerProductVariation)
	t.Run("CartProductToCartUsingCartProducts", testCartProductToOneRemoveOpCartUsingCart)
	t.Run("CartProductToSellerProductUsingCartProducts", testCartProductToOneRemoveOpSellerProductUsingSellerProduct)
	t.Run("CartToUserUsingCarts", testCartToOneRemoveOpUserUsingUser)
	t.Run("CategoryToCategoryUsingParentCategories", testCategoryToOneRemoveOpCategoryUsingParent)
	t.Run("CheckoutProductToSellerProductVariationUsingCheckoutProducts", testCheckoutProductToOneRemoveOpSellerProductVariationUsingSellerProductVariation)
	t.Run("CheckoutProductToCheckoutUsingCheckoutProducts", testCheckoutProductToOneRemoveOpCheckoutUsingCheckout)
	t.Run("CheckoutProductToSellerProductUsingCheckoutProducts", testCheckoutProductToOneRemoveOpSellerProductUsingSellerProduct)
	t.Run("CheckoutProductToUserUsingCheckoutProducts", testCheckoutProductToOneRemoveOpUserUsingUser)
	t.Run("CheckoutProductToUserUsingSellingSellerCheckoutProducts", testCheckoutProductToOneRemoveOpUserUsingSellingSeller)
	t.Run("CheckoutToCartUsingCheckouts", testCheckoutToOneRemoveOpCartUsingCart)
	t.Run("CheckoutToUserLocationUsingCheckouts", testCheckoutToOneRemoveOpUserLocationUsingUserLocation)
	t.Run("CheckoutToUserUsingCheckouts", testCheckoutToOneRemoveOpUserUsingUser)
	t.Run("SellerProductCategoryToSellerProductUsingSellerProductCategories", testSellerProductCategoryToOneRemoveOpSellerProductUsingSellerProduct)
	t.Run("SellerProductImageToSellerProductUsingSellerProductImages", testSellerProductImageToOneRemoveOpSellerProductUsingSellerProduct)
	t.Run("SellerProductToBrandUsingSellerProducts", testSellerProductToOneRemoveOpBrandUsingBrand)
	t.Run("SellerProductToSellerShopUsingSellerProducts", testSellerProductToOneRemoveOpSellerShopUsingSellerShop)
	t.Run("SellerProductToUserUsingSellerProducts", testSellerProductToOneRemoveOpUserUsingUser)
	t.Run("SellerRequestToUserUsingSellerRequests", testSellerRequestToOneRemoveOpUserUsingUser)
	t.Run("SellerShopProductToSellerProductUsingSellerShopProducts", testSellerShopProductToOneRemoveOpSellerProductUsingSellerProduct)
	t.Run("SellerShopProductToSellerShopUsingSellerShopProducts", testSellerShopProductToOneRemoveOpSellerShopUsingSellerShop)
	t.Run("SellerShopToShopCategoryUsingSellerShops", testSellerShopToOneRemoveOpShopCategoryUsingShopCategory)
	t.Run("SellerShopToUserUsingAdminSellerShops", testSellerShopToOneRemoveOpUserUsingAdmin)
	t.Run("SellerShopToUserUsingSellerShops", testSellerShopToOneRemoveOpUserUsingUser)
	t.Run("UserLocationToUserUsingUserLocations", testUserLocationToOneRemoveOpUserUsingUser)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AttributeToSellerProductVariationValues", testAttributeToManyAddOpSellerProductVariationValues)
	t.Run("BrandToSellerProducts", testBrandToManyAddOpSellerProducts)
	t.Run("CartToCartProducts", testCartToManyAddOpCartProducts)
	t.Run("CartToCheckouts", testCartToManyAddOpCheckouts)
	t.Run("CategoryToParentCategories", testCategoryToManyAddOpParentCategories)
	t.Run("CheckoutToCheckoutProducts", testCheckoutToManyAddOpCheckoutProducts)
	t.Run("SellerProductVariationToCartProducts", testSellerProductVariationToManyAddOpCartProducts)
	t.Run("SellerProductVariationToCheckoutProducts", testSellerProductVariationToManyAddOpCheckoutProducts)
	t.Run("SellerProductVariationToSellerProductVariationValues", testSellerProductVariationToManyAddOpSellerProductVariationValues)
	t.Run("SellerProductToCartProducts", testSellerProductToManyAddOpCartProducts)
	t.Run("SellerProductToCheckoutProducts", testSellerProductToManyAddOpCheckoutProducts)
	t.Run("SellerProductToSellerProductCategories", testSellerProductToManyAddOpSellerProductCategories)
	t.Run("SellerProductToSellerProductImages", testSellerProductToManyAddOpSellerProductImages)
	t.Run("SellerProductToSellerProductVariations", testSellerProductToManyAddOpSellerProductVariations)
	t.Run("SellerProductToSellerShopProducts", testSellerProductToManyAddOpSellerShopProducts)
	t.Run("SellerShopToSellerProducts", testSellerShopToManyAddOpSellerProducts)
	t.Run("SellerShopToSellerShopProducts", testSellerShopToManyAddOpSellerShopProducts)
	t.Run("ShopCategoryToCategories", testShopCategoryToManyAddOpCategories)
	t.Run("ShopCategoryToSellerRequests", testShopCategoryToManyAddOpSellerRequests)
	t.Run("ShopCategoryToSellerShops", testShopCategoryToManyAddOpSellerShops)
	t.Run("UserLocationToCheckouts", testUserLocationToManyAddOpCheckouts)
	t.Run("UserToCarts", testUserToManyAddOpCarts)
	t.Run("UserToCheckoutProducts", testUserToManyAddOpCheckoutProducts)
	t.Run("UserToSellingSellerCheckoutProducts", testUserToManyAddOpSellingSellerCheckoutProducts)
	t.Run("UserToCheckouts", testUserToManyAddOpCheckouts)
	t.Run("UserToSellerProducts", testUserToManyAddOpSellerProducts)
	t.Run("UserToSellerRequests", testUserToManyAddOpSellerRequests)
	t.Run("UserToAdminSellerShops", testUserToManyAddOpAdminSellerShops)
	t.Run("UserToSellerShops", testUserToManyAddOpSellerShops)
	t.Run("UserToUserLocations", testUserToManyAddOpUserLocations)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("BrandToSellerProducts", testBrandToManySetOpSellerProducts)
	t.Run("CartToCartProducts", testCartToManySetOpCartProducts)
	t.Run("CartToCheckouts", testCartToManySetOpCheckouts)
	t.Run("CategoryToParentCategories", testCategoryToManySetOpParentCategories)
	t.Run("CheckoutToCheckoutProducts", testCheckoutToManySetOpCheckoutProducts)
	t.Run("SellerProductVariationToCartProducts", testSellerProductVariationToManySetOpCartProducts)
	t.Run("SellerProductVariationToCheckoutProducts", testSellerProductVariationToManySetOpCheckoutProducts)
	t.Run("SellerProductToCartProducts", testSellerProductToManySetOpCartProducts)
	t.Run("SellerProductToCheckoutProducts", testSellerProductToManySetOpCheckoutProducts)
	t.Run("SellerProductToSellerProductCategories", testSellerProductToManySetOpSellerProductCategories)
	t.Run("SellerProductToSellerProductImages", testSellerProductToManySetOpSellerProductImages)
	t.Run("SellerProductToSellerShopProducts", testSellerProductToManySetOpSellerShopProducts)
	t.Run("SellerShopToSellerProducts", testSellerShopToManySetOpSellerProducts)
	t.Run("SellerShopToSellerShopProducts", testSellerShopToManySetOpSellerShopProducts)
	t.Run("ShopCategoryToSellerShops", testShopCategoryToManySetOpSellerShops)
	t.Run("UserLocationToCheckouts", testUserLocationToManySetOpCheckouts)
	t.Run("UserToCarts", testUserToManySetOpCarts)
	t.Run("UserToCheckoutProducts", testUserToManySetOpCheckoutProducts)
	t.Run("UserToSellingSellerCheckoutProducts", testUserToManySetOpSellingSellerCheckoutProducts)
	t.Run("UserToCheckouts", testUserToManySetOpCheckouts)
	t.Run("UserToSellerProducts", testUserToManySetOpSellerProducts)
	t.Run("UserToSellerRequests", testUserToManySetOpSellerRequests)
	t.Run("UserToAdminSellerShops", testUserToManySetOpAdminSellerShops)
	t.Run("UserToSellerShops", testUserToManySetOpSellerShops)
	t.Run("UserToUserLocations", testUserToManySetOpUserLocations)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("BrandToSellerProducts", testBrandToManyRemoveOpSellerProducts)
	t.Run("CartToCartProducts", testCartToManyRemoveOpCartProducts)
	t.Run("CartToCheckouts", testCartToManyRemoveOpCheckouts)
	t.Run("CategoryToParentCategories", testCategoryToManyRemoveOpParentCategories)
	t.Run("CheckoutToCheckoutProducts", testCheckoutToManyRemoveOpCheckoutProducts)
	t.Run("SellerProductVariationToCartProducts", testSellerProductVariationToManyRemoveOpCartProducts)
	t.Run("SellerProductVariationToCheckoutProducts", testSellerProductVariationToManyRemoveOpCheckoutProducts)
	t.Run("SellerProductToCartProducts", testSellerProductToManyRemoveOpCartProducts)
	t.Run("SellerProductToCheckoutProducts", testSellerProductToManyRemoveOpCheckoutProducts)
	t.Run("SellerProductToSellerProductCategories", testSellerProductToManyRemoveOpSellerProductCategories)
	t.Run("SellerProductToSellerProductImages", testSellerProductToManyRemoveOpSellerProductImages)
	t.Run("SellerProductToSellerShopProducts", testSellerProductToManyRemoveOpSellerShopProducts)
	t.Run("SellerShopToSellerProducts", testSellerShopToManyRemoveOpSellerProducts)
	t.Run("SellerShopToSellerShopProducts", testSellerShopToManyRemoveOpSellerShopProducts)
	t.Run("ShopCategoryToSellerShops", testShopCategoryToManyRemoveOpSellerShops)
	t.Run("UserLocationToCheckouts", testUserLocationToManyRemoveOpCheckouts)
	t.Run("UserToCarts", testUserToManyRemoveOpCarts)
	t.Run("UserToCheckoutProducts", testUserToManyRemoveOpCheckoutProducts)
	t.Run("UserToSellingSellerCheckoutProducts", testUserToManyRemoveOpSellingSellerCheckoutProducts)
	t.Run("UserToCheckouts", testUserToManyRemoveOpCheckouts)
	t.Run("UserToSellerProducts", testUserToManyRemoveOpSellerProducts)
	t.Run("UserToSellerRequests", testUserToManyRemoveOpSellerRequests)
	t.Run("UserToAdminSellerShops", testUserToManyRemoveOpAdminSellerShops)
	t.Run("UserToSellerShops", testUserToManyRemoveOpSellerShops)
	t.Run("UserToUserLocations", testUserToManyRemoveOpUserLocations)
}

func TestReload(t *testing.T) {
	t.Run("Attributes", testAttributesReload)
	t.Run("Brands", testBrandsReload)
	t.Run("CartProducts", testCartProductsReload)
	t.Run("Carts", testCartsReload)
	t.Run("Categories", testCategoriesReload)
	t.Run("CheckoutProducts", testCheckoutProductsReload)
	t.Run("Checkouts", testCheckoutsReload)
	t.Run("SellerProductCategories", testSellerProductCategoriesReload)
	t.Run("SellerProductImages", testSellerProductImagesReload)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesReload)
	t.Run("SellerProductVariations", testSellerProductVariationsReload)
	t.Run("SellerProducts", testSellerProductsReload)
	t.Run("SellerRequests", testSellerRequestsReload)
	t.Run("SellerShopProducts", testSellerShopProductsReload)
	t.Run("SellerShops", testSellerShopsReload)
	t.Run("ShopCategories", testShopCategoriesReload)
	t.Run("UserLocations", testUserLocationsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Attributes", testAttributesReloadAll)
	t.Run("Brands", testBrandsReloadAll)
	t.Run("CartProducts", testCartProductsReloadAll)
	t.Run("Carts", testCartsReloadAll)
	t.Run("Categories", testCategoriesReloadAll)
	t.Run("CheckoutProducts", testCheckoutProductsReloadAll)
	t.Run("Checkouts", testCheckoutsReloadAll)
	t.Run("SellerProductCategories", testSellerProductCategoriesReloadAll)
	t.Run("SellerProductImages", testSellerProductImagesReloadAll)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesReloadAll)
	t.Run("SellerProductVariations", testSellerProductVariationsReloadAll)
	t.Run("SellerProducts", testSellerProductsReloadAll)
	t.Run("SellerRequests", testSellerRequestsReloadAll)
	t.Run("SellerShopProducts", testSellerShopProductsReloadAll)
	t.Run("SellerShops", testSellerShopsReloadAll)
	t.Run("ShopCategories", testShopCategoriesReloadAll)
	t.Run("UserLocations", testUserLocationsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Attributes", testAttributesSelect)
	t.Run("Brands", testBrandsSelect)
	t.Run("CartProducts", testCartProductsSelect)
	t.Run("Carts", testCartsSelect)
	t.Run("Categories", testCategoriesSelect)
	t.Run("CheckoutProducts", testCheckoutProductsSelect)
	t.Run("Checkouts", testCheckoutsSelect)
	t.Run("SellerProductCategories", testSellerProductCategoriesSelect)
	t.Run("SellerProductImages", testSellerProductImagesSelect)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesSelect)
	t.Run("SellerProductVariations", testSellerProductVariationsSelect)
	t.Run("SellerProducts", testSellerProductsSelect)
	t.Run("SellerRequests", testSellerRequestsSelect)
	t.Run("SellerShopProducts", testSellerShopProductsSelect)
	t.Run("SellerShops", testSellerShopsSelect)
	t.Run("ShopCategories", testShopCategoriesSelect)
	t.Run("UserLocations", testUserLocationsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Attributes", testAttributesUpdate)
	t.Run("Brands", testBrandsUpdate)
	t.Run("CartProducts", testCartProductsUpdate)
	t.Run("Carts", testCartsUpdate)
	t.Run("Categories", testCategoriesUpdate)
	t.Run("CheckoutProducts", testCheckoutProductsUpdate)
	t.Run("Checkouts", testCheckoutsUpdate)
	t.Run("SellerProductCategories", testSellerProductCategoriesUpdate)
	t.Run("SellerProductImages", testSellerProductImagesUpdate)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesUpdate)
	t.Run("SellerProductVariations", testSellerProductVariationsUpdate)
	t.Run("SellerProducts", testSellerProductsUpdate)
	t.Run("SellerRequests", testSellerRequestsUpdate)
	t.Run("SellerShopProducts", testSellerShopProductsUpdate)
	t.Run("SellerShops", testSellerShopsUpdate)
	t.Run("ShopCategories", testShopCategoriesUpdate)
	t.Run("UserLocations", testUserLocationsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Attributes", testAttributesSliceUpdateAll)
	t.Run("Brands", testBrandsSliceUpdateAll)
	t.Run("CartProducts", testCartProductsSliceUpdateAll)
	t.Run("Carts", testCartsSliceUpdateAll)
	t.Run("Categories", testCategoriesSliceUpdateAll)
	t.Run("CheckoutProducts", testCheckoutProductsSliceUpdateAll)
	t.Run("Checkouts", testCheckoutsSliceUpdateAll)
	t.Run("SellerProductCategories", testSellerProductCategoriesSliceUpdateAll)
	t.Run("SellerProductImages", testSellerProductImagesSliceUpdateAll)
	t.Run("SellerProductVariationValues", testSellerProductVariationValuesSliceUpdateAll)
	t.Run("SellerProductVariations", testSellerProductVariationsSliceUpdateAll)
	t.Run("SellerProducts", testSellerProductsSliceUpdateAll)
	t.Run("SellerRequests", testSellerRequestsSliceUpdateAll)
	t.Run("SellerShopProducts", testSellerShopProductsSliceUpdateAll)
	t.Run("SellerShops", testSellerShopsSliceUpdateAll)
	t.Run("ShopCategories", testShopCategoriesSliceUpdateAll)
	t.Run("UserLocations", testUserLocationsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}