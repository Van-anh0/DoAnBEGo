package model

type Product struct {
	BaseModel
	Name     string  `json:"name" gorm:"type:varchar(250); name;index" `
	Quantity float64 `json:"quantity" gorm:"type:float;default:0"`
	Uom      string  `json:"uom"` // don vi
	Price    float64 `json:"price" gorm:"type:float;default:0"`
	Sold     float64 `json:"sold" gorm:"type:float;default:0"`
}

func (Product) TableName() string {
	return "product"
}

type ProductRequest struct {
	ID       *string  `json:"id"`
	Name     *string  `json:"name"`
	Quantity *float64 `json:"quantity"`
	Uom      *string  `json:"uom"` // don vi
	Price    *float64 `json:"price"`
	Sold     *float64 `json:"sold"`
}

type ProductResponse struct {
	*Product
}

type ProductParams struct {
	BaseParam
}

type ListProductResponse struct {
	Data []Product              `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
