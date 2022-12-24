package model

type CategoryHasProduct struct {
	BaseModel
	ProductId  string `json:"product_id" gorm:"not null;char(36)"`
	CategoryId string `json:"category_id" gorm:"not null;char(36)"`
}

func (CategoryHasProduct) TableName() string {
	return "category_has_product"
}

type CategoryHasProductRequest struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type CategoryHasProductParams struct {
	BaseParam
}

type CategoryHasProductResponse struct {
	Data []CategoryHasProduct   `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
