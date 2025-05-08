package structs

import (
	"encoding/json"
	"time"
)

type NullField[T any] struct {
	Set   bool
	Value *T
}

func (nf *NullField[T]) UnmarshalJSON(data []byte) error {
	nf.Set = true
	if string(data) == "null" {
		nf.Value = nil
		return nil
	}

	var val T
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	nf.Value = &val
	return nil
}

type Product struct {
	ProductID     string  `json:"product_id" db:"ProductID"`
	Name          string  `json:"name" db:"ProductName"`
	Description   *string `json:"description" db:"ProductDesc"`
	ImgURL        *string `json:"img_url" db:"ProductImgUrl"`
	Price         float64 `json:"price" db:"Price"`
	StockQuantity int     `json:"stock_quantity" db:"StockQuantity"`
	CategoryName  *string `json:"category_name" db:"Category"`
	BrandName     *string `json:"brand_name" db:"Brand"`
}

type PatchProduct struct {
	ProductID     string             `json:"product_id" db:"ProductID"`
	Name          NullField[string]  `json:"name" db:"ProductName"`
	Description   NullField[string]  `json:"description" db:"ProductDesc"`
	ImgURL        NullField[string]  `json:"img_url" db:"ProductImgUrl"`
	Price         NullField[float64] `json:"price" db:"Price"`
	StockQuantity NullField[int]     `json:"stock_quantity" db:"StockQuantity"`
	CategoryName  NullField[string]  `json:"category_name" db:"Category"`
	BrandName     NullField[string]  `json:"brand_name" db:"Brand"`
}

type Category struct {
	Name        string  `json:"category_name" db:"CategoryName"`
	Description *string `json:"category_desc" db:"CategoryDesc"`
}

type Brand struct {
	Name        string  `json:"brand_name" db:"BrandName"`
	Description *string `json:"brand_desc" db:"BrandDesc"`
}

type User struct {
	UserID    string  `json:"user_id" db:"UserID"`
	Username  string  `json:"username" db:"Username"`
	Password  string  `json:"password" db:"Password"` // Maybe not, yeah?
	Email     string  `json:"email" db:"Email"`
	FirstName string  `json:"first_name" db:"FirstName"`
	LastName  string  `json:"last_name" db:"LastName"`
	Address   *string `json:"address" db:"Address"`
}

type Member struct {
	UserID          string     `json:"user_id" db:"UserID"`
	MembershipLevel string     `json:"membership_level" db:"MembershipLevel"`
	MembershipStart *time.Time `json:"membership_start" db:"MembershipStart"`
}

type Order struct {
	ID          string      `json:"order_id" db:"OrderID"`
	UserID      string      `json:"user_id" db:"UserID"`
	OrderDate   *time.Time  `json:"order_date" db:"OrderDate"`
	OrderStatus string      `json:"order_status" db:"OrderStatus"`
	OrderTotal  int         `json:"order_total" db:"OrderTotal"`
	Items       []OrderItem `json:"items"`
}

type OrderItem struct {
	OrderID   string `json:"order_id" db:"OrderID"`
	ProductID string `json:"product_id" db:"ProductID"`
	Quantity  int    `json:"quantity" db:"Quantity"`
}

type OrderStatus struct {
	StatusName string  `json:"status_name" db:"StatusName"`
	StatusDesc *string `json:"status_desc" db:"StatusDesc"`
}

type Payment struct {
	PaymentID     string  `json:"payment_id" db:"PaymentID"`
	OrderID       string  `json:"order_id" db:"OrderID"`
	PaymentMethod string  `json:"payment_method" db:"PaymentMethod"`
	Amount        float32 `json:"amount" db:"Amount"`
	PaymentDate   string  `json:"payment_date" db:"PaymentDate"`
	PaymentStatus *string `json:"payment_status" db:"PaymentStatus"`
}

type CartItem struct {
	UserID    string `json:"cart_id" db:"UserID"`
	ProductID string `json:"product_id" db:"ProductID"`
	Quantity  *int64 `json:"quantity" db:"Quantity"`
}

type Review struct {
	UserID    string    `json:"user_id" db:"UserID"`
	ProductID string    `json:"product_id" db:"ProductID"`
	Comment   *string   `json:"comment" db:"Comment"`
	Rating    *int16    `json:"rating" db:"Rating"`
	PostDate  time.Time `json:"post_date" db:"PostDate"`
}
