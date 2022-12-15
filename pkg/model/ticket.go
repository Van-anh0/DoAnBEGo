package model

type Ticket struct {
	BaseModel
	OrderId string `json:"order_id" gorm:"not null;type:char(36);"`
	SeatId  string `json:"seat_id" gorm:"not null;type:char(36);"`
}

func (Ticket) TableName() string {
	return "ticket"
}

type TicketRequest struct {
	ID      *string `json:"id"`
	OrderId *string `json:"order_id"`
	SeatId  *string `json:"seat_id"`
}

type TicketParams struct {
	BaseParam
}

type TicketResponse struct {
	Data []Ticket               `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
