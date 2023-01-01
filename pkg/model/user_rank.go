package model

type UserRank struct {
	BaseModel
	Name        string  `json:"name" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:varchar(255);not null"`
	MinPoint    int     `json:"min_point" gorm:"type:int;not null"`
	Discount    float64 `json:"discount" gorm:"type:float;not null"`
}

func (UserRank) TableName() string {
	return "user_rank"
}

type UserRankRequest struct {
	ID          *string  `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	MinPoint    *int     `json:"min_point"`
	Discount    *float64 `json:"discount"`
}

type RankParams struct {
	BaseParam
}

type RankResponse struct {
	Data []UserRank             `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
