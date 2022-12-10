package model

type Metadata struct {
	BaseModel
	Key      string  `json:"key" gorm:"type:varchar(100);not null"`
	Type     string  `json:"type" gorm:"type:varchar(250);not null"`
	Name     string  `json:"name" gorm:"type:varchar(250);not null"`
	Priority float64 `json:"priority" gorm:"priority;not null;default:0"`
}

func (Metadata) TableName() string {
	return "metadata"
}
