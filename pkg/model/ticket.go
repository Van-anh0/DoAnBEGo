package model

import "github.com/google/uuid"

type Ticket struct {
	BaseModel
	OrderId uuid.UUID `json:"order_id" gorm:"type:uuid;not null"`
	SeatId  uuid.UUID `json:"seat_id" gorm:"type:uuid;not null"`
}

func (Ticket) TableName() string {
	return "ticket"
}

type TicketRequest struct {
}

type TicketResponse struct {
}
