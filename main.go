package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Customer represents a customer in our system
type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// For now, we'll store customers in memory (no database yet)
var customers []Customer
var nextID = 1

func main() {
	// Add some sample customers
	customers = append(customers, Customer{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Phone:     "555-0101",
	})
	customers = append(customers, Customer{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane@example.com",
		Phone:     "555-0102",
	})
	nextID = 3

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health Insurance App API - Go to /api/customers to see customers")
	})

	// API endpoint to get all customers
	http.HandleFunc("/api/customers", getCustomers)

	// Start the server
	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Try: http://localhost:8080/api/customers")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Function to handle GET /api/customers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Convert customers to JSON and send response
	json.NewEncoder(w).Encode(customers)
}
