package constants

import (
	"github.com/jmoiron/sqlx"
)

// Constants for database table and column names
const (
	PRODUCTS_TABLE = "products"
	PRODUCTID      = "productid"
	PRODUCT_NAME   = "name"
	PRODUCT_DESC   = "description"
	PRODUCT_PRICE  = "price"
	PRODUCT_STOCK  = "stockQuantity"

	CATEGORIES_TABLE = "categories"
	CATEGORYID       = "categoryID"
	CATEGORY_NAME    = "name"
	CATEGORY_DESC    = "description"

	BRANDS_TABLE = "brands"
	BRANDID      = "brandID"
	BRAND_NAME   = "name"
	BRAND_DESC   = "description"

	USERS_TABLE = "users"
	USERID      = "userid"
)

// SQL queries for products, categories, and brands

var QueryProducts = "SELECT * FROM " + PRODUCTS_TABLE                                     //fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM "+PRODUCTS_TABLE, PRODUCTID, PRODUCT_NAME, PRODUCT_DESC, PRODUCT_PRICE, PRODUCT_STOCK, CATEGORYID, BRANDID)
var QueryProduct = "SELECT * FROM " + PRODUCTS_TABLE + " WHERE " + PRODUCTID + " = ?"     //fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM "+PRODUCTS_TABLE+" WHERE %s = ?", PRODUCTID, PRODUCT_NAME, PRODUCT_DESC, PRODUCT_PRICE, PRODUCT_STOCK, CATEGORYID, BRANDID, PRODUCTID)
var QueryCategories = "SELECT * FROM " + CATEGORIES_TABLE                                 // fmt.Sprintf("SELECT %s, %s, %s FROM "+CATEGORIES_TABLE, PRODUCTID, CATEGORY_NAME, CATEGORY_DESC)
var QueryCategory = "SELECT * FROM " + CATEGORIES_TABLE + " WHERE " + CATEGORYID + " = ?" // fmt.Sprintf("SELECT %s, %s FROM "+CATEGORIES_TABLE+" WHERE %s = ?", CATEGORY_NAME, CATEGORY_DESC, PRODUCTID)
var QueryBrands = "SELECT * FROM " + BRANDS_TABLE                                         //fmt.Sprintf("SELECT %s, %s, %s FROM "+BRANDS_TABLE, PRODUCTID, BRAND_NAME, BRAND_DESC)
var QueryBrand = "SELECT * FROM " + BRANDS_TABLE + " WHERE " + BRANDID + " = ?"           //fmt.Sprintf("SELECT %s, %s FROM "+BRANDS_TABLE+" WHERE %s = ?", BRAND_NAME, BRAND_DESC, PRODUCTID)

var QueryUser = "SELECT * FROM " + USERS_TABLE + " WHERE " + USERID + " = ?"

var QueryUserLogin = "SELECT * FROM " + USERS_TABLE + " WHERE email = ?"

var DB *sqlx.DB
