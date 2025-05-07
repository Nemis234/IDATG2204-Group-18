package constants

import (
	"github.com/jmoiron/sqlx"
)

// Constants for database table and column names
const (
	PRODUCTS_TABLE   = "Product"
	PRODUCTID        = "ProductID"
	PRODUCT_NAME     = "ProductName"
	PRODUCT_DESC     = "ProductDesc"
	PRODUCT_IMG_URL  = "ProductImgUrl"
	PRODUCT_PRICE    = "Price"
	PRODUCT_STOCK    = "StockQuantity"
	PRODUCT_BRAND    = "Brand"
	PRODUCT_CATEGORY = "Category"

	CATEGORIES_TABLE = "Category"
	CATEGORYID       = "categoryID"
	CATEGORY_NAME    = "CategoryName"
	CATEGORY_DESC    = "CategoryDesc"

	BRANDS_TABLE = "Brand"
	BRANDID      = "brandID"
	BRAND_NAME   = "BrandName"
	BRAND_DESC   = "BrandDesc"

	USERS_TABLE = "Users"
	USERID      = "UserID"
)

var DB *sqlx.DB
