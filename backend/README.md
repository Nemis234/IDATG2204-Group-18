

## API Endpoints

### 🏷️ Brands
- [`GET /brands` → `BrandsHandler`](#brands)
- [`GET /brands/{id}` → `BrandHandler`](#brand)

### 📂 Categories
- [`GET /categories` → `CategoriesHandler`](#categories)
- [`GET /categories/{id}` → `CategoryHandler`](#category)

### 📦 Orders
- [`GET /orders` → `OrdersHandler`](#orders)
- [`GET /orders/{order_id}` → `OrderHandler`](#order)
- [`GET /orders/{order_id}/items` → `OrderItemsHandler`](#orderitems)
- [`GET /orders/{order_id}/items/{item_id}` → `OrderItemHandler`](#orderitem)
- [`GET /orders/{order_id}/payment` → `PaymentHandler`](#payment) WIP

### 🔄 Order Status
- [`GET /orderstatus` → `StatusHandler`](#status)
- [`GET /orderstatus/{statusName}` → `UpdateStatusHandler`](#update-status)

### 🛍️ Products
- [`GET /products` → `ProductsHandler`](#products)
- [`GET /products/{product_id}` → `ProductHandler`](#product)
- [`GET /products/{product_id}/reviews` → `ProductReviewsHandler`](#productreviews)
- [`GET /products/{product_id}/reviews/{user_id}` → `ReviewHandler`](#review)

### 👤 Users
- [`GET /users` → `UsersHandler`](#users)
- [`GET /users/{user_id}` → `UserHandler`](#user) WIP
- [`GET /users/{user_id}/member` → `MemberHandler`](#member)
- [`GET /user/{user_id}/cart` → `CartHandler`](#cart)
- [`GET /user/{user_id}/cart/{product_id}` → `CartItemHandler`](#cart-item)
- [`GET /users/{user_id}/reviews` → `UserReviewsHandler`](#userreviews)
- [`GET /users/{user_id}/reviews/{product_id}` → `ReviewHandler`](#review)
- [`POST /users/login` → `LoginHandler`](#login)

---

### Brands
<details>
<summary>Details about BrandsHandler</summary>

BrandsHandler supports these HTTP methods:
  - GET: Retrieves a list of all brands from the database.
  - POST: Inserts a new brand into the database.

#### GET

GET handles requests for a list of brands.
It retrieves all brands from the database and returns them in JSON format.

Example usage:

	Method: GET
	Route: /brands
	Response:
	[
		{
			"name": "Brand 1",
			"description": "Brand Description 1"
		},
		{
			"name": "Brand 2",
			"description": "Brand Description 2"
		}
	]

#### POST

POST handles requests to insert a new brand into the database.
It expects a JSON payload with the brand details.

Only admins can access this endpoint.

Example usage:

	Method: POST
	Route: /brands
	Request Body:
	{
		"name": "Brand 3",
		"description": "Brand Description 3"
	}
	Response:
	HTTP Status: 201 Created

</details>

### Brand
<details>
<summary>Details about `BrandHandler`</summary>

BrandHandler supports these HTTP methods:
  - GET: Retrieves details of a specific brand based on the name in the URL path.
  - PUT: Updates the details of a specific brand.
  - DELETE: Deletes a specific brand from the database.

#### GET

GET handles requests for a single brand.
It retrieves the brand details from the database based on the provided name in the URL path,
and returns the brand details in JSON format.

Example usage:

	Method: GET
	Route: /brands/Brand 1
	Response:
	{
		"name": "Brand 1",
		"description": "Brand Description 1"
	}

#### PUT

PUT handles requests to update a specific brand in the database.
It expects a JSON payload with the brand details.

Only admins can access this endpoint.

Example usage:

	Method: PUT
	Route: /brands/Brand 1
	Request Body:
	{
		"description": "Updated Brand Description"
	}
	Response:
	HTTP Status: 200 OK

#### DELETE

DELETE handles requests to delete a specific brand from the database.

Only admins can access this endpoint.

Example usage:

	Method: DELETE
	Route: /brands/Brand 1
	Response:
	HTTP Status: 204 No Content

</details>

### Categories
<details>
<summary>Details about `CategoriesHandler`</summary>
CategoriesHandler supports these HTTP methods:
  - GET: Retrieves a list of all categories from the database.
  - POST: Inserts a new category into the database.

#### GET

GET handles requests for a list of categories.
It retrieves all categories from the database and returns them in JSON format.

Example usage:

	Method: GET
	Route: /categories
	Response:
	[
		{
			"name": "Category 1",
			"description": "Category Description 1"
		},
		{
			"name": "Category 2",
			"description": "Category Description 2"
		}
	]

#### POST

POST handles requests to insert a new category into the database.
It expects a JSON payload with the category details.

Only admins can access this endpoint.

Example usage:

	Method: POST
	Route: /categories
	Request Body:
	{
		"name": "Category 3",
		"description": "Category Description 3"
	}
	Response:
	HTTP Status: 201 Created

</details>

### Category
<details>
<summary>Details about `CategoryHandler`</summary>
CategoryHandler supports these HTTP methods:
  - GET: Retrieves details of a specific category based on the name in the URL path.
  - PUT: Updates the details of a specific category.
  - DELETE: Deletes a specific category based on the name in the URL path.

#### GET

GET handles requests for a single category.
It retrieves the category details from the database based on the provided name in the URL path,
and returns the category details in JSON format.

Example usage:

	Method: GET
	Route: /categories/category 1
	Response:
	{
		"name": "Category 1",
		"description": "Category Description 1"
	}

#### PUT

PUT handles requests to update a specific category in the database.
It expects a JSON payload with the category details.

Only admins can access this endpoint.

Example usage:

	Method: PUT
	Route: /categories/category 1
	Request Body:
	{
		"description": "Updated Category Description"
	}
	Response:
	HTTP Status: 200 OK

#### DELETE

DELETE handles requests to delete a specific category from the database.
It retrieves the category details from the database based on the provided name in the URL path,
and deletes the category from the database.

Only admins can access this endpoint.

Example usage:

	Method: DELETE
	Route: /categories/category 1
	Response:
	HTTP Status: 204 No Content


</details>

### Orders
<details>
<summary>Details about `OrdersHandler`</summary>

OrdersHandler supports the following methods:
  - GET: Fetches all orders.
  - POST: Creates a new order.

#### GET

GET handles the retrieval of all orders from the database.
It queries the database for all orders and returns them as a JSON response.
It does not respond with the orders items.

Only admin users can access this endpoint.

Example usage:

	Method: GET
	Route: /orders
	Response:
	[
		{
			"order_id": "12345",
			"user_id": "67890",
			"order_date": "2023-10-01T12:00:00Z",
			"order_status": "Pending",
			"order_total": 100.50
			"items": null
		},
		{
			"order_id": "54321",
			"user_id": "09876",
			"order_date": "2023-10-02T12:00:00Z",
			"order_status": "Shipped",
			"order_total": 200.75
			"items": null
		}
	]

#### POST

POST handles the creation of a new order in the database.
It expects a JSON payload with the order details, in JSON format.
The order ID is generated by the database, and the new order is inserted into the database.

To add items to the order, an array of order items is included in the request body.
If any items are provided, they will be inserted into the database as well.
If any items fail to be inserted, like if a product ID does not exist, the entire transaction will be rolled back.

A user can only make an order if they are logged in,
and only if the user ID in the request matches the user ID in the JWT token.
Admin users can place orders for any user.

Mandatory fields cannot be null, while optional can be null.

The request body should contain the following fields:

	{
	- user_id	(string)| mandatory	: The ID of the user placing the order.
	- order_date	(string)| optional	: The date of the order in ISO 8601 format.
	- order_total	(float64)| mandatory	: The total amount of the order.
	- order_status	(string)| mandatory	: The status of the order, must be found in the OrderStatus table.
	- items	(array of objects)| optional	: The order_items in the order. Each item must contain the product ID and quantity. See below for the format.

		The order items should contain the following fields:
		{
		- product_id	(string)| mandatory	: The ID of the product.
		- quantity		(int)| mandatory	: The quantity of the product.
		}
	}

Example usage:

	Method: POST
	Route: /orders
	Request Body:
	{
		"user_id": "67890",
		"order_date": "2023-10-01T12:00:00Z",
		"order_status": "Pending",
		"order_total": 100.50,
		"items": [
			{
				"product_id": "54321",
				"quantity": 2
			},
			{
				"product_id": "67890",
				"quantity": 1
			}
		]
	}
	Response:
	Http Status: 201 Created
	{
		"id": "12345"
	}


</details>

### Order
<details>
<summary>Details about `OrderHandler`</summary>

OrderHandler supports the following methods:
  - GET: Fetches a specific order by ID.
  - PUT: Updates an existing order.
  - PATCH: Partially updates an existing order.
  - DELETE: Deletes an existing order.

Only users that created the order can see, update or delete it.
Admin users can update any order.

#### GET

GET handles the retrieval of a specific order by ID from the database.
It queries the database for the order with the specified ID and returns it as a JSON response.
It also fetches the order items associated with that order and includes them in the response.

Example usage:

	Method: GET
	Route: /orders/12345
	Response:
	Http Status: 200 OK
	{
		"order_id": "12345",
		"user_id": "67890",
		"order_date": "2023-10-01T12:00:00Z",
		"order_status": "Pending",
		"order_total": 100.50,
		"items": [
			{
				"order_id": "12345",
				"product_id": "54321",
				"quantity": 2
			},
			{
				"order_id": "12345",
				"product_id": "67890",
				"quantity": 1
			}
		]
	}

#### PUT

PUT handles the update of an existing order in the database.
It expects a JSON payload with the order details, in JSON format.
The order ID is specified in the URL, and the order details are provided in the request body.

Items associated with the order cannot be updated in this request.
User ID and Order ID cannot be changed.

Mandatory fields cannot be null, while optional can be null.

The request body should contain the following fields:

	{
	- order_date	(string)| optional	: The date of the order in ISO 8601 format.
	- order_total	(float64)| mandatory	: The total amount of the order.
	- order_status	(string)| mandatory	: The status of the order. Must be a status found in the OrderStatus table.
	}

Example usage:

	Method: PUT
	Route: /orders/12345
	Request Body:
	{
		"order_date": "2023-10-01T12:00:00Z",
		"order_status": "Pending",
		"order_total": 100.50
	}
	Response:
	Http Status: 200 OK

#### PATCH

PATCH handles the partial update of an existing order in the database.
The order ID is specified in the URL, and the order details are provided in the request body.

The request body can use any field(s) available in the PUT method, in the same format.

Mandatory fields cannot be null, while optional can be null.

#### DELETE

DELETE handles the deletion of an existing order in the database.
It expects the order ID to be specified in the URL.
The order is deleted from the database, and a 204 No Content response is returned.

Example usage:

	Method: DELETE
	Route: /orders/12345
	Response:
	Http Status: 204 No Content


</details>


### OrderItems
<details>
<summary>Details about `OrderItemsHandler`</summary>


OrderItemsHandler supports the following methods:
  - GET: Fetches all order items for a specific order by ID.
  - POST: Adds a new order item to an existing order.

Only users that created the order can see, update or delete it.
Admin users can update any order.

#### GET

GET handles the retrieval of all order items for a specific order by ID from the database.
It queries the database for all order items associated with the specified order ID and returns them as a JSON response.

Example usage:

	Method: GET
	Route: /orders/12345/items
	Response:
	Http Status: 200 OK
	[
		{
			"order_id": "12345",
			"product_id": "67890",
			"quantity": 2
		},
		{
			"order_id": "12345",
			"product_id": "54321",
			"quantity": 1
		}
	]

#### POST

POST handles the addition of a new order item to an existing order.
It expects a JSON payload with the order item details, in JSON format.

The order ID is specified in the URL, and the order item details are provided in the request body.
The order ID cannot be changed.
Mandatory fields cannot be null, while optional can be null.

The request body should contain the following fields:

	{
	- product_id	(string)| mandatory	: The ID of the product being ordered.
	- quantity	(int)| mandatory	: The quantity of the product in the order.
	}

Example usage:

	Method: POST
	Route: /orders/12345/items
	Request Body:
	{
		"product_id": "67890",
		"quantity": 2
	}
	Response:
	Http Status: 201 Created


</details>

### OrderItem
<details>
<summary>Details about `OrderItemHandler`</summary>


OrderItemHandler supports the following methods:
  - GET: Fetches a specific order item by order ID and item ID.
  - PUT: Updates an existing order item.
  - DELETE: Deletes an existing order item.

Only users that created the order can see, update or delete it.
Admin users can update any order.

#### GET

GET handles the retrieval of a specific order item by order ID and item ID from the database.
It queries the database for the order item with the specified order ID and item ID and returns it as a JSON response.

Example usage:

	Method: GET
	Route: /orders/12345/items/67890
	Response:
	Http Status: 200 OK
	{
		"order_id": "12345",
		"product_id": "67890",
		"quantity": 2
	}

#### PUT

PUT handles the update of an existing order item in the database.
It expects a JSON payload with the order item details, in JSON format.

The order ID and item ID are specified in the URL, and the order item details are provided in the request body.
The order ID and product ID cannot be changed.

Mandatory fields cannot be null, while optional can be null.

The request body should contain the following fields:

	{
	- quantity	(int)| mandatory	: The quantity of the product in the order.
	}

Example usage:

	Method: PUT
	Route: /orders/12345/items/67890
	Request Body:
	{
		"quantity": 2
	}
	Response:
	Http Status: 200 OK

#### DELETE

DELETE handles the deletion of an existing order item in the database.
It expects the order ID and item ID to be specified in the URL.
The order item is deleted from the database, and a 204 No Content response is returned.

Example usage:

	Method: DELETE
	Route: /orders/12345/items/67890
	Response:
	Http Status: 204 No Content


</details>

### Payment
<details>
<summary>Details about `PaymentHandler`</summary>

...

</details>

### Status
<details>
<summary>Details about `StatusHandler`</summary>


StatusHandler supports the following methods:
- GET: Fetches all order statuses. Requires admin privileges.
- POST: Creates a new order status. Requires admin privileges.

#### GET
The function retrieves all order statuses from the database and returns them as a JSON response.

Example usage:

	Method: GET
	URL: /orderstatus
	Response:
	HTTP code: 200 OK
	[
		{
			"status_name": "Pending",
			"status_desc": "Order is pending"
		},
		{
			"status_name": "Shipped",
			"status_desc": "Order has been shipped"
		}
	]

#### POST
The function creates a new order status in the database. The request body should contain the order status details in JSON format.

Example usage:

	Method: POST
	URL: /orderstatus
	Request body:
	{
		"status_name": "Delivered",
		"status_desc": "Order has been delivered"
	}
	Response:
	HTTP code: 201 Created

</details>

### Update Status
<details>
<summary>Details about `UpdateStatusHandler`</summary>

UpdateStatusHandler supports the following methods:
- PUT: Updates an existing order status. Requires admin privileges.
- DELETE: Deletes an existing order status. Requires admin privileges.

#### PUT
The function updates an existing order status in the database. The request body should contain the updated order status details in JSON format.

Example usage:

	Method: PUT
	URL: /orderstatus/{statusName}
	Request body:
	{
		"status_desc": "Order has been shipped"
	}
	Response:
	HTTP code: 200 OK

#### DELETE
The function deletes an existing order status from the database.

Example usage:

	Method: DELETE
	URL: /orderstatus/{statusName}
	Response:
	HTTP code: 204 No Content

</details>

### Products
<details>
<summary>Details about `ProductsHandler`</summary>


ProductsHandler support these methods:

  - GET requests for a list of products with optional search paramaters.
  - POST requests to create a new product.

#### GET

The function supports pagination and filtering based on various query parameters.
Pagination is done using the page number as a query parameter, and is mandatory.
The page number is used to determine the offset for the SQL query.

Supported query parameters:

	{
	- page 		(string)| mandatory	: The page number for pagination
	- search 	(string)| optional	: Filter products by name or description
	- category 	(string)| optional	: Filter products by category name
	- brand 	(string)| optional	: Filter products by brand name
	- priceMin 	(float64)| optional	: Filter products by minimum price
	- priceMax 	(float64)| optional	: Filter products by maximum price
	- pageLimit	(int64)| optional	: The number of products to return per page (default is 10)
	}

Example usage:

	Method: GET
	Route: /products?page=1&search=example&category=example1&brand=example2&priceMin=9.9&priceMax=100
	Response:
	HTTP code: 200 OK
	[
		{
			"product_id": "12345",
			"name": "example product",
			"description": "Product Description",
			"img_url": "https://example.com/image9.jpg",
			"price": 100.00,
			"stock_quantity": 50,
			"category_name": "example1",
			"brand_name": "example2",
		},
		{
			"product_id": "10",
			"name": "TechBrave Laptop Pro",
			"description": null,
			"img_url": null,
			"price": 15999.99,
			"stock_quantity": 40,
			"category_name": null,
			"brand_name": null,
	  	},
	]

#### POST

When using the POST method, the request body should contain the product details in JSON format.
Mandatory fields cannot be null, optional fields can be null.

Only admins can access this endpoint.

The request body should include the following fields:

	{
	- product_id     (string)| optional	: The ID of the product
	- name           (string)| mandatory	: The name of the product
	- description    (string)| optional	: The description of the product
	- img_url        (string)| optional	: The URL of the product image
	- price          (float64)| mandatory	: The price of the product
	- stock_quantity (int64)| mandatory	: The quantity of the product in stock
	- category_name  (string)| optional	: The name of the category
	- brand_name     (string)| optional	: The name of the brand
	}

Example usage:

	Method: POST
	Route: /products
	Request body:
	{
		"name": "Product 1, example text",
		"description": "Description 1",
		"img_url": "https://example.com/image1.jpg",
		"price": 99.99,
		"stock_quantity": 50,
		"category_name": "Audio & Headphones",
		"brand_name": "TechBrave"
	}

	Response:
	HTTP code: 201 Created
	{
		"id" : "7c063863-2b6e-11f0-ad3b-58cdc90b0639"
	}

</details>

### Product
<details>
<summary>Details about `ProductHandler`</summary>


ProductHandler support these methods:

  - GET requests for a single product.
  - PUT requests to update a whole product.
  - PATCH requests to update indevidual fields for a product.
  - DELETE requests to delete a product.

#### GET

It retrieves the product details from the database based on the provided ID in the URL path,
and returns the product details in JSON format.

Example usage:

	Method: GET
	Route: /products/12345
	Response:
	HTTP code: 200 OK
	{
		"product_id": "12345",
		"name": "example product",
		"description": "Product Description",
		"img_url": "https://example.com/image9.jpg",
		"price": 100.00,
		"stock_quantity": 50,
		"category_name": "example1",
		"brand_name": "example2",
	}

#### PUT

When using the PUT method, the request body should contain the product details in JSON format.
Mandatory fields cannot be null, while optional can be null.

Only admins can access this endpoint.

The request body should include the following fields:

	{
	- name           (string)| mandatory	: The name of the product
	- description    (string)| optional	: The description of the product
	- img_url        (string)| optional	: The URL of the product image
	- price          (float64)| mandatory	: The price of the product
	- stock_quantity (int64)| mandatory	: The quantity of the product in stock
	- category_name  (string)| optional	: The name of the category
	- brand_name     (string)| optional	: The name of the brand
	}

Example usage:

	Method: PUT
	Route: /products/12345
	Request body:
	{
		"name": "Updated Product",
		"description": "Updated Description",
		"img_url": "https://example.com/updated_image.jpg",
		"price": 150.00,
		"stock_quantity": 30,
		"category_name": "Updated Category",
		"brand_name": "Updated Brand"
	}
	Response:
	HTTP code: 200 OK

#### PATCH

When using the PATCH method, the request body should contain the product details in JSON format.
Any amount of fields can be updated, but mandatory fields cannot be null, while optional can be null.
The request body can use any field(s) available in the PUT method, in the same format.

Only admins can access this endpoint.

#### DELETE

When using the DELETE method, the product will be deleted from the database.
If the product is in a foreign key constraint, the delete will fail with a 409 Conflict error.

Only admins can access this endpoint.

Example usage:

	Method: DELETE
	Route: /products/12345
	Response:
	HTTP code: 204 No Content

</details>

### ProductReviews
<details>
<summary>Details about `ProductReviewsHandler`</summary>


ProductReviewsHandler support these methods:
  - GET retrieve all reviews for a product.
  - POST create a new review for a product.

#### GET

Retrieves all reviews for a product from the database.
Example usage:

	Method: GET
	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews
	Response:
	HTTP code: 200 OK
	[
		{
			"product_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
			"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
			"rating": 5,
			"comment": "Great product!",
			"created_at": "2023-10-01T12:00:00Z"
		},
	]

#### POST

Creates a new review for a product in the database.
Only a user that provides its own user_id can create a review.
An admin can create a review for any user.

Example usage:

	Method: POST
	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews
	Request body:
	{
		"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
		"rating": 5,
		"comment": "Great product!"
	}
	Response:
	HTTP code: 201 Created

</details>

### Review
<details>
<summary>Details about `ReviewHandler`</summary>


ReviewHandler support these methods:
  - GET retrieve a review for a product by user and product ID.
  - PUT update a review for a product by user and product ID.
  - PATCH partially update a review for a product by user and product ID.
  - DELETE delete a review for a product by user and product ID.

#### GET

Retrieves a review for a product by user and product ID from the database.

Example usage:

	Method: GET

	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews/0df01f83-a9a7-4afa-9b62-8d0bb9722849
	Response:
	HTTP code: 200 OK
	{
		"product_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
		"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
		"rating": 5,
		"comment": "Great product!",
		"created_at": "2023-10-01T12:00:00Z"
	}

#### PUT

Updates a review for a product by user and product ID in the database.
Only the user that created the review  or an admin can update it.

Example usage:

	Method: PUT
	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews/0df01f83-a9a7-4afa-9b62-8d0bb9722849
	Request body:
	{
		"rating": 4,
		"comment": "Good product!"
	}
	Response:
	HTTP code: 200 OK

#### PATCH
Partially updates a review for a product by user and product ID in the database.
Only the user that created the review  or an admin can update it.

All fields are optional, and all fields used in the PUT method are available.

#### DELETE
Deletes a review for a product by user and product ID from the database.
Only the user that created the review  or an admin can delete it.

Example usage:

	Method: DELETE
	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews/0df01f83-a9a7-4afa-9b62-8d0bb9722849
	Response:
	HTTP code: 204 No Content

</details>

### Users
<details>
<summary>Details about `UsersHandler`</summary>


UsersHandler supports these methods:

  - GET request for a list of users.
  - POST creates a new user.

#### GET

The functions requires admin privileges, this is sent by the frontend through a JWT token in the Header (Authorization) upon request.

General function flow:
-> Validates the Token sent from the client
-> Checks privileges
-> Query the database
-> Write a respond to the client

Example usage:

	Method: GET
	Route: /users
	Response:
	HTTP code: 200 OK
	[
		{
			"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
			"username": "helloWorld",
			"password": "$2a$10$a8ZHeovSX0/NitUQBkHHHeTe8FqVRlJEmet29pUYDjjqkQA6KGaum",
			"email": "john_doe@gmail.com",
			"first_name": "John",
			"last_name": "Doe",
			"address": "Yolostreet 15",
			"role": "admin"
		},
		{
			"user_id": "8d29df7f-887d-4a02-aa78-89f25a669a3a",
			"username": "janedoe",
			"password": "$2a$10$qOd0Mt3q9oHZrBwEYGZa2.D8w6xaofMoUSFMG4ZE2VhHOK2j6Pf/.",
			"email": "jane_doe@gmail.com",
			"first_name": "Jane",
			"last_name": "Doe",
			"address": "Main Street 5",
			"role": "user"
		}
	]

#### POST

Receives a payload with user informations and creates a new user.

General function flow:
-> Extract the user infor from the payload
-> Check if mandatory fields are missing
-> Query the database for existing username/email
-> Creates the struct to be stored in the database
-> Stores the data(new user) in the database
-> Write a respond to the client

Example usage:

	Method: POST
	Route: /users
	Request body:

	{

		"username":"will00",
		"password":"son00",
		"email":"willson00gmail.com",
		"first_name":"Will",
		"last_name":"Son",
		"address":"Big Street 28"
	}

	Response:
	HTTP code: 201 Created
	{
		"id" : "fiadg09b0-200b-4d08-939a-f90105f6546s"
	}

</details>

### User
<details>
<summary>Details about `UserHandler`</summary>

...

</details>

### Member
<details>
<summary>Details about `MemberHandler`</summary>


MemberHandler supports the following methods:
  - GET: Get member data for a user
  - POST: Add a new member
  - PUT: Update member data for a user
  - DELETE: Delete member data for a user

Only administrators and the user themselves can access any endpoint.

#### GET

This method retrieves member data for a user. It requires the user ID to be passed in the URL path.
Example usage:

	Method: GET
	URL: /users/456/member
	Response:
	HTTP Status: 200 OK
	Body:
	{
		"UserID": "456",
		"MembershipLevel": "Gold",
		"MembershipStart": "2023-01-01"
	}

#### POST

This method adds a new member.
It requires the user ID to be passed in the URL path and the membership level and start date to be passed in the request body.

Example usage:

	Method: POST
	URL: /users/456/member
	Body:
	{
		"MembershipLevel": "Gold",
		"MembershipStart": "2023-01-01"
	}
	Response:
	HTTP Status: 201 Created

#### PUT

This method updates member data for a user.
It requires the user ID to be passed in the URL path and the membership level and start date to be passed in the request body.

Example usage:

	Method: PUT
	URL: /users/456/member
	Body:
	{
		"MembershipLevel": "Platinum",
		"MembershipStart": "2023-01-01"
	}
	Response:
	HTTP Status: 200 OK

#### DELETE

This method deletes member data for a user.
It requires the user ID to be passed in the URL path.

Example usage:

	Method: DELETE
	URL: /users/456/member
	Response:
	HTTP Status: 204 No Content

</details>

### Cart
<details>
<summary>Details about `CartHandler`</summary>

CartHandler supports the following methods:
  - GET: Get all cart items for a user
  - POST: Add a new item to the cart
  - DELETE: Delete all car items for a user

Only administrators and the user themselves can access any endpoint.

#### GET

This method retrieves all cart items for a user. It requires the user ID to be passed in the URL path.

Example usage:

	Method: GET
	URL: /users/456/cart

	Response:
	HTTP Status: 200 OK
	Body:
	[
		{
			"ProductID": "123",
			"UserID": "456",
			"Quantity": 2
		},
		{
			"ProductID": "789",
			"UserID": "456",
			"Quantity": 1
		}
	]

#### POST

This method adds a new item to the cart for a user. It requires the user ID to be passed in the URL path and the product ID and quantity to be passed in the request body.

Example usage:

	Method: POST
	URL: /users/456/cart
	Body:
	{
		"ProductID": "123",
		"Quantity": 2
	}

	Response:
	HTTP Status: 201 Created

#### DELETE

This method deletes all cart items for a user. It requires the user ID to be passed in the URL path.

Example usage:

	Method: DELETE
	URL: /users/456/cart
	Response:
	HTTP Status: 204 No Content


</details>

### Cart item
<details>
<summary>Details about `CartItemHandler`</summary>


CartItemHandler supports the following methods:
  - GET: Get a specific cart item for a user
  - PUT: Update a specific cart item for a user
  - DELETE: Delete a specific cart item for a user

Only administrators and the user themselves can access any endpoint.

#### GET

This method retrieves a specific cart item for a user.
It requires the user ID and product ID to be passed in the URL path.

Example usage:

	Method: GET
	URL: /users/456/cart/123
	Response:
	HTTP Status: 200 OK
	Body:
	{
		"ProductID": "123",
		"UserID": "456",
		"Quantity": 2
	}

#### PUT

This method updates a specific cart item for a user.
It requires the user ID and product ID to be passed in the URL path and the new quantity to be passed in the request body.
Example usage:

	Method: PUT
	URL: /users/456/cart/123
	Body:
	{
		"Quantity": 3
	}
	Response:
	HTTP Status: 200 OK

#### DELETE

This method deletes a specific cart item for a user.
It requires the user ID and product ID to be passed in the URL path.

Example usage:

	Method: DELETE
	URL: /users/456/cart/123
	Response:
	HTTP Status: 204 No Content

</details>

### UserReviews
<details>
<summary>Details about `UserReviewsHandler`</summary>


UserReviewsHandler supports the following methods:
- GET: Fetches all reviews for a specific user.
- POST: Creates a new review for a specific user.

#### GET
The function retrieves all reviews for a specific user from the database and returns them as a JSON response.

Example usage:

	Method: GET
	URL: /users/{user_id}/reviews
	Response:

	HTTP code: 200 OK
	[
		{
		"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
		"product_id": "12345678-1234-5678-1234-567812345678",
		"comment": "Great product!",
		"rating": 5,
		"post_date": "2023-10-01T12:00:00Z"
		}
	]

#### POST

The function creates a new review for a specific user in the database. The request body should contain the review details in JSON format.
Only the user who made the review or an admin can access this endpoint.

Example usage:

	Method: POST
	URL: /users/{user_id}/reviews
	Request body:

	{
		"product_id": "12345678-1234-5678-1234-567812345678",
		"comment": "Great product!",
		"rating": 5
	}

	Response:
	HTTP code: 201 Created

</details>

### Login
<details>
<summary>Details about `LoginHandler`</summary>


LoginHandler support these methods:

  - POST Request for comparing login credentials and returns a JWT token used for authentications.

#### POST

It retrieves the user information from the database and compares the password.
Then generates a JWT and returns this to the client as a Header. This must be stored in the frontend
and will be used for later on protected requests that requires admin privileges.
It can be made to return a JSON body containing the JWT token if needed, but it does not support that for now.

General function flow :
-> Extract the Payload
-> Check if the payload is not empty
-> Query the database for the user login information
-> Compare the password
-> Generates a JWT token, this is sent back to the client. Used to distinguish user roles on protected requests
-> Writes a respond to the client

Possible Responds:
200: OK, user is authenticated
400: Bad Request, Payload is not complete/missing fields required
401: Unauthorized, wrong password or username/email
404: Not found, query returned 0 rows, no user found in the database with the given username/email
500: Internal Server Error

Example usage:

	Method: Post
	Route: /users/login
	Request body:
	{
		"Email": "john_doe@gmail.com",
		"Password": "helloWorld"
	}

	Respond body:
	HTTP code: 200 No Content
	Respond header:
	Authorization: Bearer "jwt-token"

</details>




# Progress

## Handlers

#### Brands - Done


#### Categories - Done


#### Orders - Done


#### OrderItems - Done


#### Payment - WIP
- Commenting - WIP
- Plural
    - GET - WIP
    - POST - WIP
- Singluar
    - GET - WIP
    - PUT - WIP
    - PATCH - WIP
    - DELETE - WIP

#### Status - Done


#### Product - Done


#### Product review - Done


#### Cart - Done

#### Login - Done

#### Members - Done

#### Users review - Done

#### Users - Done