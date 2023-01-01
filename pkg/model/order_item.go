package model

type OrderItem struct {
	BaseModel
	OrderId   string  `json:"order_id" gorm:"not null;type:char(36);"`
	ProductId string  `json:"product_id" gorm:"not null;type:char(36);"`
	Price     float64 `json:"price" gorm:"type:float;not null"`
	Quantity  int     `json:"quantity" gorm:"type:int;not null"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

type OrderItemRequest struct {
	ID        *string  `json:"id"`
	OrderId   *string  `json:"order_id"`
	ProductId *string  `json:"product_id"`
	Price     *float64 `json:"price"`
	Quantity  *int     `json:"quantity"`
}

type OrderItemParams struct {
	BaseParam
}

type OrderItemResponse struct {
	Data []OrderItem            `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
