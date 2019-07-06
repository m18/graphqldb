package model

import (
	"time"
)

// Order encapsulates order data
type Order struct {
	ID int32
	CustomerID int32
	Time time.Time
}