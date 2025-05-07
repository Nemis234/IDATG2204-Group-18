package constants

var QueryProducts = "SELECT * FROM " + PRODUCTS_TABLE                                 //fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM "+PRODUCTS_TABLE, PRODUCTID, PRODUCT_NAME, PRODUCT_DESC, PRODUCT_PRICE, PRODUCT_STOCK, CATEGORYID, BRANDID)
var QueryProduct = "SELECT * FROM " + PRODUCTS_TABLE + " WHERE " + PRODUCTID + " = ?" //fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM "+PRODUCTS_TABLE+" WHERE %s = ?", PRODUCTID, PRODUCT_NAME, PRODUCT_DESC, PRODUCT_PRICE, PRODUCT_STOCK, CATEGORYID, BRANDID, PRODUCTID)
var DeleteProduct = "DELETE FROM " + PRODUCTS_TABLE + " WHERE " + PRODUCTID + " = ?"
var InsertProduct = "INSERT INTO " + PRODUCTS_TABLE + " (" + PRODUCTID + ", " + PRODUCT_NAME + ", " + PRODUCT_DESC + ", " + PRODUCT_IMG_URL + ", " + PRODUCT_PRICE + ", " + PRODUCT_STOCK + ", " + PRODUCT_CATEGORY + ", " + PRODUCT_BRAND + ") VALUES (:" + PRODUCTID + ", :" + PRODUCT_NAME + ", :" + PRODUCT_DESC + ", :" + PRODUCT_IMG_URL + ", :" + PRODUCT_PRICE + ", :" + PRODUCT_STOCK + ", :" + PRODUCT_CATEGORY + ", :" + PRODUCT_BRAND + ")"

var QueryCategories = "SELECT * FROM " + CATEGORIES_TABLE                                    // fmt.Sprintf("SELECT %s, %s, %s FROM "+CATEGORIES_TABLE, PRODUCTID, CATEGORY_NAME, CATEGORY_DESC)
var QueryCategory = "SELECT * FROM " + CATEGORIES_TABLE + " WHERE " + CATEGORY_NAME + " = ?" // fmt.Sprintf("SELECT %s, %s FROM "+CATEGORIES_TABLE+" WHERE %s = ?", CATEGORY_NAME, CATEGORY_DESC, PRODUCTID)

var QueryBrands = "SELECT * FROM " + BRANDS_TABLE                                  //fmt.Sprintf("SELECT %s, %s, %s FROM "+BRANDS_TABLE, PRODUCTID, BRAND_NAME, BRAND_DESC)
var QueryBrand = "SELECT * FROM " + BRANDS_TABLE + " WHERE " + BRAND_NAME + " = ?" //fmt.Sprintf("SELECT %s, %s FROM "+BRANDS_TABLE+" WHERE %s = ?", BRAND_NAME, BRAND_DESC, PRODUCTID)

var QueryUser = "SELECT * FROM " + USERS_TABLE + " WHERE " + USERID + " = ?"
var QueryUserLogin = "SELECT * FROM " + USERS_TABLE + " WHERE email = ?"
