package model

import "encoding/json"

type Sku struct {
	BaseModel
	ProductId    string          `json:"product_id" gorm:"type:char(36);not null"`
	Attribute    json.RawMessage `json:"attribute" gorm:"type:json"`
	Cost         float64         `json:"cost" gorm:"type:numeric(10,2);not null"`
	Quantity     float64         `json:"quantity" gorm:"type:numeric(10,2);not null"`
	SoldQuantity float64         `json:"sold_quantity" gorm:"type:numeric(10,2);not null"`
	Image        string          `json:"image" gorm:"type:varchar(255)"`
	Type         string          `json:"type" gorm:"type:varchar(255)"`
	Showtime     []Showtime      `json:"showtime" gorm:"foreignKey:sku_id;association_foreignkey:id"`
}

func (Sku) TableName() string {
	return "sku"
}

type SkuRequest struct {
	ID           *string          `json:"id"`
	ProductId    *string          `json:"product_id"`
	Attribute    *json.RawMessage `json:"attribute"`
	Cost         *float64         `json:"cost"`
	Quantity     *float64         `json:"quantity"`
	SoldQuantity *float64         `json:"sold_quantity"`
	Image        *string          `json:"image"`
}

type SkuParams struct {
	BaseParam
	MovieTheaterId string `json:"movie_theater_id" form:"movie_theater_id"`
	Day            string `json:"day" form:"day"`
}

type SkuResponse struct {
	Data []Sku                  `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
