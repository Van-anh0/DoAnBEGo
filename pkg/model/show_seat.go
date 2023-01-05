package model

type ShowSeat struct {
	BaseModel
	OrderId    string  `json:"order_id" gorm:"not null;type:char(36);"`
	SeatId     string  `json:"seat_id" gorm:"not null;type:char(36);"`
	ShowtimeId string  `json:"showtime_id" gorm:"not null;type:char(36);"`
	Price      float64 `json:"price" gorm:"type:float;not null"`
	Row        string  `json:"row" gorm:"type:varchar(100);not null"`
	Col        int     `json:"col" gorm:"type:int;not null"`
}

func (ShowSeat) TableName() string {
	return "show_seat"
}

type ShowSeatRequest struct {
	ID         *string  `json:"id"`
	OrderId    *string  `json:"order_id"`
	SeatId     *string  `json:"seat_id"`
	ShowtimeId *string  `json:"showtime_id"`
	Price      *float64 `json:"price"`
	Row        *string  `json:"row"`
	Col        *int     `json:"col"`
}

type ShowSeatParams struct {
	BaseParam
	SeatId     []string `json:"seat_id"`
	ShowtimeId string   `json:"showtime_id" form:"showtime_id"`
}

type ShowSeatResponse struct {
	Data []ShowSeat             `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
