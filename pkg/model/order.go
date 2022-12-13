package model

type Order struct {
	BaseModel
	TotalPrice    float64  `json:"total_price" gorm:"type:float;not null"`
	Status        string   `json:"status" gorm:"type:varchar(100);not null"`
	PaymentMethod string   `json:"payment_method" gorm:"type:varchar(100);not null"`
	UserID        string   `json:"user_id" gorm:"not null;char(36)"`
	MovieId       string   `json:"movie_id" gorm:"not null;char(36)"`
	SlotId        string   `json:"slot_id" gorm:"not null;char(36)"`
	Ticket        []Ticket `json:"ticket" gorm:"foreignKey:order_id;references:id"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderRequest struct {
}

type OrderResponse struct {
}
