package model

import "github.com/google/uuid"

type Order struct {
	BaseModel
	TotalPrice    float64   `json:"total_price" gorm:"type:float;not null"`
	Status        string    `json:"status" gorm:"type:varchar(100);not null"`
	PaymentMethod string    `json:"payment_method" gorm:"type:varchar(100);not null"`
	UserID        uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	MovieId       uuid.UUID `json:"movie_id" gorm:"type:uuid;not null"`
	SlotId        uuid.UUID `json:"slot_id" gorm:"type:uuid;not null"`
	Ticket        []Ticket  `json:"ticket" gorm:"foreignKey:order_id"`
}

func (Order) TableName() string {
	return "order"
}

type OrderRequest struct {
}

type OrderResponse struct {
}
