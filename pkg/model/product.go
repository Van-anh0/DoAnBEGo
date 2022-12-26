package model

type Product struct {
	BaseModel
	Type     string  `gorm:"column:Type;type:varchar(255);not null"` // movie | event | ticket | product
	Name     string  `gorm:"type:varchar(250); name;index" json:"name"`
	Price    float64 `json:"price" gorm:"type:float"`
	Cost     float64 `json:"cost" gorm:"type:float"`
	Quantity float64 `json:"quantity" gorm:"type:float"`
	Image    string  `json:"image"`
}

func (Product) TableName() string {
	return "product"
}

type ProductRequest struct {
	ID       *string  `json:"id"`
	Name     *string  `json:"name"`
	Price    *float64 `json:"price"`
	Cost     *float64 `json:"cost"`
	Quantity *float64 `json:"quantity"`
	Image    *string  `json:"image"`
}

type ProductParams struct {
	BaseParam
}

type ProductResponse struct {
	Data []Product              `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
