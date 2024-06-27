package handler

import (
    "receiptProcessor/internal/model"
    "receiptProcessor/internal/util"
    "encoding/json"
    "github.com/google/uuid"
    "net/http"
    "strings"
    
)

var receipts = make(map[string] model.ReceiptData)

func ProcessReceiptsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
        return
    }

    var receipt model.Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, "Invalid receipt format", http.StatusBadRequest)
        return
    }


	points := util.CalculatePoints(receipt)
    id := uuid.New().String()


	

	
	receipts[id] = model.ReceiptData{Receipt: receipt, Points: points}
    json.NewEncoder(w).Encode(struct{ ID string `json:"id"` }{ID: id})
}

func GetReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
        return
    }

    pathSegments := strings.Split(r.URL.Path, "/")
    
    // Check if the path format is correct
    if len(pathSegments) < 4 || pathSegments[3] != "points" {
        http.Error(w, "Invalid URL format", http.StatusBadRequest)
        return
    }

    id := pathSegments[2]
    receiptData, exists := receipts[id]
    if !exists {
        http.Error(w, "Receipt not found", http.StatusNotFound)
        return
    }

    // Return the stored points
    json.NewEncoder(w).Encode(struct{ Points int `json:"points"` }{Points: receiptData.Points})
}
