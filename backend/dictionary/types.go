package dictionary

import (
	"time"
)

type Product struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	ShopName     string  `json:"shop_name"`
	ProductPrice float64 `json:"product_price"`
	ImageURL     string  `json:"image_url"`
}


type APIResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error_message"`
}

type Banner struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	ImgSrc string `json:"imgSrc"`
	StartDate time.Time `json:"startDate"`
	EndDate time.Time `json:"endDate"`
	MinTier string `json:"minTier"`
	MaxTier string `json:"maxTier"`
	MinBalance int64 `json:"minBalance"`
	MaxBalance int64 `json:"maxBalance"`
	MinTokopoint int64 `json:"minTokopoint"`
	MaxTokopoint int64 `json:"maxTokopoint"`
	isActive bool `json:"isActive"`
}

type User struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Balance int64 `json:"balance"`
	Tokopoint int64 `json:"tokopoint"`
	Tier string `json:"tier"`
	BannerList []Banner `json:"banner"`
	Location string `json:"email"`
}