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

	REVIEWS_TABLE    = "Review"
	REVIEW_COMMENT   = "Comment"
	REVIEW_RATING    = "Rating"
	REVIEW_POST_DATE = "PostDate"

	CATEGORIES_TABLE = "Category"
	CATEGORY_NAME    = "CategoryName"
	CATEGORY_DESC    = "CategoryDesc"

	BRANDS_TABLE = "Brand"
	BRAND_NAME   = "BrandName"
	BRAND_DESC   = "BrandDesc"

	ORDER_TABLE  = "OrderTable"
	ORDER_ID     = "OrderID"
	ORDER_DATE   = "OrderDate"
	ORDER_STATUS = "OrderStatus"
	ORDER_TOTAL  = "OrderTotal"

	ORDER_ITEMS_TABLE = "OrderItem"
	ORDER_QUANTITY    = "Quantity"

	ORDER_STATUS_TABLE = "OrderStatus"
	ORDER_STATUS_NAME  = "StatusName"
	ORDER_STATUS_DESC  = "StatusDesc"

	PAYMENT_TABLE  = "Payment"
	PAYMENT_ID     = "PaymentID"
	PAYMENT_METHOD = "PaymentMethod"
	PAYMENT_AMOUNT = "Amount"
	PAYMENT_DATE   = "PaymentDate"
	PAYMENT_STATUS = "PaymentStatus"

	USERS_TABLE    = "Users"
	USER_ID        = "UserID"
	USER_USERNAME  = "Username"
	USER_PASSWORD  = "Password"
	USER_EMAIL     = "Email"
	USER_FIRSTNAME = "FirstName"
	USER_LASTNAME  = "LastName"
	USER_ADDRESS   = "Address"

	MEMBER_TABLE     = "Members"
	MEMBERSHIP_LEVEL = "MembershipLevel"
	MEMBERSHIP_START = "MembershipStart"

	CART_TABLE    = "CartItem"
	CART_QUANTITY = "Quantity"

	ADMINSTRATORS_TABLE    = "Administrators"
	ADMINSTRATORS_ROLENAME = "RoleName"
)

var DB *sqlx.DB
var JwtKey []byte //Token used by the frontend/backend to check user privilegies
