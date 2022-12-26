package model

type Product struct {
	BaseModel
	Type          string  `json:"type" gorm:"type:varchar(255);not null;index"`
	Status        string  `json:"status" gorm:"type:varchar(255);index"`
	Name          string  `gorm:"type:varchar(250); name;index" json:"name"`
	TotalQuantity float64 `gorm:"type:numeric(10,2); total_quantity" json:"total_quantity"`
	TotalSold     float64 `gorm:"type:numeric(10,2); total_sold" json:"total_sold"`
	Sku           []Sku   `json:"sku" gorm:"foreignKey:product_id;association_foreignkey:id"`
}

func (Product) TableName() string {
	return "product"
}

type ProductRequest struct {
	ID            *string      `json:"id"`
	Name          *string      `json:"name"`
	Type          *string      `json:"type"`
	Status        *string      `json:"status"`
	TotalQuantity *float64     `json:"total_quantity"`
	TotalSold     *float64     `json:"total_sold"`
	Sku           []SkuRequest `json:"sku"`
}

type ProductResponse struct {
	*Product
	Sku []*Sku `json:"sku"`
}

type ProductParams struct {
	BaseParam
	Day            string `json:"day" form:"day"`
	MovieTheaterId string `json:"movie_theater_id" form:"movie_theater_id"`
}

type ListProductResponse struct {
	Data []Product              `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
