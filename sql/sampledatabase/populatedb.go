package main

import (
    "database/sql"
    "fmt"
    "math/rand"
    "time" 
    "golang.org/x/crypto/bcrypt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Replace the "test_project" with the name of your database
    dsn := "root:@tcp(127.0.0.1:3306)/test_project" // << replace this

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

    //Clearing the database first
    clearTables(db)
    
    //Populating the brand table
    populateBrand(db)
    
    //Populating the category table
    populateCategory(db)

    //Populating the product table
    populateProduct(db)

    //Populating the user table
    populateUser(db)

    //Populating the review table
    populateReview(db)

    //Populating the orderstatus table
    populateOrderStatus(db)

    //Populate the ordertable
    populateOrderTable(db)

    //Populate order items table
    populateOrderItems(db)

    //Populate the payment table
    populatePayment(db)

    //Populate the cart table
    populateCart(db)

    //Populate cart items
    populateCartItems(db)

    //Udating ordertotals, payment amount accordingly to the other tables, to make it more realistic
    updateOrderTotals(db)
    updatePaymentAmounts(db)
}

/*
*   Clearing tables before populating the db with sample data.
*
*   db - This is the databased passed from the main function.
*/
func clearTables(db *sql.DB) {
    tables := []string{"Ordertable", "User", "Product","Brand" ,"Review", "Cart","Cartitem", "Category","OrderStatus", "OrderItem", "Payment"}
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
func populateUser(db *sql.DB) {
    // Hash the passwords before inserting
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
        INSERT INTO User (UserID, Username, Password, Email, FirstName, LastName, Address) VALUES
        (?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?),
        (?, ?, ?, ?, ?, ?, ?)
    `

    _, err = db.Exec(query,
        "0", "helloWorld", hashedPassword1, "john_doe@gmail.com", "John", "Doe", "Yolostreet 15",
        "1", "janedoe", hashedPassword2, "jane_doe@gmail.com", "Jane", "Doe", "Main Street 5",
        "2", "alice123", hashedPassword3, "alice@gmail.com", "Alice", "Smith", "River Road 42",
        "3", "bob88", hashedPassword4, "bob88@gmail.com", "Bob", "Johnson", "Mountain View 10",
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
func populateProduct(db *sql.DB) {
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
        "1", "TechBrave X100 Smartphone", "High-end smartphone with AI features.", "https://example.com/img1.jpg", 8999.99, 100, "TechBrave", "Mobile Phones & Tablets",
        "2", "E-Gadgets Ultra HD TV 55\"", "Ultra HD smart TV with vibrant colors.", "https://example.com/img2.jpg", 12999.99, 30, "E-Gadgets", "TV & Home Theater",
        "3", "DigiTech Gaming Headset Pro", "Immersive gaming headset with surround sound.", "https://example.com/img3.jpg", 1999.99, 80, "DigiTech", "Audio & Headphones",
        "4", "TechBrave Smartwatch 2", "Fitness tracking smartwatch.", "https://example.com/img4.jpg", 2999.99, 150, "TechBrave", "Wearable Technology",
        "5", "E-Gadgets Home Assistant", "Voice-controlled smart home hub.", "https://example.com/img5.jpg", 1499.99, 50, "E-Gadgets", "Smart Home",
        "6", "DigiTech Bluetooth Speaker", "Portable speaker with rich bass.", "https://example.com/img6.jpg", 899.99, 200, "DigiTech", "Audio & Headphones",
        "7", "TechBrave Tablet 10\"", "Lightweight tablet for work and play.", "https://example.com/img7.jpg", 3999.99, 90, "TechBrave", "Mobile Phones & Tablets",
        "8", "E-Gadgets Noise Cancelling Headphones", "Noise-free music experience.", "https://example.com/img8.jpg", 2499.99, 70, "E-Gadgets", "Audio & Headphones",
        "9", "DigiTech Wireless Charger", "Fast wireless charging pad.", "https://example.com/img9.jpg", 599.99, 300, "DigiTech", "Mobile Phones & Tablets",
        "10", "TechBrave Laptop Pro", "Powerful laptop for professionals.", "https://example.com/img10.jpg", 15999.99, 40, "TechBrave", "Computers & Accessories",
        "11", "E-Gadgets Security Camera", "Smart WiFi home security camera.", "https://example.com/img11.jpg", 999.99, 100, "E-Gadgets", "Smart Home",
        "12", "DigiTech Gaming Mouse", "Precision gaming mouse.", "https://example.com/img12.jpg", 799.99, 120, "DigiTech", "Gaming",
        "13", "TechBrave Drone Pro", "Professional drone with 4K camera.", "https://example.com/img13.jpg", 8999.99, 25, "TechBrave", "Cameras & Photography",
        "14", "E-Gadgets 4K Projector", "Cinema-quality projector for home.", "https://example.com/img14.jpg", 4999.99, 45, "E-Gadgets", "TV & Home Theater",
        "15", "DigiTech Portable SSD 1TB", "High-speed portable storage.", "https://example.com/img15.jpg", 1499.99, 70, "DigiTech", "Computers & Accessories",
        "16", "TechBrave Fitness Band", "Basic fitness tracker with heart monitor.", "https://example.com/img16.jpg", 699.99, 110, "TechBrave", "Wearable Technology",
        "17", "E-Gadgets Smart Light Bulb", "WiFi-enabled color-changing bulb.", "https://example.com/img17.jpg", 299.99, 500, "E-Gadgets", "Smart Home",
        "18", "DigiTech VR Headset", "Immersive VR gaming experience.", "https://example.com/img18.jpg", 3999.99, 60, "DigiTech", "Gaming",
        "19", "TechBrave Power Bank 20000mAh", "Portable charger for mobile devices.", "https://example.com/img19.jpg", 499.99, 250, "TechBrave", "Mobile Phones & Tablets",
        "20", "E-Gadgets Robot Vacuum", "Smart robot vacuum cleaner.", "https://example.com/img20.jpg", 2999.99, 30, "E-Gadgets", "Smart Home",
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
func populateReview(db *sql.DB) {
    users := []string{"0", "1", "2", "3"}
    products := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}

    for i := 0; i < 10; i++ {
        userID := users[i%len(users)]
        productID := products[i%len(products)]
        rating := rand.Intn(10) + 1 // 1 to 10
        comment := getCommentBasedOnRating(rating)
        postDate := time.Now().AddDate(0, 0, -rand.Intn(100)).Format("2006-01-02") // Random date in the past 100 days

        _, err := db.Exec(`INSERT INTO Review (UserID, ProductID, Comment, Rating, PostDate) VALUES (?, ?, ?, ?, ?)`,
            userID, productID, comment, rating, postDate)
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
func populateOrderTable(db *sql.DB) {
    users := []string{"0", "1", "2", "3"}
    orderstatuses := []string{"Pending", "Processing", "Shipped", "Delivered", "Cancelled", "Returned", "Refunded", "Failed"}

    for i := 0; i < 10; i++ {
        userID := users[i%len(users)]
        orderstatus := orderstatuses[i%len(orderstatuses)]
        ordertotal := 0 //Using update to calculate this later
        orderDate := time.Now().AddDate(0, 0, -rand.Intn(100)).Format("2006-01-02") // Random date in the past 100 days

        _, err := db.Exec(`INSERT INTO Ordertable (OrderID, UserID, OrderDate, OrderStatus, OrderTotal) VALUES (?, ?, ?, ?, ?)`,
            i, userID, orderDate, orderstatus, ordertotal)
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
func populateOrderItems(db *sql.DB) {
    orders := []string{"0", "1", "2", "3","4","5","6","7","8","9"}
    products := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}

    for i := 0; i < 20; i++ {
        ordersID := orders[i%len(orders)]
        productID := products[i%len(products)]
        quantity := rand.Intn(3) + 1 // 1 to 3

        _, err := db.Exec(`INSERT INTO Orderitem (OrderID, ProductID, Quantity) VALUES (?, ?, ?)`,
            ordersID, productID, quantity)
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
    err := db.QueryRow("SELECT OrderStatus FROM Ordertable WHERE OrderID = ?", orderID).Scan(&orderStatus)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No order found with that ID")
            return "", nil
        }
        fmt.Println("Error querying ordertable:", err)
        return "", err
    }

    if (orderStatus == "Pending" || orderStatus == "Processing"){
        return paymentStatus[0] ,nil
    }else if(orderStatus == "Shipped" || orderStatus == "Delivered"){
        return paymentStatus[1] ,nil
    }else{
        return paymentStatus[2] ,nil
    }

}

/*
*   Populating the payment table, this will create 10 payments in the table.
*
*   db - This is the databased passed from the main function.
*/
func populatePayment(db *sql.DB) {
    orders := []string{"0", "1", "2", "3","4","5","6","7","8","9"}
    paymentmethod := []string{"card", "vipps", "bank_transfer"}

    for i := 0; i < 10; i++ {
        ordersID := orders[i]
        paymentMethodUsed := paymentmethod[i%len(paymentmethod)] 
        paymentDate := time.Now().AddDate(0, 0, -rand.Intn(100)).Format("2006-01-02") // Random date in the past 100 days
        paymentStatus, err := checkOrderStatus(ordersID, db)
        if err != nil {
            fmt.Println("Could not find paymentstatus for order id:", ordersID)
        }

        _, err = db.Exec(`INSERT INTO Payment (PaymentID, OrderID, PaymentMethod, Amount, PaymentDate, PaymentStatus) VALUES (?, ?, ?, ? ,? ,?)`,
            i, ordersID, paymentMethodUsed, 0, paymentDate, paymentStatus)
        if err != nil {
            fmt.Println("Failed to insert payments.")
            panic(err)
        }
    }

    fmt.Println("Inserted 10 payments.")
}


/*
*   Populating the cart table, one cart per user.
*
*   db - This is the database passed from the main function.
*/
func populateCart(db *sql.DB) {
    users := []string{"0", "1", "2", "3"}

    for i, userID := range users {
        _, err := db.Exec(`INSERT INTO Cart (CartID, UserID) VALUES (?, ?)`, i, userID)
        if err != nil {
            fmt.Println("Failed to insert cart for user:", userID)
            panic(err)
        }
    }

    fmt.Println("Inserted carts for all users.")
}

/*
*   Populating the cartitem table, this will create 10 cart items in the table.
*
*   db - This is the databased passed from the main function.
*/
func populateCartItems(db *sql.DB) {
    carts := []string{"0","1","2","3"}
    products := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}

    for i := 0; i < 10; i++ {
        cartID := carts[i%len(carts)]
        productID := products[i%len(products)]
        quantity := rand.Intn(3) + 1 //1 to 3

        _, err := db.Exec(`INSERT INTO Cartitem (CartID, ProductID, Quantity) VALUES (?, ?, ?)`,
            cartID, productID, quantity)
        if err != nil {
            fmt.Println("Failed to insert cart items.")
            panic(err)
        }
    }

    fmt.Println("Inserted 10 cart items.")
}

/*
*   Updating the ordertable total in order to make i more realistic.
*
*   db - This is the databased passed from the main function.
*/
func updateOrderTotals(db *sql.DB) {
    rows, err := db.Query("SELECT OrderID FROM Ordertable")
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

        _, err = db.Exec("UPDATE Ordertable SET OrderTotal = ? WHERE OrderID = ?", totalAmount, orderID)
        if err != nil {
            panic(err)
        }
    }

    fmt.Println("Updated order totals.")
}

/*
*   Updating the payment table's Amount in order to make i more realistic.
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
