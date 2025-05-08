package constants

var QueryProducts = "SELECT * FROM " + PRODUCTS_TABLE
var QueryProduct = "SELECT * FROM " + PRODUCTS_TABLE + " WHERE " + PRODUCT_ID + " = ?"
var DeleteProduct = "DELETE FROM " + PRODUCTS_TABLE + " WHERE " + PRODUCT_ID + " = ?"
var InsertProduct = "INSERT INTO " + PRODUCTS_TABLE + " (" + PRODUCT_ID + ", " + PRODUCT_NAME + ", " + PRODUCT_DESC + ", " + PRODUCT_IMG_URL + ", " + PRODUCT_PRICE + ", " + PRODUCT_STOCK + ", " + PRODUCT_CATEGORY + ", " + PRODUCT_BRAND + ") VALUES (:" + PRODUCT_ID + ", :" + PRODUCT_NAME + ", :" + PRODUCT_DESC + ", :" + PRODUCT_IMG_URL + ", :" + PRODUCT_PRICE + ", :" + PRODUCT_STOCK + ", :" + PRODUCT_CATEGORY + ", :" + PRODUCT_BRAND + ")"
var UpdateProduct = "UPDATE " + PRODUCTS_TABLE + " SET " + PRODUCT_NAME + " = :" + PRODUCT_NAME + ", " + PRODUCT_DESC + " = :" + PRODUCT_DESC + ", " + PRODUCT_IMG_URL + " = :" + PRODUCT_IMG_URL + ", " + PRODUCT_PRICE + " = :" + PRODUCT_PRICE + ", " + PRODUCT_STOCK + " = :" + PRODUCT_STOCK + ", " + PRODUCT_CATEGORY + " = :" + PRODUCT_CATEGORY + ", " + PRODUCT_BRAND + " = :" + PRODUCT_BRAND + " WHERE " + PRODUCT_ID + " = :" + PRODUCT_ID

var QueryCategories = "SELECT * FROM " + CATEGORIES_TABLE
var QueryCategory = "SELECT * FROM " + CATEGORIES_TABLE + " WHERE " + CATEGORY_NAME + " = ?"
var InsertCategory = "INSERT INTO " + CATEGORIES_TABLE + " (" + CATEGORY_NAME + ", " + CATEGORY_DESC + ") VALUES (:" + CATEGORY_NAME + ", :" + CATEGORY_DESC + ")"
var UpdateCategory = "UPDATE " + CATEGORIES_TABLE + " SET " + CATEGORY_DESC + " = :" + CATEGORY_DESC + " WHERE " + CATEGORY_NAME + " = :" + CATEGORY_NAME
var DeleteCategory = "DELETE FROM " + CATEGORIES_TABLE + " WHERE " + CATEGORY_NAME + " = ?"

var QueryBrands = "SELECT * FROM " + BRANDS_TABLE
var QueryBrand = "SELECT * FROM " + BRANDS_TABLE + " WHERE " + BRAND_NAME + " = ?"
var InsertBrand = "INSERT INTO " + BRANDS_TABLE + " (" + BRAND_NAME + ", " + BRAND_DESC + ") VALUES (:" + BRAND_NAME + ", :" + BRAND_DESC + ")"
var UpdateBrand = "UPDATE " + BRANDS_TABLE + " SET " + BRAND_DESC + " = :" + BRAND_DESC + " WHERE " + BRAND_NAME + " = :" + BRAND_NAME
var DeleteBrand = "DELETE FROM " + BRANDS_TABLE + " WHERE " + BRAND_NAME + " = ?"

var QueryUser = "SELECT * FROM " + USERS_TABLE + " WHERE " + USERID + " = ?"
var QueryUserLogin = "SELECT * FROM " + USERS_TABLE + " WHERE email = ?"
var QueryUser = "SELECT * FROM " + USERS_TABLE + " WHERE " + USER_ID + " = ?"
