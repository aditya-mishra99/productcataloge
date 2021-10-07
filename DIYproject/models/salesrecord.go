package models

import (
	"time"
)

type SalesRecord struct {
	Id           int       `json:"id"`
	ProductId    int       `json:"productid"`
	QuantitySold int       `json:"quantitysold"`
	SalesTime    time.Time `json:"salestime"`
}
