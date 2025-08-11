package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func getAllShoes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(shoes)
}

// getShoeByID handles GET /shoes/{id} - returns a specific shoe by ID
// This function extracts the shoe ID from the URL, finds the shoe in the slice, and
// returns it as a JSON response.
func getShoeByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/shoes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, shoe := range shoes {
		if shoe.ID == id {
			json.NewEncoder(w).Encode(shoe)
			return
		}
	}
	http.NotFound(w, r)
}

// createShoe handles POST /shoes - creates a new shoe
// This function decodes the JSON request body into a Shoe struct and appends it to the shoes slice.
func createShoe(w http.ResponseWriter, r *http.Request) {
	var newShoe Shoe
	if err := json.NewDecoder(r.Body).Decode(&newShoe); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//auto-increment ID
	maxID := 0
	for _, s := range shoes {
		if s.ID > maxID {
			maxID = s.ID
		}
	}
	newShoe.ID = maxID + 1
	shoes = append(shoes, newShoe)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newShoe)
}

// update existing shoes
func updateShoe(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/shoes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updated Shoe
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if updated.ID != id {
		http.Error(w, "ID in URL and body do not match", http.StatusBadRequest)
		return
	}
	for i, s := range shoes {
		if s.ID == id {
			shoes[i] = updated
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.NotFound(w, r)
}

// deleteShoe handles DELETE /shoes/{id} - deletes a shoe by ID
// This function extracts the shoe ID from the URL, finds the shoe in the slice, and removes it.
func deleteShoe(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/shoes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, s := range shoes {
		if s.ID == id {
			shoes = append(shoes[:i], shoes[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "shoe with ID %d deleted\n", id)
			return
		}
	}
	http.NotFound(w, r)
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

// main function starts the HTTP server and defines the API routes
func main() {
	http.HandleFunc("/shoes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getShoesHandler(w, r)
		case http.MethodPost:
			createShoe(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/shoes/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getShoeByID(w, r)
		case http.MethodPut:
			updateShoe(w, r)
		case http.MethodDelete:
			deleteShoe(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("🦶 Shoes Shop API starting on http://localhost:8080")
	fmt.Println("📡 Available endpoints:")
	fmt.Println("   GET /shoes       - Get all shoes")
	fmt.Println("   POST /shoes      - Create a new shoe")
	fmt.Println("   GET /shoes/{id} - Get shoe by ID")
	fmt.Println("   PUT /shoes/{id} - Update shoe by ID")
	fmt.Println("   DELETE /shoes/{id} - Delete shoe by ID")
	fmt.Println("\n💡 Try it: curl http://localhost:8080/shoes")

	// Start server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
