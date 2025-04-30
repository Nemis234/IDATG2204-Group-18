package handler

type ListProducts struct {
	Products []Product `json:"products"`
}

type Product struct {
	ProductID     int      `json:"product_id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Price         int      `json:"price"`
	StockQuantity int      `json:"stock_quantity"`
	Category      Category `json:"category"`
	Brand         Brand    `json:"brand"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Brand struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"` // Maybe not, yeah?
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

type Order struct {
	ID          int         `json:"id"`
	User        User        `json:"user"`
	OrderDate   string      `json:"order_date"`
	TotalAmount int         `json:"total_amount"`
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
	PaymentMethod string `json:"payment_method"`
	Amount        int    `json:"amount"`
	PaymentDate   string `json:"payment_date"`
	Status        string `json:"status"`
}
