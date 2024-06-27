package main

import (
	"log"
	"net/http"
	"receiptProcessor/internal/handler"
)


func main(){
	http.HandleFunc("/receipts/process", handler.ProcessReceiptsHandler)
	http.HandleFunc("/receipts/",handler.GetReceiptPointsHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}