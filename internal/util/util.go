package util

import (
    "strconv"
    "time"
    "receiptProcessor/internal/model"
)

func CalculatePoints(receipt model.Receipt) int {
    points := 0

    if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
        if total == float64(int(total)) {
            points += 50
        }
        if int(total*100)%25 == 0 {
            points += 25
        }
    }

    for _, char := range receipt.Retailer {
        if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
            points += 10
        }
    }

    if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
        if t.Hour() >= 14 && t.Hour() < 16 {
            points += 10
        }
    }

    points += len(receipt.Items) * 5
    return points
}
