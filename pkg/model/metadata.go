package model

type Metadata struct {
	BaseModel
	Key      string  `json:"key" gorm:"type:varchar(100);not null"`
	Type     string  `json:"type" gorm:"type:varchar(250);not null"`
	Name     string  `json:"name" gorm:"type:varchar(250);not null"`
	Position float64 `json:"position" gorm:"position;not null;default:0"`
}

func (Metadata) TableName() string {
	return "metadata"
}

type MetadataRequest struct {
	ID       *string  `json:"id"`
	Key      *string  `json:"key"`
	Type     *string  `json:"type"`
	Name     *string  `json:"name"`
	Position *float64 `json:"position"`
}

type MetadataParams struct {
	BaseParam
}

type MetadataResponse struct {
	Data []Metadata             `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
