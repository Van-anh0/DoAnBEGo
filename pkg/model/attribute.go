package model

type Attribute struct {
	BaseModel
	ProductId   string `json:"product_id" gorm:"type:char(36);not null"`
	ObjectValue string `json:"object_value" gorm:"type:jsonb;not null"`
}

func (Attribute) TableName() string {
	return "attribute"
}

type AttributeRequest struct {
	ID          *string `json:"id"`
	ProductId   *string `json:"product_id"`
	ObjectValue *string `json:"object_value"`
}

type AttributeParams struct {
	BaseParam
}

type AttributeResponse struct {
	Data []Attribute            `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
