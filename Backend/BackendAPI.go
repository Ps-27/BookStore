// Import necessary packages
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// Defination of a Book struct to hold book details
type Book struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    Writer      string `json:"writer"`
    PublishDate string `json:"publishDate"`
    Category    string `json:"category"`
} 
// Defination of a CartItem struct to represent items in the cart
type CartItem struct {
    BookID    int    `json:"book_id"`
    Quantity  int    `json:"quantity"`
    // Add more properties as needed (e.g., book details)
}

// Defination of a Cart struct to represent the user's cart
type Cart struct {
    UserID     int         `json:"user_id"`
    Items      []CartItem  `json:"items"`
}




// Fetch book details from the database
func getBookDetails(w http.ResponseWriter, r *http.Request) {
    bookID := mux.Vars(r)["id"]

    // Establish a database connection (replace with  DB credentials)
    db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/db")
    if err != nil {
        log.Fatal(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Query the database for book details on speficic key example- id or name
    var book Book
    err = db.QueryRow("SELECT id, title, author, writer, publish_date, category FROM books WHERE id=?", bookID).Scan(
        &book.ID, &book.Title, &book.Author, &book.Writer, &book.PublishDate, &book.Category)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusNotFound)
        return
      }

    // Return book details as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book) 

    // Log user activity to the database (replace with  actual user tracking logic)
    userID := r.Header.Get("X-User-Id") // Assuming you pass user ID in the header
    logUserActivity(userID, "Viewed book details for ID: "+bookID)

} 
//table name: books


// Defination of a function to log user activity
func logUserActivity(userID string, activity string) error {
    // Establish a database connection (replace with  DB credentials)
    db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/db")
    if err != nil {
        log.Println("Error opening database connection:", err)
        return err
    }
    defer db.Close()

    // Prepare an SQL statement to insert user activity
    stmt, err := db.Prepare("INSERT INTO user_activity (user_id, activity) VALUES (?, ?)")
    if err != nil {
        log.Println("Error preparing SQL statement:", err)
        return err
    }
    defer stmt.Close() 
    // Execute the SQL statement to insert user activity
    _, err = stmt.Exec(userID, activity)
    if err != nil {
        log.Println("Error inserting user activity:", err)
        return err
    }

    log.Printf("User activity logged - UserID: %s, Activity: %s\n", userID, activity)
    return nil
}

//inserted user activity into a table named user_activity. table with appropriate columns (e.g., id, user_id, activity, timestamp).


// API endpoint to add a book to the user's cart
func addToCart(w http.ResponseWriter, r *http.Request) {
    // Parse book ID and quantity from the request
    // Add the book to the user's cart (create/update cart in  data store)
    // Respond with updated cart data
}

// API endpoint to view the user's cart
func viewCart(w http.ResponseWriter, r *http.Request) {
    // Retrieve the user's cart from the data store
    // Respond with cart data
}

// API endpoint to update cart item quantities or remove items
func updateCart(w http.ResponseWriter, r *http.Request) {
    // Parse item updates from the request
    // Update the user's cart in the data store
    // Respond with updated cart data
}

// Handle fetching the highest demanding books
func getHighestDemandingBooks(w http.ResponseWriter, r *http.Request) {
    // Query the useractivity table in  MySQL database
    // Calculate the highest demanding books based on user interactions
    // Return the results as JSON response
}
