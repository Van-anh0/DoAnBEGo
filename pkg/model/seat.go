package model

import "github.com/google/uuid"

type Seat struct {
	BaseModel
	Status string    `json:"status" gorm:"status;type:varchar(100);not null"`
	Type   string    `json:"type" gorm:"type;type:varchar(100);not null"`
	Price  float64   `json:"price" gorm:"price;not null;default:0"`
	RoomID uuid.UUID `json:"room_id" gorm:"room_id;type:uuid;not null"`
}

func (Seat) TableName() string {
	return "seat"
}

type SeatRequest struct {
}

type SeatResponse struct {
}
