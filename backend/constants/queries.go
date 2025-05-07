package constants

var QueryProducts = "SELECT * FROM " + PRODUCTS_TABLE
var QueryProduct = "SELECT * FROM " + PRODUCTS_TABLE + " WHERE " + PRODUCTID + " = ?"
var DeleteProduct = "DELETE FROM " + PRODUCTS_TABLE + " WHERE " + PRODUCTID + " = ?"
var InsertProduct = "INSERT INTO " + PRODUCTS_TABLE + " (" + PRODUCTID + ", " + PRODUCT_NAME + ", " + PRODUCT_DESC + ", " + PRODUCT_IMG_URL + ", " + PRODUCT_PRICE + ", " + PRODUCT_STOCK + ", " + PRODUCT_CATEGORY + ", " + PRODUCT_BRAND + ") VALUES (:" + PRODUCTID + ", :" + PRODUCT_NAME + ", :" + PRODUCT_DESC + ", :" + PRODUCT_IMG_URL + ", :" + PRODUCT_PRICE + ", :" + PRODUCT_STOCK + ", :" + PRODUCT_CATEGORY + ", :" + PRODUCT_BRAND + ")"

var QueryCategories = "SELECT * FROM " + CATEGORIES_TABLE
var QueryCategory = "SELECT * FROM " + CATEGORIES_TABLE + " WHERE " + CATEGORY_NAME + " = ?"

var QueryBrands = "SELECT * FROM " + BRANDS_TABLE
var QueryBrand = "SELECT * FROM " + BRANDS_TABLE + " WHERE " + BRAND_NAME + " = ?"

var QueryUser = "SELECT * FROM " + USERS_TABLE + " WHERE " + USERID + " = ?"
var QueryUserLogin = "SELECT * FROM " + USERS_TABLE + " WHERE email = ?"
