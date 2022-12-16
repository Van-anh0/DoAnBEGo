package model

type Order struct {
	BaseModel
	TotalPrice    float64  `json:"total_price" gorm:"type:float;not null"`
	Status        string   `json:"status" gorm:"type:varchar(100);not null"`
	PaymentMethod string   `json:"payment_method" gorm:"type:varchar(100);not null"`
	UserId        string   `json:"user_id" gorm:"not null;char(36)"`
	MovieId       string   `json:"movie_id" gorm:"not null;char(36)"`
	ShowtimeId    string   `json:"showtime_id" gorm:"not null;char(36)"`
	Ticket        []Ticket `json:"ticket" gorm:"foreignKey:order_id;references:id"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderRequest struct {
	ID            *string   `json:"id"`
	TotalPrice    *float64  `json:"total_price"`
	Status        *string   `json:"status"`
	PaymentMethod *string   `json:"payment_method"`
	UserID        *string   `json:"user_id"`
	MovieId       *string   `json:"movie_id"`
	ShowtimeId    *string   `json:"showtime_id"`
	Ticket        *[]Ticket `json:"ticket"`
}

type OrderParams struct {
	BaseParam
}

type OrderResponse struct {
	Data []Order                `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
