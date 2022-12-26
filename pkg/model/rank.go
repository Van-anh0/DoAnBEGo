package model

type Rank struct {
	BaseModel
	Name        string  `json:"name" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:varchar(255);not null"`
	MinPoint    int     `json:"min_point" gorm:"type:int;not null"`
	MaxPoint    int     `json:"max_point" gorm:"type:int;not null"`
	IsActive    bool    `json:"is_active" gorm:"type:boolean;not null"`
	Position    float64 `json:"position" gorm:"type:float;not null"`
}

func (Rank) TableName() string {
	return "rank"
}

type RankRequest struct {
	ID          *string  `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	MinPoint    *int     `json:"min_point"`
	MaxPoint    *int     `json:"max_point"`
	IsActive    *bool    `json:"is_active"`
	Position    *float64 `json:"position"`
}

type RankParams struct {
	BaseParam
}

type RankResponse struct {
	Data []Rank                 `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
