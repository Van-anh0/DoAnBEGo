package model

type Seat struct {
	BaseModel
	RoomID string  `json:"room_id" gorm:"room_id;char(36);not null"`
	Name   string  `json:"name" gorm:"name;type:varchar(100);not null"`
	Row    string  `json:"row" gorm:"row;type:varchar(10);not null"`
	Col    int     `json:"col" gorm:"col;type:int;not null"`
	Price  float64 `json:"price" gorm:"price;type:float;not null"`
}

func (Seat) TableName() string {
	return "seat"
}

type SeatRequest struct {
	ID     *string  `json:"id"`
	RoomID *string  `json:"room_id"`
	Name   *string  `json:"name"`
	Row    *string  `json:"row"`
	Col    *int     `json:"col"`
	Price  *float64 `json:"price"`
}

type SeatParams struct {
	BaseParam
	ShowtimeId string `json:"showtime_id" form:"showtime_id"`
}

type SeatResponse struct {
	BaseModel
	RoomID string  `json:"room_id"`
	Name   string  `json:"name"`
	Row    string  `json:"row"`
	Col    int     `json:"col"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

type ListSeatResponse struct {
	Data []SeatResponse         `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
