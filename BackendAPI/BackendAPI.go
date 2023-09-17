// Import necessary packages
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// Define a Book struct to hold book details
type Book struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    Writer      string `json:"writer"`
    PublishDate string `json:"publishDate"`
    Category    string `json:"category"`
}
// Fetch book details from the database
func getBookDetails(w http.ResponseWriter, r *http.Request) {
    bookID := mux.Vars(r)["id"]

    // Establish a database connection (replace with your DB credentials)
    db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/yourdb")
    if err != nil {
        log.Fatal(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Query the database for book details
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

    // Log user activity to the database (replace with your actual user tracking logic)
    userID := r.Header.Get("X-User-Id") // Assuming you pass user ID in the header
    logUserActivity(userID, "Viewed book details for ID: "+bookID)

} 
//table name: books


// Define a function to log user activity
func logUserActivity(userID string, activity string) error {
    // Establish a database connection (replace with your DB credentials)
    db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/yourdb")
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
