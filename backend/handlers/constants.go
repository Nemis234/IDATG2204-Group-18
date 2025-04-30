package handlers

import (
	"database/sql"
	"fmt"
)

// Constants for database column names
const (
	ID            = "id"
	NAME          = "name"
	DESCRIPTION   = "description"
	PRICE         = "price"
	STOCKQUANTITY = "stockQuantity"

	CATEGORYID = "categoryID"
	BRANDID    = "brandID"

	CATEGORYNAME        = "name"
	CATEGORYDESCRIPTION = "description"

	BRANDNAME        = "name"
	BRANDDESCRIPTION = "description"
)

// SQL queries for products, categories, and brands
var queryProducts = fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM products", ID, NAME, DESCRIPTION, PRICE, STOCKQUANTITY, CATEGORYID, BRANDID)
var queryProduct = fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM products WHERE %s = ?", ID, NAME, DESCRIPTION, PRICE, STOCKQUANTITY, CATEGORYID, BRANDID, ID)
var queryCategories = fmt.Sprintf("SELECT %s, %s FROM categories WHERE %s = ?", CATEGORYNAME, CATEGORYDESCRIPTION, ID)
var queryBrands = fmt.Sprintf("SELECT %s, %s FROM brands WHERE %s = ?", BRANDNAME, BRANDDESCRIPTION, ID)

var db *sql.DB

func SetDB(DB *sql.DB) {
	db = DB
}
