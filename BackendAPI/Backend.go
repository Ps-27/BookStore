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
} 

//table name: books
