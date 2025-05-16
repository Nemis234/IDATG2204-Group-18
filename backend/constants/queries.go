package constants

var QueryProducts = "SELECT * FROM " + PRODUCTS_TABLE
var QueryProduct = "SELECT * FROM " + PRODUCTS_TABLE + " WHERE " + PRODUCT_ID + " = ?"
var DeleteProduct = "DELETE FROM " + PRODUCTS_TABLE + " WHERE " + PRODUCT_ID + " = ?"
var InsertProduct = "INSERT INTO " + PRODUCTS_TABLE + " (" + PRODUCT_ID + ", " + PRODUCT_NAME + ", " + PRODUCT_DESC + ", " + PRODUCT_IMG_URL + ", " + PRODUCT_PRICE + ", " + PRODUCT_STOCK + ", " + PRODUCT_CATEGORY + ", " + PRODUCT_BRAND + ") VALUES (:" + PRODUCT_ID + ", :" + PRODUCT_NAME + ", :" + PRODUCT_DESC + ", :" + PRODUCT_IMG_URL + ", :" + PRODUCT_PRICE + ", :" + PRODUCT_STOCK + ", :" + PRODUCT_CATEGORY + ", :" + PRODUCT_BRAND + ")"
var UpdateProduct = "UPDATE " + PRODUCTS_TABLE + " SET " + PRODUCT_NAME + " = :" + PRODUCT_NAME + ", " + PRODUCT_DESC + " = :" + PRODUCT_DESC + ", " + PRODUCT_IMG_URL + " = :" + PRODUCT_IMG_URL + ", " + PRODUCT_PRICE + " = :" + PRODUCT_PRICE + ", " + PRODUCT_STOCK + " = :" + PRODUCT_STOCK + ", " + PRODUCT_CATEGORY + " = :" + PRODUCT_CATEGORY + ", " + PRODUCT_BRAND + " = :" + PRODUCT_BRAND + " WHERE " + PRODUCT_ID + " = :" + PRODUCT_ID

var QueryReviewsByProduct = "SELECT * FROM " + REVIEWS_TABLE + " WHERE " + PRODUCT_ID + " = ?"
var QueryReviewsByUser = "SELECT * FROM " + REVIEWS_TABLE + " WHERE " + USER_ID + " = ?"
var QueryReview = "SELECT * FROM " + REVIEWS_TABLE + " WHERE " + PRODUCT_ID + " = ? AND " + USER_ID + " = ?"
var InsertReview = "INSERT INTO " + REVIEWS_TABLE + " (" + PRODUCT_ID + ", " + USER_ID + ", " + REVIEW_COMMENT + ", " + REVIEW_RATING + ", " + REVIEW_POST_DATE + ") VALUES (:" + PRODUCT_ID + ", :" + USER_ID + ", :" + REVIEW_COMMENT + ", :" + REVIEW_RATING + ", :" + REVIEW_POST_DATE + ")"
var UpdateReview = "UPDATE " + REVIEWS_TABLE + " SET " + REVIEW_COMMENT + " = :" + REVIEW_COMMENT + ", " + REVIEW_RATING + " = :" + REVIEW_RATING + ", " + REVIEW_POST_DATE + " = :" + REVIEW_POST_DATE + " WHERE " + PRODUCT_ID + " = :" + PRODUCT_ID + " AND " + USER_ID + " = :" + USER_ID
var DeleteReview = "DELETE FROM " + REVIEWS_TABLE + " WHERE " + PRODUCT_ID + " = ? AND " + USER_ID + " = ?"

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

var QueryOrders = "SELECT * FROM " + ORDER_TABLE
var QueryOrder = "SELECT * FROM " + ORDER_TABLE + " WHERE " + ORDER_ID + " = ?"
var InsertOrder = "INSERT INTO " + ORDER_TABLE + " (" + ORDER_ID + ", " + USER_ID + ", " + ORDER_DATE + ", " + ORDER_STATUS + ", " + ORDER_TOTAL + ") VALUES (:" + ORDER_ID + ", :" + USER_ID + ", :" + ORDER_DATE + ", :" + ORDER_STATUS + ", :" + ORDER_TOTAL + ")"
var UpdateOrder = "UPDATE " + ORDER_TABLE + " SET " + ORDER_DATE + " = :" + ORDER_DATE + ", " + ORDER_STATUS + " = :" + ORDER_STATUS + ", " + ORDER_TOTAL + " = :" + ORDER_TOTAL + " WHERE " + ORDER_ID + " = :" + ORDER_ID
var DeleteOrder = "DELETE FROM " + ORDER_TABLE + " WHERE " + ORDER_ID + " = ?"
var GetUserIDByOrderID = "SELECT " + USER_ID + " FROM " + ORDER_TABLE + " WHERE " + ORDER_ID + " = ?"

var QueryOrderItemsByID = "SELECT * FROM " + ORDER_ITEMS_TABLE + " WHERE " + ORDER_ID + " = ?"
var QueryOrderItemByID = "SELECT * FROM " + ORDER_ITEMS_TABLE + " WHERE " + ORDER_ID + " = ? AND " + PRODUCT_ID + " = ?"
var InsertOrderItem = "INSERT INTO " + ORDER_ITEMS_TABLE + " (" + ORDER_ID + ", " + PRODUCT_ID + ", " + ORDER_QUANTITY + ") VALUES (:" + ORDER_ID + ", :" + PRODUCT_ID + ", :" + ORDER_QUANTITY + ")"
var UpdateOrderItem = "UPDATE " + ORDER_ITEMS_TABLE + " SET " + ORDER_QUANTITY + " = :" + ORDER_QUANTITY + " WHERE " + ORDER_ID + " = :" + ORDER_ID + " AND " + PRODUCT_ID + " = :" + PRODUCT_ID
var DeleteOrderItem = "DELETE FROM " + ORDER_ITEMS_TABLE + " WHERE " + ORDER_ID + " = ? AND " + PRODUCT_ID + " = ?"

var QueryOrderStatus = "SELECT * FROM " + ORDER_STATUS_TABLE
var InsertOrderStatus = "INSERT INTO " + ORDER_STATUS_TABLE + " (" + ORDER_STATUS_NAME + ", " + ORDER_STATUS_DESC + ") VALUES (:" + ORDER_STATUS_NAME + ", :" + ORDER_STATUS_DESC + ")"
var UpdateOrderStatus = "UPDATE " + ORDER_STATUS_TABLE + " SET " + ORDER_STATUS_DESC + " = :" + ORDER_STATUS_DESC + " WHERE " + ORDER_STATUS_NAME + " = :" + ORDER_STATUS_NAME
var DeleteOrderStatus = "DELETE FROM " + ORDER_STATUS_TABLE + " WHERE " + ORDER_STATUS_NAME + " = ?"

var QueryUserLoginByEmail =  "SELECT u."+USER_ID+", u."+USER_USERNAME+", u."+USER_PASSWORD+", u."+USER_EMAIL+", u."+USER_FIRSTNAME+", u."+USER_LASTNAME+", u."+USER_LASTNAME+", a."+ADMINSTRATORS_ROLENAME+" FROM "+ USERS_TABLE +" u LEFT JOIN " + ADMINSTRATORS_TABLE + " a ON u."+USER_ID+" = a."+USER_ID+" WHERE u."+USER_EMAIL+"= ?"
var QueryUserLoginByUsername = "SELECT u."+USER_ID+", u."+USER_USERNAME+", u."+USER_PASSWORD+", u."+USER_EMAIL+", u."+USER_FIRSTNAME+", u."+USER_LASTNAME+", u."+USER_LASTNAME+", a."+ADMINSTRATORS_ROLENAME+" FROM "+ USERS_TABLE +" u LEFT JOIN " + ADMINSTRATORS_TABLE + " a ON u."+USER_ID+" = a."+USER_ID+" WHERE u."+USER_USERNAME+"= ?"
var QueryUserCountByEmailOrUsername = "SELECT COUNT(*) FROM " + USERS_TABLE + " WHERE Email = ? OR Username = ?"
var InsertUser = "INSERT INTO " + USERS_TABLE + " (" + USER_ID + ", " + USER_USERNAME + ", " + USER_PASSWORD + ", " + USER_EMAIL + ", " + USER_FIRSTNAME + ", " + USER_LASTNAME + ", " + USER_ADDRESS + ") VALUES (:" + USER_ID + ", :" + USER_USERNAME + ", :" + USER_PASSWORD + ", :" + USER_EMAIL + ", :" + USER_FIRSTNAME + ", :" + USER_LASTNAME + ", :" + USER_ADDRESS +")"
var DeleteUser = "DELETE FROM "+USERS_TABLE+" WHERE "+USER_ID+" = ?"
var QueryUser = "SELECT * FROM " + USERS_TABLE + " WHERE " + USER_ID + " = ?"
var QueryUsers = "SELECT * FROM " + USERS_TABLE

var QueryAdministrator = "SELECT * FROM " + ADMINSTRATORS_TABLE
var InsertAdministrator = "INSERT INTO "+ ADMINSTRATORS_TABLE +" ("+USER_ID+", "+ADMINSTRATORS_ROLENAME+") VALUES (:"+USER_ID+", :"+ADMINSTRATORS_ROLENAME+")"