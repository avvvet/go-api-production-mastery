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
=======
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

type Shoe struct {
    ID      int     `json:"ID"`
    Brand   string  `json:"Brand"`
    Model   string  `json:"Model"`
    Size    float64 `json:"Size"`
    Color   string  `json:"Color"`
    Price   float64 `json:"Price"`
    InStock bool    `json:"InStock"`
}

var shoes = []Shoe{
    {1, "Puma", "Suede Classic", 8.0, "Blue", 75.00, true},
    {2, "New Balance", "990v5", 9.5, "Gray", 185.00, true},
    {3, "Jordan", "Air Jordan 1", 10.5, "Red/Black", 170.00, false},
    {4, "Reebok", "Club C", 7.5, "White", 80.00, true},
    {5, "Asics", "Gel-Kayano", 11.0, "Navy", 160.00, true},
}

func getShoes(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(shoes)
}

func getShoe(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, s := range shoes {
        if s.ID == id {
            json.NewEncoder(w).Encode(s)
            return
        }
    }
    http.Error(w, "Shoe not found", http.StatusNotFound)
}

func createShoe(w http.ResponseWriter, r *http.Request) {
    var shoe Shoe
    json.NewDecoder(r.Body).Decode(&shoe)
    maxID := 0
    for _, s := range shoes {
        if s.ID > maxID {
            maxID = s.ID
        }
    }
    shoe.ID = maxID + 1
    shoes = append(shoes, shoe)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(shoe)
}

func updateShoe(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var updated Shoe
    json.NewDecoder(r.Body).Decode(&updated)
    if updated.ID != id {
        http.Error(w, "ID in URL and body must match", http.StatusBadRequest)
        return
    }
    for i, s := range shoes {
        if s.ID == id {
            shoes[i] = updated
            json.NewEncoder(w).Encode(updated)
            return
        }
    }
    http.Error(w, "Shoe not found", http.StatusNotFound)
}

func deleteShoe(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for i, s := range shoes {
        if s.ID == id {
            shoes = append(shoes[:i], shoes[i+1:]...)
            json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
            return
        }
    }
    http.Error(w, "Shoe not found", http.StatusNotFound)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/shoes", getShoes).Methods("GET")
    r.HandleFunc("/shoes/{id}", getShoe).Methods("GET")
    r.HandleFunc("/shoes", createShoe).Methods("POST")
    r.HandleFunc("/shoes/{id}", updateShoe).Methods("PUT")
    r.HandleFunc("/shoes/{id}", deleteShoe).Methods("DELETE")
    log.Println("Server running on http://localhost:3000")
    log.Fatal(http.ListenAndServe(":3000", r))
}
>>>>>>> 4e7a09e (Add full Shoes API implementation in Go)
