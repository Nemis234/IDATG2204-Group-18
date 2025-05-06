package structs

import "database/sql"

type Product struct {
	ProductID     int      `json:"product_id" db:"ProductID"`
	Name          string   `json:"name" db:"ProductName"`
	Description   string   `json:"description" db:"ProductDesc"`
	ImgURL        string   `json:"img_url" db:"ProductImgUrl"`
	Price         int      `json:"price"`
	StockQuantity int      `json:"stock_quantity"`
	Category      Category `json:"category"`
	Brand         Brand    `json:"brand"`
}

type Category struct {
	ID          int            `json:"id"`
	Name        string         `json:"name" db:"CategoryName"`
	Description sql.NullString `json:"description" db:"CategoryDesc"`
}

type Brand struct {
	ID          int            `json:"id"`
	Name        string         `json:"name" db:"BrandName"`
	Description sql.NullString `json:"description" db:"BrandDesc"`
}

type User struct {
	ID        int            `json:"id" db:"UserID"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"` // Maybe not, yeah?
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Address   sql.NullString `json:"address"`
}

type Order struct {
	ID          int         `json:"id"`
	User        User        `json:"user"`
	OrderDate   string      `json:"order_date" db:"order_date"`
	TotalAmount int         `json:"total_amount" db:"total_amount"`
	Status      string      `json:"status"`
	Items       []OrderItem `json:"items"`
}

type OrderItem struct {
	ID       int     `json:"id"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	Subtotal int     `json:"subtotal"`
}

type Payment struct {
	ID            int    `json:"id"`
	Order         Order  `json:"order"`
	PaymentMethod string `json:"payment_method" db:"payment_method"`
	Amount        int    `json:"amount"`
	PaymentDate   string `json:"payment_date" db:"payment_date"`
	Status        string `json:"status"`
}

type Cart struct {
	ID        int        `json:"id" db:"CartID"`
	UserID    int        `json:"user_id"`
	CartItems []CartItem `json:"cart_items"`
}

type CartItem struct {
	CartID   int     `json:"cart_id" db:"CartID"`
	Product  Product `json:"product" db:"ProductID"`
	Quantity int     `json:"quantity"`
}
