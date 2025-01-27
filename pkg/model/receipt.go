package model

import (
	"time"
	"github.com/google/uuid"
	"encoding/json"
)

type Receipt struct {
    ID           uuid.UUID `json:"id,omitempty"`
    Retailer     string    `json:"retailer" validate:"required"`
    PurchaseDate time.Time `json:"purchaseDate" validate:"required"`
    PurchaseTime time.Time `json:"purchaseTime" validate:"required"`
    Items        []Item    `json:"items" validate:"required,min=1"`
    Total        string    `json:"total" validate:"required,regexp=^\\d+\\.\\d{2}$"`
    Score        *int64    `json:"score,omitempty"`
}

func NewReceipt(retailer string, purchaseDate, purchaseTime time.Time, items []Item, total string) Receipt {
    return Receipt{
        ID:           uuid.New(),
        Retailer:     retailer,
        PurchaseDate: purchaseDate,
        PurchaseTime: purchaseTime,
        Items:        items,
        Total:        total,
    }
}

// Custom unmarshaling for Receipt
// allows string -> time.Time conversion
func (r *Receipt) UnmarshalJSON(data []byte) error {
    type Alias Receipt
    aux := &struct {
        PurchaseDate string `json:"purchaseDate"`
        PurchaseTime string `json:"purchaseTime"`
        *Alias
    }{
        Alias: (*Alias)(r),
    }
    if err := json.Unmarshal(data, aux); err != nil {
        return err
    }
    pd, err := time.Parse("2006-01-02", aux.PurchaseDate)
    if err != nil {
        return err
    }
    pt, err := time.Parse("15:04", aux.PurchaseTime)
    if err != nil {
        return err
    }
    r.PurchaseDate = pd
    r.PurchaseTime = pt

	// Ensure a new UUID is generated if not provided
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}

	// Ensure each item has a new UUID if not provided
	for i := range r.Items {
		if r.Items[i].ID == uuid.Nil {
			r.Items[i].ID = uuid.New()
		}
	}

    return nil
}

// Custom marshaling for Receipt
// allows string -> time.Time conversion
func (r Receipt) MarshalJSON() ([]byte, error) {
    type Alias Receipt
    return json.Marshal(&struct {
        PurchaseDate string `json:"purchaseDate"`
        PurchaseTime string `json:"purchaseTime"`
        Alias
    }{
        PurchaseDate: r.PurchaseDate.Format("2006-01-02"),
        PurchaseTime: r.PurchaseTime.Format("15:04"),
        Alias:        (Alias)(r),
    })
}

type Item struct {
    ID               uuid.UUID `json:"id"`
    ShortDescription string    `json:"shortDescription" validate:"required,regexp=^[\\w\\s\\-]+$"`
    Price            string    `json:"price" validate:"required,regexp=^\\d+\\.\\d{2}$"`
    ReceiptID        uuid.UUID `json:"receiptId,omitempty"`
}

func NewItem(shortDescription, price string) Item {
    return Item{
        ID:               uuid.New(),
        ShortDescription: shortDescription,
        Price:            price,
    }
}