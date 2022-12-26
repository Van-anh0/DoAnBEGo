package model

type OrderItem struct {
	BaseModel
	OrderId         string  `json:"order_id" gorm:"not null;type:char(36);"`
	ShowtimeId      string  `json:"showtime_id" gorm:"not null;char(36)"`
	ProductId       string  `json:"product_id" gorm:"not null;char(36)"`
	ProductQuantity float64 `json:"product_quantity" gorm:"not null;type:float;"`
	ProductPrice    float64 `json:"product_price" gorm:"not null;type:float;"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

type TicketRequest struct {
	ID      *string  `json:"id"`
	OrderId *string  `json:"order_id"`
	SeatId  *string  `json:"seat_id"`
	Price   *float64 `json:"price"`
}

type TicketParams struct {
	BaseParam
}

type TicketResponse struct {
	Data []OrderItem            `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
