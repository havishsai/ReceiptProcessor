package util

import (
    "fmt"
	"receiptProcessor/internal/model"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints(receipt model.Receipt) int {
    points := 0

    if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
        if total == float64(int(total)) {
            points += 50 // Adding 50 Points if Purchase Total is complete dollars without any cents
        }

        if int(total*100)%25 == 0 {
            points += 25 //Adding 25 points if cents are multiple of 25
        }
    }

    for _, char := range receipt.Retailer {
        if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
            points += 1 //Adding 1 point for each Alphanumeric Character
        }
    }

    if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
        if t.Hour() >= 14 && t.Hour() < 16 {
            points += 10 //Adding 10 points if purchase Time is between 2:00 PM and 4:00 PM
        }
    }

    date, date_err := time.Parse("2006-01-02", receipt.PurchaseDate) 
    if date_err == nil{
        fmt.Println(date.Day())
        if date.Day() % 2 != 0{
            points += 6 //Adding 6 points if day is odd
        }
    } else{
        fmt.Println(date_err)
    }

    for _, item := range receipt.Items{
        sd := item.ShortDescription
        trimmed_sd := strings.Trim(sd," ")
        if len(trimmed_sd)%3 == 0{
            points += 3 //Adding 3 points for each item if the trimmed Short Description is having length multiple of 3
        }
    }

    points += int((len(receipt.Items)/2)) * 5 // Calculating Pairs of Items and adding 5 points for each pair
    return points
}
