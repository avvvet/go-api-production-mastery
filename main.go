package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// this represents our shoes data
type Shoe struct {
	ID      int     `json:"id"`
	Brand   string  `json:"brand"`
	Model   string  `json:"model"`
	Size    float64 `json:"size"`
	Color   string  `json:"color"`
	Price   float64 `json:"price"`
	InStock bool    `json:"in_stock"`
}

// Sample data - in a real app, this would come from a database
var shoes = []Shoe{
	{ID: 1, Brand: "Nike", Model: "Air Max 90", Size: 9.0, Color: "White", Price: 120.00, InStock: true},
	{ID: 2, Brand: "Adidas", Model: "Stan Smith", Size: 8.5, Color: "White/Green", Price: 85.00, InStock: true},
	{ID: 3, Brand: "Converse", Model: "Chuck Taylor", Size: 10.0, Color: "Black", Price: 65.00, InStock: false},
	{ID: 4, Brand: "Vans", Model: "Old Skool", Size: 9.5, Color: "Black/White", Price: 70.00, InStock: true},
	{ID: 5, Brand: "Nike", Model: "Air Force 1", Size: 11.0, Color: "White", Price: 110.00, InStock: true},
}

// getShoesHandler handles GET /shoes - returns all shoes
func getShoesHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Convert shoes slice to JSON and send response
	if err := json.NewEncoder(w).Encode(shoes); err != nil {
		// If encoding fails, return server error
		http.Error(w, "Failed to encode shoes data", http.StatusInternalServerError)
		return
	}
}

func main() {
	// register our first handler for the /shoes endpoint
	http.HandleFunc("/shoes", getShoesHandler)

	// lets start our server
	fmt.Println("🦶 Shoes Shop API starting on http://localhost:8080")
	fmt.Println("📡 Available endpoints:")
	fmt.Println("   GET /shoes - Get all shoes")
	fmt.Println("\n💡 Try it: curl http://localhost:8080/shoes")

	// Start server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
