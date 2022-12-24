package model

import "time"

type Promotion struct {
	BaseModel
	Name         string    `json:"name" gorm:"type:varchar(255);not null"`
	StartTime    time.Time `json:"start_date" gorm:"not null"`
	EndTime      time.Time `json:"end_date" gorm:"not null"`
	IsActive     bool      `json:"is_active" gorm:"type:boolean;not null"`
	Position     float64   `json:"position" gorm:"type:float;not null"`
	Type         string    `json:"type" gorm:"type:varchar(255);not null"`   // percent, amount
	MinPrice     float64   `json:"min_price" gorm:"type:float;not null"`     // ap dung cho don hang co gia tri tu min_price
	MaxPromotion float64   `json:"max_promotion" gorm:"type:float;not null"` // so tien toi da duoc giam gia
	MaxUser      int       `json:"max_user" gorm:"type:int;not null"`        // so luong nguoi toi da duoc ap dung
}

func (Promotion) TableName() string {
	return "promotion"
}

type PromotionRequest struct {
	ID       *string  `json:"id"`
	Name     *string  `json:"name"`
	Price    *float64 `json:"price"`
	Cost     *float64 `json:"cost"`
	Quantity *float64 `json:"quantity"`
	Image    *string  `json:"image"`
}

type PromotionParams struct {
	BaseParam
}

type PromotionResponse struct {
	Data []Promotion            `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
