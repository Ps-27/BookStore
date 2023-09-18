package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Secret key for JWT (we should use a strong secret in a real application)
var secretKey = []byte("secret_key")


// CustomClaims represents the JWT claims
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func initDB() error {
	// Need to Replace these with  MySQL database credentials
	dsn := "username:password@tcp(mysql-host:port)/database"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return db.Ping()
}

// Middleware for authentication
func tokenRequired((next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Token is missing")
			return
		}

		tokenString := strings.Split(tokenHeader, " ")[1]
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid token")
			return
		}

    r = r.WithContext(userContext(r.Context(), claims.Username))
		next.ServeHTTP(w, r)
	})
}

func userContext(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, "user", username)
}

// Registration handler
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

// Login handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

// Create a JWT token
func createToken(username string) string {
	// ...
}

// Product details handler
func productDetailsHandler(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["productID"]
	product, exists := products[productID]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Product not found")
		return
	}

	// need to Implement recommendation logic here
	searchTerms := getUserSearchHistory(getCurrentUserID(r))
	recommendations := getRecommendations(searchTerms, getCurrentUserID(r))
  response := map[string]interface{}{
		"product":        product,
		"recommendations": recommendations,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Get recommendations based on user search history
func getRecommendations(searchTerms []string, userID string) []string {
	recommendedProductIDs := []string{}

	// Fetch users who have made similar searches
	similarUserIDs := []string{}
	query := "SELECT DISTINCT user_id FROM user_search_history WHERE search_term IN (?) AND user_id != ?"
	rows, err := db.Query(query, strings.Join(searchTerms, ","), userID)
	if err != nil {
		fmt.Printf("Error fetching similar users: %v\n", err)
		return recommendedProductIDs
	}
	defer rows.Close()

  for rows.Next() {
		var similarUserID string
		err := rows.Scan(&similarUserID)
		if err != nil {
			fmt.Printf("Error scanning rows: %v\n", err)
			continue
		}
		similarUserIDs = append(similarUserIDs, similarUserID)
	}

	// Fetching books searched by similar users but not searched by the current user
	query = "SELECT DISTINCT product_id FROM user_search_history WHERE user_id IN (?) AND product_id NOT IN (SELECT DISTINCT product_id FROM user_search_history WHERE user_id=?)"
	rows, err = db.Query(query, strings.Join(similarUserIDs, ","), userID)
	if err != nil {
		fmt.Printf("Error fetching recommended products: %v\n", err)
		return recommendedProductIDs
	}
	defer rows.Close() 

  for rows.Next() {
		var productID string
		err := rows.Scan(&productID)
		if err != nil {
			fmt.Printf("Error scanning rows: %v\n", err)
			continue
		}
		recommendedProductIDs = append(recommendedProductIDs, productID)
	}

	return recommendedProductIDs
}

// Getting the current user's ID (need to replace with  user identification logic)
func getCurrentUserID(r *http.Request) string {
	// need to Replace this with  logic to retrieve the user's ID based on the JWT token or other authentication mechanism
	return "user1"
}

// Get the search history of the current user
func getUserSearchHistory(userID string) []string {
	searchTerms := []string{}

	// Retrieve search history from the database for the given user
	query := "SELECT DISTINCT search_term FROM user_search_history WHERE user_id=?"
	rows, err := db.Query(query, userID)
	if err != nil {
		fmt.Printf("Error fetching search history: %v\n", err)
		return searchTerms
	}
	defer rows.Close()


  for rows.Next() {
		var searchTerm string
		err := rows.Scan(&searchTerm)
		if err != nil {
			fmt.Printf("Error scanning rows: %v\n", err)
			continue
		}
		searchTerms = append(searchTerms, searchTerm)
	}

	return searchTerms
}

// ...
func main() {
	if err := initDB(); err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}

	r := mux.NewRouter()

	// Register and login endpoints
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")

	// Product details endpoint with authentication middleware
	r.Handle("/product/{productID}", tokenRequired(http.HandlerFunc(productDetailsHandler))).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}


  

  


  
