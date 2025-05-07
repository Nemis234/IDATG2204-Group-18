package handlers

import (
	brandshandler "backend/handlers/brandsHandler"
	categorieshandler "backend/handlers/categoriesHandler"
	producthandler "backend/handlers/productHandler"
	usershandler "backend/handlers/usersHandler"
)

var ProductsHandler = producthandler.ProductsHandler
var ProductHandler = producthandler.ProductHandler

var CategoriesHandler = categorieshandler.CategoriesHandler
var CategoryHandler = categorieshandler.CategoryHandler

var BrandsHandler = brandshandler.BrandsHandler
var BrandHandler = brandshandler.BrandHandler

var UsersHandler = usershandler.UsersHandler
var UserHandler = usershandler.UserHandler
var LoginHandler = usershandler.LoginHandler
