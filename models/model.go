package models

import _ "encoding/json"

type Orders struct {
	OrderUid          string   `json:"order_uid"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	Delivery          Delivery `json:"delivery"`
	Payment           Payment  `json:"payment"`
	Items             []Items  `json:"items"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerId        string   `json:"customer_id"`
	DeliveryService   string   `json:"delivery_service"`
	Shardkey          string   `json:"shardkey"`
	SmId              int64    `json:"sm_id"`
	DateCreated       string   `json:"date_created"`
	OofShard          string   `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int64  `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int64  `json:"delivery_cost"`
	GoodsTotal   int64  `json:"goods_total"`
	CustomFee    int64  `json:"custom_fee"`
}

type Items struct {
	ChrtId      int64  `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int64  `json:"price"`
	rid         string `json:"rid"`
	name        string `json:"name"`
	sale        int64  `json:"sale"`
	size        string `json:"size"`
	totalPrice  int64  `json:"total_price"`
	nmId        int64  `json:"nm_id"`
	brand       string `json:"brand"`
	Status      int64  `json:"status"`
}
