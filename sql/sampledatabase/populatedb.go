package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"log"
	"time"
	"strings"
    "github.com/google/uuid"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)


func main() {
	var databasename string
	rawName, err := os.ReadFile("DATABASE_NAME.txt")
	if err != nil {
		fmt.Println("Failed to read the name of the database from file, will use the default name written in the code.")
		databasename = "idatg2204" //<---- REPLACE THIS MANUALLY IF YOU SEE THE ERROR OVER WHEN YOU RUN THIS CODE.
	}else {
		databasename = string(rawName)
	}

	dsn := "root:@tcp(127.0.0.1:3306)/" + databasename
	os.Setenv("DSN", dsn)

	//Open the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to open the database")
		panic(err)
	}
	defer db.Close()

	// Check connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		panic(err)
	}
	fmt.Println("Connected to the database")


	//Creates the tables
	sqlBytes, err := os.ReadFile("databasetables.sql")
	if err != nil {
		log.Fatal("Failed to read SQL file:", err)
	}
	sqlStatements := string(sqlBytes)

	queries := strings.Split(string(sqlStatements), ";")

	for _, query := range queries {
		trimmed := strings.TrimSpace(query)
		if trimmed == "" {
			continue
		}

		_, err := db.Exec(trimmed)
		if err != nil {
			log.Fatalf("Failed to execute SQL: %v\nStatement: %s", err, trimmed)
		}
	}

	//Creating a uuid for users to be put in the sample database.
	userIDs := []string{}
	for i := 0; i < 4; i++{
		id := uuid.New()
		userIDs = append(userIDs, id.String())
	}
	
	//Creating a uuid for products to be put in the sample database.
	productIDs := []string{}
	for i := 0; i < 20; i++{
		id := uuid.New()
		productIDs = append(productIDs, id.String())
	}

	//Creating a uuid for the order table to be put in the sample database.
	orderIDs := []string{}
	for i := 0; i < 10; i++{
		id := uuid.New()
		orderIDs = append(orderIDs, id.String())
	}

	//Clearing the database first
	clearTables(db)

	//Populating the brand table
	populateBrand(db)

	//Populating the category table
	populateCategory(db)

	//Populating the product table
	populateProduct(db, productIDs)

	//Populating the user table
	populateUser(db, userIDs)

	//Populating the review table
	populateReview(db, userIDs, productIDs)

	//Populating the orderstatus table
	populateOrderStatus(db)

	//Populate the ordertable
	populateOrderTable(db, userIDs, orderIDs)

	//Populate order items table
	populateOrderItems(db, productIDs, orderIDs)

	//Populate the payment table
	populatePayment(db, orderIDs)

	//Populate cart items
	populateCartItems(db, userIDs, productIDs)

	//Populate the member table
	populateMembers(db, userIDs)

	//Udating ordertotals, payment amount accordingly to the other tables, to make it more realistic
	updateOrderTotals(db)
	updatePaymentAmounts(db)

	//Assigning admin
	assignAdmin(db, userIDs)
}

/*
*   Clearing tables before populating the db with sample data.
*
*   db - This is the databased passed from the main function.
 */
func clearTables(db *sql.DB) {
	tables := []string{"OrderTable", "Users", "Administrators" ,"Members", "Product", "Brand", "Review", "CartItem", "Category", "OrderStatus", "OrderItem", "Payment"}
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			fmt.Printf("Failed to clear table %s: %v\n", table, err)
		} else {
			fmt.Printf("Cleared table %s\n", table)
		}
	}
}

/*
*   Hash the password before storing it. Using the bycrypt which is in GO's libary.
*   It includes salt.
*
*   password - This is the password string that is to be hashed
 */
func hashPassword(password string) (string, error) {
	// Generate a salt + hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

/*
*   Populating the user table, this will create 4 users in the table.
*
*   db - This is the databased passed from the main function.
 */
func populateUser(db *sql.DB, userID []string) {
	//Creating an uuid for the users and Hash the passwords before inserting
	hashedPassword1, err := hashPassword("helloWorld")
	if err != nil {
		fmt.Println("Failed to hash password.")
		panic(err)
	}

	hashedPassword2, err := hashPassword("janedoe")
	if err != nil {
		fmt.Println("Failed to hash password.")
		panic(err)
	}

	hashedPassword3, err := hashPassword("alice123")
	if err != nil {
		fmt.Println("Failed to hash password.")
		panic(err)
	}

	hashedPassword4, err := hashPassword("bob88")
	if err != nil {
		fmt.Println("Failed to hash password.")
		panic(err)
	}

	query := `
        INSERT INTO Users (UserID, Username, Password, Email, FirstName, LastName, Address) VALUES
        (?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?)
    `

	_, err = db.Exec(query,
		userID[0], "helloWorld", hashedPassword1, "john_doe@gmail.com", "John", "Doe", "Yolostreet 15",
		userID[1], "janedoe", hashedPassword2, "jane_doe@gmail.com", "Jane", "Doe", "Main Street 5", 
		userID[2], "alice123", hashedPassword3, "alice@gmail.com", "Alice", "Smith", "River Road 42", 
		userID[3], "bob88", hashedPassword4, "bob88@gmail.com", "Bob", "Johnson", "Mountain View 10", 
	)
	if err != nil {
		fmt.Println("Failed to insert users.")
		panic(err)
	}

	fmt.Println("Inserted 4 users.")

}

/*
*   Populating the brand table, this will create 3 brands in the table.
*
*   db - This is the databased passed from the main function.
 */
func populateBrand(db *sql.DB) {
	query := `
        INSERT INTO Brand (BrandName, BrandDesc) VALUES
        (?, ?),
        (?, ?),
        (?, ?)
    `

	_, err := db.Exec(query,
		"E-Gadgets", "E-Gadget brings everyday innovation to your fingertips with a wide selection of electronics and accessories that make life simpler, smarter, and more exciting.",
		"TechBrave", "TechBrave delivers bold, cutting-edge technology across smartphones, laptops, cameras, and more, empowering pioneers in every part of life.",
		"DigiTech", "DigiTech powers your lifestyle with a full range of smart electronics — from smartphones to home appliances — all designed for seamless digital living.",
	)
	if err != nil {
		fmt.Println("Failed to insert brands.")
		panic(err)
	}
	
	fmt.Println("Inserted 3 brands.")

}

/*
*   Populating the category table, this will create 10 categories in the table.
*
*   db - This is the databased passed from the main function.
 */
func populateCategory(db *sql.DB) {
	query := `
        INSERT INTO Category (CategoryName, CategoryDesc) VALUES
        (?, ?),
        (?, ?),
        (?, ?),
        (?, ?),
        (?, ?),
        (?, ?),
        (?, ?),
        (?, ?),
        (?, ?),
        (?, ?)
    `

	_, err := db.Exec(query,
		"Electronics", "Gadgets, smartphones, computers, and accessories.",
		"Computers & Accessories", "Desktops, laptops, components, and peripherals like keyboards and mice.",
		"Mobile Phones & Tablets", "Smartphones, tablets, and mobile accessories.",
		"Cameras & Photography", "Digital cameras, lenses, tripods, and photography gear.",
		"Home Appliances", "Refrigerators, washing machines, microwaves, and small kitchen devices.",
		"Gaming", "Consoles, video games, gaming accessories, and VR devices.",
		"Audio & Headphones", "Speakers, headphones, earphones, and audio systems.",
		"Smart Home", "Smart lights, thermostats, security cameras, and home automation.",
		"TV & Home Theater", "Televisions, projectors, sound systems, and accessories.",
		"Wearable Technology", "Smartwatches, fitness trackers, and health tech devices.",
	)
	if err != nil {
		fmt.Println("Failed to insert categories.")
		panic(err)
	}

	fmt.Println("Inserted 10 categories into Category table.")
}

/*
*   Populating the product table, this will create 20 products in the table.
*
*   db - This is the databased passed from the main function.
 */
func populateProduct(db *sql.DB, productID []string) {
	query := `
        INSERT INTO Product (ProductID, ProductName, ProductDesc, ProductImgUrl, Price, StockQuantity, Brand, Category)
        VALUES
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err := db.Exec(query,
		productID[0], "TechBrave X100 Smartphone", "High-end smartphone with AI features.", "https://example.com/img1.jpg", 8999.99, 100, nil, "Mobile Phones & Tablets",
		productID[1], "E-Gadgets Ultra HD TV 55\"", "Ultra HD smart TV with vibrant colors.", "https://example.com/img2.jpg", 12999.99, 30, "E-Gadgets", "TV & Home Theater",
		productID[2], "DigiTech Gaming Headset Pro", "Immersive gaming headset with surround sound.", "https://example.com/img3.jpg", 1999.99, 80, "DigiTech", "Audio & Headphones",
		productID[3], "TechBrave Smartwatch 2", "Fitness tracking smartwatch.", "https://example.com/img4.jpg", 2999.99, 150, "TechBrave", "Wearable Technology",
		productID[4], "E-Gadgets Home Assistant", "Voice-controlled smart home hub.", "https://example.com/img5.jpg", 1499.99, 50, "E-Gadgets", "Smart Home",
		productID[5], "DigiTech Bluetooth Speaker", "Portable speaker with rich bass.", "https://example.com/img6.jpg", 899.99, 200, "DigiTech", "Audio & Headphones",
		productID[6], "TechBrave Tablet 10\"", "Lightweight tablet for work and play.", "https://example.com/img7.jpg", 3999.99, 90, "TechBrave", "Mobile Phones & Tablets",
		productID[7], "E-Gadgets Noise Cancelling Headphones", "Noise-free music experience.", "https://example.com/img8.jpg", 2499.99, 70, "E-Gadgets", "Audio & Headphones",
		productID[8], "DigiTech Wireless Charger", "Fast wireless charging pad.", "https://example.com/img9.jpg", 599.99, 300, "DigiTech", "Mobile Phones & Tablets",
		productID[9], "TechBrave Laptop Pro", "Powerful laptop for professionals.", "https://example.com/img10.jpg", 15999.99, 40, "TechBrave", "Computers & Accessories",
		productID[10], "E-Gadgets Security Camera", "Smart WiFi home security camera.", "https://example.com/img11.jpg", 999.99, 100, "E-Gadgets", "Smart Home",
		productID[11], "DigiTech Gaming Mouse", "Precision gaming mouse.", "https://example.com/img12.jpg", 799.99, 120, "DigiTech", "Gaming",
		productID[12], "TechBrave Drone Pro", "Professional drone with 4K camera.", "https://example.com/img13.jpg", 8999.99, 25, "TechBrave", "Cameras & Photography",
		productID[13], "E-Gadgets 4K Projector", "Cinema-quality projector for home.", "https://example.com/img14.jpg", 4999.99, 45, "E-Gadgets", "TV & Home Theater",
		productID[14], "DigiTech Portable SSD 1TB", "High-speed portable storage.", "https://example.com/img15.jpg", 1499.99, 70, "DigiTech", "Computers & Accessories",
		productID[15], "TechBrave Fitness Band", "Basic fitness tracker with heart monitor.", "https://example.com/img16.jpg", 699.99, 110, "TechBrave", "Wearable Technology",
		productID[16], "E-Gadgets Smart Light Bulb", "WiFi-enabled color-changing bulb.", "https://example.com/img17.jpg", 299.99, 500, "E-Gadgets", "Smart Home",
		productID[17], "DigiTech VR Headset", "Immersive VR gaming experience.", "https://example.com/img18.jpg", 3999.99, 60, "DigiTech", "Gaming",
		productID[18], "TechBrave Power Bank 20000mAh", "Portable charger for mobile devices.", "https://example.com/img19.jpg", 499.99, 250, "TechBrave", "Mobile Phones & Tablets",
		productID[19], "E-Gadgets Robot Vacuum", "Smart robot vacuum cleaner.", "https://example.com/img20.jpg", 2999.99, 30, "E-Gadgets", "Smart Home",
	)

	if err != nil {
		fmt.Println("Failed to insert products.")
		panic(err)
	}

	fmt.Println("Inserted 20 products into Product table.")
}

/*
*   Get comments for review based on randomized rating
 */
func getCommentBasedOnRating(rating int) string {
	positiveComments := []string{
		"Fantastic product!", "Exceeded my expectations.", "Highly recommend it!", "Amazing quality!",
	}
	neutralComments := []string{
		"It's okay, does the job.", "Average, nothing special.", "Good, but could be better.",
	}
	negativeComments := []string{
		"Very disappointing.", "Not worth the money.", "Broke after a few days.", "Terrible quality!",
	}

	if rating >= 8 {
		return positiveComments[rand.Intn(len(positiveComments))]
	} else if rating >= 4 {
		return neutralComments[rand.Intn(len(neutralComments))]
	} else {
		return negativeComments[rand.Intn(len(negativeComments))]
	}
}

/*
*   Populating the review table, this will create 10 reviews in the table.
*
*   db - This is the databased passed from the main function.
 */
func populateReview(db *sql.DB, userID []string, productID []string) {

	for i := 0; i < 10; i++ {
		userIDToStore := userID[i%len(userID)]
		productIDToStore := productID[i%len(productID)]
		rating := rand.Intn(10) + 1 // 1 to 10
		comment := getCommentBasedOnRating(rating)
		postDate := time.Now().AddDate(0, 0, -rand.Intn(100)).Format("2006-01-02") // Random date in the past 100 days

		_, err := db.Exec(`INSERT INTO Review (UserID, ProductID, Comment, Rating, PostDate) VALUES (?, ?, ?, ?, ?)`,
			userIDToStore, productIDToStore, comment, rating, postDate)
		if err != nil {
			fmt.Println("Failed to insert review.")
			panic(err)
		}
	}
	fmt.Println("Inserted 10 reviews.")
}

/*
*   Populating the OrderStatus table.
*
*   db - This is the database passed from the main function.
 */
func populateOrderStatus(db *sql.DB) {
	statuses := []struct {
		Name string
		Desc string
	}{
		{"Pending", "Order received but not yet processed."},
		{"Processing", "Order is currently being prepared."},
		{"Shipped", "Order has been shipped to the customer."},
		{"Delivered", "Order successfully delivered to the customer."},
		{"Cancelled", "Order was cancelled by the customer or store."},
		{"Returned", "Customer has returned the order."},
		{"Refunded", "Customer has been refunded for the order."},
		{"Failed", "Order payment failed or could not be processed."},
	}

	for _, status := range statuses {
		_, err := db.Exec(`INSERT INTO OrderStatus (StatusName, StatusDesc) VALUES (?, ?)`, status.Name, status.Desc)
		if err != nil {
			fmt.Println("Failed to insert order status:", status.Name)
			panic(err)
		}
	}

	fmt.Println("Inserted order statuses into OrderStatus table.")
}

/*
*   Populating the order table, this will create 10 orders in the table.
*
*   db - This is the databased passed from the main function.
 */
func populateOrderTable(db *sql.DB, userID []string, orderID []string) {
	orderstatuses := []string{"Pending", "Processing", "Shipped", "Delivered", "Cancelled", "Returned", "Refunded", "Failed"}

	for i := 0; i < 10; i++ {
		orderIDToStore := orderID[i%len(orderID)]
		userIDToStore := userID[i%len(userID)]
		orderstatus := orderstatuses[i%len(orderstatuses)]
		ordertotal := 0                                                             //Using update to calculate this later
		orderDate := time.Now().AddDate(0, 0, -rand.Intn(100)).Format("2006-01-02") // Random date in the past 100 days

		_, err := db.Exec(`INSERT INTO OrderTable (OrderID, UserID, OrderDate, OrderStatus, OrderTotal) VALUES (?, ?, ?, ?, ?)`,
			orderIDToStore, userIDToStore, orderDate, orderstatus, ordertotal)
		if err != nil {
			fmt.Println("Failed to insert orders.")
			panic(err)
		}
	}

	fmt.Println("Inserted 10 orders.")
}

/*
*   Populating the orderitems table, this will create 20 orderitems in the table.
*
*   db - This is the databased passed from the main function.
 */
func populateOrderItems(db *sql.DB, productID []string, orderID []string) {

	for i := 0; i < 20; i++ {
		ordersIDToStore := orderID[i%len(orderID)]
		productIDToStore := productID[i%len(productID)]
		quantity := rand.Intn(3) + 1 // 1 to 3

		_, err := db.Exec(`INSERT INTO OrderItem (OrderID, ProductID, Quantity) VALUES (?, ?, ?)`,
			ordersIDToStore, productIDToStore, quantity)
		if err != nil {
			fmt.Println("Failed to insert order items.")
			panic(err)
		}
	}

	fmt.Println("Inserted 20 orders items.")
}

/*
*   Helper function to get the orderstatus to set the correct paymentstatus.
*
*   orderID - The id of the order.
*   db - This is the databased passed from the main function.
 */
func checkOrderStatus(orderID string, db *sql.DB) (string, error) {
	paymentStatus := []string{"pending", "successful", "failed"}
	var orderStatus string
	err := db.QueryRow("SELECT OrderStatus FROM OrderTable WHERE OrderID = ?", orderID).Scan(&orderStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No order found with that ID")
			return "", nil
		}
		fmt.Println("Error querying ordertable:", err)
		return "", err
	}

	if orderStatus == "Pending" || orderStatus == "Processing" {
		return paymentStatus[0], nil
	} else if orderStatus == "Shipped" || orderStatus == "Delivered" {
		return paymentStatus[1], nil
	} else {
		return paymentStatus[2], nil
	}

}

/*
*   Populating the payment table, this will create 10 payments in the table.
*
*   db - This is the databased passed from the main function.
 */
func populatePayment(db *sql.DB, orderID []string) {
	paymentmethod := []string{"card", "vipps", "bank_transfer"}

	for i := 0; i < 10; i++ {
		paymentID := uuid.New()
		orderIDToStore := orderID[i%len(orderID)]
		paymentMethodUsed := paymentmethod[i%len(paymentmethod)]
		paymentDate := time.Now().AddDate(0, 0, -rand.Intn(100)).Format("2006-01-02") // Random date in the past 100 days
		paymentStatus, err := checkOrderStatus(orderIDToStore, db)
		if err != nil {
			fmt.Println("Could not find paymentstatus for order id:", orderIDToStore)
		}

		_, err = db.Exec(`INSERT INTO Payment (PaymentID, OrderID, PaymentMethod, Amount, PaymentDate, PaymentStatus) VALUES (?, ?, ?, ? ,? ,?)`,
			paymentID, orderIDToStore, paymentMethodUsed, 0, paymentDate, paymentStatus)
		if err != nil {
			fmt.Println("Failed to insert payments.")
			panic(err)
		}
	}

	fmt.Println("Inserted 10 payments.")
}

/*
*   Populating the cartitem table, this will create 10 cart items in the table.
*
*   db - This is the databased passed from the main function.
 */
func populateCartItems(db *sql.DB, userID []string, productID []string) {

	for i := 0; i < 10; i++ {
		userIDToStore := userID[i%len(userID)]
		productIDToStore := productID[i%len(productID)]
		quantity := rand.Intn(3) + 1 //1 to 3

		_, err := db.Exec(`INSERT INTO CartItem (UserID, ProductID, Quantity) VALUES (?, ?, ?)`,
			userIDToStore , productIDToStore, quantity)
		if err != nil {
			fmt.Println("Failed to insert cart items.")
			panic(err)
		}
	}

	fmt.Println("Inserted 10 cart items.")
}

/*
*   Populating the member table, this will create 3 members in the table.
 */
func populateMembers(db *sql.DB, userID []string) {
	query := `
		INSERT INTO Members (UserID, MembershipLevel, MembershipStart) VALUES
		(?, ?, ?),
		(?, ?, ?),
		(?, ?, ?)
	`

	_, err := db.Exec(query,
		userID[0], "Gold", time.Now().Format(time.DateTime),
		userID[1], "Silver", time.Now().Format(time.DateTime),
		userID[2], "Platinum", time.Now().Format(time.DateTime),
	)
	if err != nil {
		fmt.Println("Failed to insert members.")
		panic(err)
	}
	fmt.Println("Inserted 3 members into Member table.")
}

/*
*   Updating the ordertable total in order to make it more realistic.
*
*   db - This is the databased passed from the main function.
 */
func updateOrderTotals(db *sql.DB) {
	rows, err := db.Query("SELECT OrderID FROM OrderTable")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var orderID string
		if err := rows.Scan(&orderID); err != nil {
			panic(err)
		}

		var totalAmount float64
		err := db.QueryRow(`
            SELECT IFNULL(SUM(Product.Price * OrderItem.Quantity), 0)
            FROM OrderItem
            JOIN Product ON OrderItem.ProductID = Product.ProductID
            WHERE OrderItem.OrderID = ?`, orderID).Scan(&totalAmount)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec("UPDATE OrderTable SET OrderTotal = ? WHERE OrderID = ?", totalAmount, orderID)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Updated order totals.")
}

/*
*   Updating the payment table's Amount in order to make it more realistic.
*
*   db - This is the databased passed from the main function.
 */
func updatePaymentAmounts(db *sql.DB) {
	_, err := db.Exec(`
        UPDATE Payment
        JOIN OrderTable ON Payment.OrderID = OrderTable.OrderID
        SET Payment.Amount = OrderTable.OrderTotal
    `)
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated payment amounts.")
}

/*
*	Assigning adminstrators in the adminstrator table.
*
*	db - This is the databased passed from the main function.
*/
func assignAdmin(db *sql.DB, userID []string){
		query := `
		INSERT INTO Administrators (UserID, RoleName) VALUES
		(?, ?)
	`
	_, err := db.Exec(query,
		userID[0], "admin", 
	)
	if err != nil {
		panic(err)
	}
	msg := fmt.Sprintf("Assign user: %s , as admin.", userID[0])
	fmt.Println(msg)
}
