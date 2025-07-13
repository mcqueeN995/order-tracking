package models

import (
	"encoding/json"
	"errors"
	"strings"
)

type Order struct {
	OrderUID        string   `json:"order_uid" db:"order_uid"`
	TrackNumber     string   `json:"track_number" db:"track_number"`
	Entry           string   `json:"entry" db:"entry"`
	Delivery        Delivery `json:"delivery" db:"delivery"`
	Payment         Payment  `json:"payment" db:"payment"`
	Items           []Items  `json:"items" db:"items"`
	locale          string   `json:"locale" db:"locale"`
	CustomerId      string   `json:"customer_id" db:"customer_id"`
	DeliveryService string   `json:"delivery_service" db:"delivery_service"`
	Shardkey        string   `json:"shardkey" db:"shardkey"`
	SmID            int      `json:"sm_id" db:"sm_id"`
	DateCreated     string   `json:"date_created" db:"date_created"`
	OofShard        string   `json:"oof_shard" db:"oof_shard"`
}

type Delivery struct {
	OrderUID string `json:"order_uid" db:"order_uid"`
	Name     string `json:"name" db:"name"`
	Phone    string `json:"phone" db:"phone"`
	Zip      string `json:"zip" db:"zip"`
	City     string `json:"city" db:"city"`
	Address  string `json:"address" db:"address"`
	Region   string `json:"region" db:"region"`
	Email    string `json:"email" db:"email"`
}

type Payment struct {
	OrderUID     string `json:"order_uid" db:"order_uid"`
	Transaction  string `json:"transation" db:"transation"`
	RequestID    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int    `json:"amount" db:"amount"`
	PaymentDt    int64  `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee"`
}

type Items struct {
	OrderUID    string `json:"order_uid" db:"order_uid"`
	ChrtID      int    `json:"chrt_id" db:"chrt_id"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Price       int    `json:"price" db:"price"`
	Rid         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int    `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  int    `json:"total_price" db:"total_price"`
	NmID        int    `json:"nm_id" db:"nm_id"`
	Brand       string `json:"brand" db:"brand"`
	Status      int    `json:"status" db:"status"`
}

func (order *Order) Validate() error {
	if order.OrderUID == "" {
		return errors.New("Order UID is required")
	}
	if order.TrackNumber == "" {
		return errors.New("Track number is required")
	}
	if order.Payment.Transaction == "" {
		return errors.New("Transaction is required")
	}
	if len(order.Items) == 0 {
		return errors.New("Items is required")
	}
	if order.Payment.Transaction != order.OrderUID {
		return errors.New("payment transaction must match order_uid")
	}
	if order.Delivery.Email != "" {
		if !strings.Contains(order.Delivery.Email, "@") {
			return errors.New("invalid email format")
		}
	}
	return nil
}

func (order *Order) ToJSON() ([]byte, error) {
	return json.Marshal(order)
}

func (order *Order) FromJSON(data []byte) error {
	return json.Unmarshal(data, &order)
}
