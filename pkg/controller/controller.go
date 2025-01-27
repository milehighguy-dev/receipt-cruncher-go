package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/milehighguy-dev/receipt-cruncher-go/pkg/model"
    "strings"
	"github.com/google/uuid"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func PostReceiptHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var receipt model.Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, "Error parsing JSON", http.StatusBadRequest)
        return
    }

    fmt.Println("Received receipt: %+v\n", receipt)

	// dummy response
	response := model.ProcessResponse{
		ID: uuid.New(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Extract the ID from the URL path
	// The URL path will be in the format /receipts/{id}/points
	// net/http does not have a built-in way to extract path parameters
    path := strings.TrimPrefix(r.URL.Path, "/receipts/")
    parts := strings.Split(path, "/")
    if len(parts) < 2 || parts[1] != "points" {
        http.Error(w, "Invalid URL path", http.StatusBadRequest)
        return
    }
    id := parts[0]
	fmt.Println("got receipt with id: %+v\n", id)

	//TODO logic for retrieving points from DB

    points := model.Points{
        Points: 100, // Dummy points value
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(points); err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }
}

func RegisterRoutes() {
	baseURL := "/receipts"
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc(baseURL + "/process", PostReceiptHandler)
	http.HandleFunc(baseURL + "/", GetPointsHandler)
}