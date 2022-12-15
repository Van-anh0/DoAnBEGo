package model

type Seat struct {
	BaseModel
	Status string  `json:"status" gorm:"status;type:varchar(100);not null"`
	Type   string  `json:"type" gorm:"type;type:varchar(100);not null"`
	Price  float64 `json:"price" gorm:"price;not null;default:0"`
	RoomID string  `json:"room_id" gorm:"room_id;char(36);not null"`
}

func (Seat) TableName() string {
	return "seat"
}

type SeatRequest struct {
	ID     *string  `json:"id"`
	Status *string  `json:"status"`
	Type   *string  `json:"type"`
	Price  *float64 `json:"price"`
	RoomID *string  `json:"room_id"`
}

type SeatParams struct {
	BaseParam
}

type SeatResponse struct {
	Data []Seat                 `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
