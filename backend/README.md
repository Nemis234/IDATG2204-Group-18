in [Section `Features`](#feature)

## Feature


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
- [`GET /orders/{order_id}/payment` → `PaymentHandler`](#payment)
- [`GET /orders/{order_id}/items` → `OrderItemsHandler`](#orderitems)
- [`GET /orders/{order_id}/items/{item_id}` → `OrderItemHandler`](#orderitem)

### 🔄 Order Status
- [`GET /orderstatus` → `StatusHandler`](#status)
- [`GET /orderstatus/{statusName}` → `StatusHandler`](#status)

### 🛍️ Products
- [`GET /products` → `ProductsHandler`](#products)
- [`GET /products/{product_id}` → `ProductHandler`](#product)
- [`GET /products/{product_id}/reviews` → `ProductReviewsHandler`](#productreviews)
- [`GET /products/{product_id}/reviews/{user_id}` → `ReviewHandler`](#review)

### 👤 Users
- [`GET /users` → `UsersHandler`](#users)
- [`GET /users/{user_id}` → `UserHandler`](#user)
- [`GET /users/{user_id}/member` → `MemberHandler`](#member)
- [`GET /user/{user_id}/cart` → `CartHandler`](#cart)
- [`GET /users/{user_id}/reviews` → `UserReviewsHandler`](#userreviews)
- [`GET /users/{user_id}/reviews/{product_id}` → `ReviewHandler`](#review)
- [`POST /users/login` → `LoginHandler`](#login)

---

### Brands
<details>
<summary>Details about `BrandsHandler`</summary>

...

</details>

### Brand
<details>
<summary>Details about `BrandHandler`</summary>

...

</details>

### Categories
<details>
<summary>Details about `CategoriesHandler`</summary>

...

</details>

### Category
<details>
<summary>Details about `CategoryHandler`</summary>

...

</details>

### Orders
<details>
<summary>Details about `OrdersHandler`</summary>

...

</details>

### Order
<details>
<summary>Details about `OrderHandler`</summary>

...

</details>

### Payment
<details>
<summary>Details about `PaymentHandler`</summary>

...

</details>

### OrderItems
<details>
<summary>Details about `OrderItemsHandler`</summary>

...

</details>

### OrderItem
<details>
<summary>Details about `OrderItemHandler`</summary>

...

</details>

### Status
<details>
<summary>Details about `StatusHandler`</summary>

...

</details>

### Products
<details>
<summary>Details about `ProductsHandler`</summary>

...

</details>

### Product
<details>
<summary>Details about `ProductHandler`</summary>

...

</details>

### ProductReviews
<details>
<summary>Details about `ProductReviewsHandler`</summary>

...

</details>

### Review
<details>
<summary>Details about `ReviewHandler`</summary>

...

</details>

### Users
<details>
<summary>Details about `UsersHandler`</summary>

...

</details>

### User
<details>
<summary>Details about `UserHandler`</summary>

...

</details>

### Member
<details>
<summary>Details about `MemberHandler`</summary>

...

</details>

### Cart
<details>
<summary>Details about `CartHandler`</summary>

...

</details>

### UserReviews
<details>
<summary>Details about `UserReviewsHandler`</summary>

...

</details>

### Login
<details>
<summary>Details about `LoginHandler`</summary>

...

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


#### Product review - WIP
- Commenting - WIP
- Plural
    - GET - done
    - POST - done
- Singluar
    - GET - done
    - PUT - done
    - DELETE - done

#### Cart - WIP
- Commenting - WIP
- Plural
    - GET - WIP
    - POST - WIP
- Singluar
    - GET - WIP
    - PUT - WIP
    - PATCH - WIP
    - DELETE - WIP

#### Login - WIP
 - Commenting - WIP
 - POST - done

#### Members - WIP
- Commenting - WIP
- GET - WIP
- POST - WIP
- PUT - WIP
- DELETE - WIP

#### Users review - WIP
- Commenting - WIP
- Plural
    - GET - done
    - POST - done
- Singluar
    - GET - done
    - PUT - done
    - DELETE - done

#### Users - WIP
- Commenting - WIP
- Plural
    - GET - done
    - POST - done
- Singluar
    - GET - done
    - PUT - WIP
    - PATCH - WIP
    - DELETE - WIP