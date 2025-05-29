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
	http.HandleFunc("/api/customers", handleCustomers)

	// Start the server
	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Try: http://localhost:8080/api/customers")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Function to handle both GET and POST for /api/customers
func handleCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// GET request - return all customers
		json.NewEncoder(w).Encode(customers)
	} else if r.Method == "POST" {
		// POST request - create new customer
		var newCustomer Customer

		// Read JSON from request body
		err := json.NewDecoder(r.Body).Decode(&newCustomer)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Assign ID and add to our list
		newCustomer.ID = nextID
		nextID++
		customers = append(customers, newCustomer)

		// Return the created customer
		json.NewEncoder(w).Encode(newCustomer)
	} else {
		// Other methods not allowed
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
