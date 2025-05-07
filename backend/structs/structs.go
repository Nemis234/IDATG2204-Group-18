package structs

import (
	"database/sql"
	"time"
)

type Product struct {
	ProductID     int            `json:"product_id" db:"ProductID"`
	Name          string         `json:"name" db:"ProductName"`
	Description   sql.NullString `json:"description" db:"ProductDesc"`
	ImgURL        sql.NullString `json:"img_url" db:"ProductImgUrl"`
	Price         float64        `json:"price" db:"Price"`
	StockQuantity int            `json:"stock_quantity" db:"StockQuantity"`
	BrandName     sql.NullString `json:"brand_name" db:"Brand"`
	CategoryName  sql.NullString `json:"category_name" db:"Category"`
}

type Category struct {
	Name        string         `json:"name" db:"CategoryName"`
	Description sql.NullString `json:"description" db:"CategoryDesc"`
}

type Brand struct {
	Name        string         `json:"name" db:"BrandName"`
	Description sql.NullString `json:"description" db:"BrandDesc"`
}

type User struct {
	UserID    int            `json:"user_id" db:"UserID"`
	Username  string         `json:"username" db:"Username"`
	Password  string         `json:"password" db:"Password"` // Maybe not, yeah?
	Email     string         `json:"email" db:"Email"`
	FirstName string         `json:"first_name" db:"FirstName"`
	LastName  string         `json:"last_name" db:"LastName"`
	Address   sql.NullString `json:"address" db:"Address"`
}

type Member struct {
	ID              int          `json:"id" db:"UserID"`
	MembershipLevel string       `json:"membership_level" db:"MembershipLevel"`
	MembershipStart sql.NullTime `json:"membership_start" db:"MembershipStart"`
}

type Order struct {
	ID          int          `json:"order_id" db:"OrderID"`
	UserID      int          `json:"user_id" db:"UserID"`
	OrderDate   sql.NullTime `json:"order_date" db:"OrderDate"`
	OrderStatus string       `json:"order_status" db:"OrderStatus"`
	OrderTotal  int          `json:"order_total" db:"OrderTotal"`
	Items       []OrderItem  `json:"items"`
}

type OrderItem struct {
	OrderID   int `json:"order_id" db:"OrderID"`
	ProductID int `json:"product_id" db:"ProductID"`
	Quantity  int `json:"quantity" db:"Quantity"`
}

type OrderStatus struct {
	ID         int            `json:"id" db:"StatusName"`
	StatusDesc sql.NullString `json:"status" db:"StatusDesc"`
}

type Payment struct {
	PaymentID     int            `json:"payment_id" db:"PaymentID"`
	OrderID       int            `json:"order_id" db:"OrderID"`
	PaymentMethod string         `json:"payment_method" db:"PaymentMethod"`
	Amount        float32        `json:"amount" db:"Amount"`
	PaymentDate   string         `json:"payment_date" db:"PaymentDate"`
	PaymentStatus sql.NullString `json:"payment_status" db:"PaymentStatus"`
}

type CartItem struct {
	UserID    int           `json:"cart_id" db:"UserID"`
	ProductID string        `json:"product_id" db:"ProductID"`
	Quantity  sql.NullInt32 `json:"quantity" db:"Quantity"`
}

type Review struct {
	UserID    int            `json:"user_id" db:"UserID"`
	ProductID int            `json:"product_id" db:"ProductID"`
	Comment   sql.NullString `json:"comment" db:"Comment"`
	Rating    sql.NullInt16  `json:"rating" db:"Rating"`
	PostDate  time.Time      `json:"post_date" db:"PostDate"`
}
