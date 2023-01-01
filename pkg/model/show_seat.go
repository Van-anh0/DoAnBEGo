package model

type ShowSeat struct {
	BaseModel
	OrderId string  `json:"order_id" gorm:"not null;type:char(36);"`
	SeatId  string  `json:"seat_id" gorm:"not null;type:char(36);"`
	ShowId  string  `json:"show_id" gorm:"not null;type:char(36);"`
	Price   float64 `json:"price" gorm:"type:float;not null"`
}

func (ShowSeat) TableName() string {
	return "show_seat"
}

type ShowSeatRequest struct {
	ID      *string  `json:"id"`
	OrderId *string  `json:"order_id"`
	SeatId  *string  `json:"seat_id"`
	ShowId  *string  `json:"show_id"`
	Price   *float64 `json:"price"`
}

type ShowSeatParams struct {
	BaseParam
}

type ShowSeatResponse struct {
	Data []ShowSeat             `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
