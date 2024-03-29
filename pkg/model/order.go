package model

type Order struct {
	BaseModel
	Total         float64     `json:"total" gorm:"type:float;not null"`
	Status        string      `json:"status" gorm:"type:varchar(100);not null"` // chua nhan ve, da nhan ve, da huy
	PaymentMethod string      `json:"payment_method" gorm:"type:varchar(100);not null"`
	UserId        string      `json:"user_id" gorm:"not null;char(36)"`
	OrderNumber   string      `json:"order_number" gorm:"type:varchar(100);not null"`
	OrderItem     []OrderItem `json:"order_item"`
	ShowSeat      []ShowSeat  `json:"show_seat"`
	MovieName     string      `json:"movie_name" gorm:"type:varchar(100);not null"`
	MovieImage    string      `json:"movie_image" gorm:"type:varchar(255);not null"`
	RoomName      string      `json:"room_name" gorm:"type:varchar(100);not null"`
	Showtime      string      `json:"showtime" gorm:"type:varchar(100);not null"`
	CinemaName    string      `json:"cinema_name" gorm:"type:varchar(100);not null"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderRequest struct {
	ID            *string      `json:"id"`
	Total         *float64     `json:"total"`
	Status        *string      `json:"status"`
	PaymentMethod *string      `json:"payment_method"`
	UserID        *string      `json:"user_id" valid:"Required"`
	ShowtimeId    *string      `json:"showtime_id" valid:"Required"`
	OrderItem     *[]OrderItem `json:"order_item" valid:"Required"`
	ShowSeat      *[]ShowSeat  `json:"show_seat" valid:"Required"`
	MovieName     *string      `json:"movie_name"`
	MovieImage    *string      `json:"movie_image"`
	RoomName      *string      `json:"room_name"`
	Showtime      *string      `json:"showtime"`
	CinemaName    *string      `json:"cinema_name"`
}

type OrderParams struct {
	BaseParam
}

type OrderResponse struct {
	Data []Order                `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
