package model

type Category struct {
	BaseModel
	Name     string  `gorm:"type:varchar(250); name;index" json:"name"`
	Position float64 `json:"position" gorm:"type:float"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryRequest struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type CategoryParams struct {
	BaseParam
}

type CategoryResponse struct {
	Data []Category             `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
