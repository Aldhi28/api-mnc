package model

import (
	"time"
)

// request dari body api

type Bill struct {
	Id         string       `json:"id"`
	BillDate   time.Time    `json:"bill_date"`
	EntryDate  time.Time    `json:"entry_date"`
	EmployeeId string       `json:"employee_id"`
	CustomerId string       `json:"customer_id"`
	BillDetail []BillDetail `json:"bill_detail"`
}

type BillDetail struct {
	Id           string    `json:"id"`
	BillId       string    `json:"bill_id"`
	ProductId    string    `json:"product_id"`
	ProductPrice int       `json:"product_price"`
	Qty          int       `json:"quantity"`
	FinishDate   time.Time `json:"finish_date"`
}
