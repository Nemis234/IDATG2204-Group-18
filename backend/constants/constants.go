package constants

import (
	"github.com/jmoiron/sqlx"
)

// Constants for database table and column names
const (
	PRODUCTS_TABLE   = "Product"
	PRODUCT_ID       = "ProductID"
	PRODUCT_NAME     = "ProductName"
	PRODUCT_DESC     = "ProductDesc"
	PRODUCT_IMG_URL  = "ProductImgUrl"
	PRODUCT_PRICE    = "Price"
	PRODUCT_STOCK    = "StockQuantity"
	PRODUCT_BRAND    = "Brand"
	PRODUCT_CATEGORY = "Category"

	CATEGORIES_TABLE = "Category"
	CATEGORY_NAME    = "CategoryName"
	CATEGORY_DESC    = "CategoryDesc"

	BRANDS_TABLE = "Brand"
	BRAND_NAME   = "BrandName"
	BRAND_DESC   = "BrandDesc"

	USERS_TABLE = "Users"
	USER_ID     = "UserID"
)

var DB *sqlx.DB
