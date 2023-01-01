package model

import "time"

type Promotion struct {
	BaseModel
	Key          string    `json:"key" gorm:"type:varchar(255);not null"`
	StartTime    time.Time `json:"start_date" gorm:"not null"`
	EndTime      time.Time `json:"end_date" gorm:"not null"`
	IsActive     bool      `json:"is_active" gorm:"type:boolean;not null"`
	Type         string    `json:"type" gorm:"type:varchar(255);not null"`   // percent, amount
	MinPrice     float64   `json:"min_price" gorm:"type:float;not null"`     // ap dung cho don hang co gia tri tu min_price
	MaxPromotion float64   `json:"max_promotion" gorm:"type:float;not null"` // so tien toi da duoc giam gia
	MaxUser      int       `json:"max_user" gorm:"type:int;default:-1"`      // so luong nguoi toi da duoc ap dung
}

func (Promotion) TableName() string {
	return "promotion"
}

type PromotionRequest struct {
	ID           *string    `json:"id"`
	Key          *string    `json:"key"`
	StartTime    *time.Time `json:"start_date"`
	EndTime      *time.Time `json:"end_date"`
	IsActive     *bool      `json:"is_active"`
	Type         *string    `json:"type"`
	MinPrice     *float64   `json:"min_price"`     // ap dung cho don hang co gia tri tu min_price
	MaxPromotion *float64   `json:"max_promotion"` // so tien toi da duoc giam gia
	MaxUser      *int       `json:"max_user"`      // so luong nguoi toi da duoc ap dung
}

type PromotionParams struct {
	BaseParam
}

type PromotionResponse struct {
	Data []Promotion            `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
