package model

type UserRank struct {
	BaseModel
	UserId      string  `json:"user_id" gorm:"type:char(36);not null"`
	Name        string  `json:"name" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:varchar(255);not null"`
	MinPoint    int     `json:"min_point" gorm:"type:int;not null"`
	MaxPoint    int     `json:"max_point" gorm:"type:int;not null"`
	IsActive    bool    `json:"is_active" gorm:"type:boolean;not null"`
	Position    float64 `json:"position" gorm:"type:float;not null"`
}

func (UserRank) TableName() string {
	return "user_rank"
}

type UserRankRequest struct {
	ID          *string  `json:"id"`
	UserId      *string  `json:"user_id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	MinPoint    *int     `json:"min_point"`
	MaxPoint    *int     `json:"max_point"`
	IsActive    *bool    `json:"is_active"`
	Position    *float64 `json:"position"`
}

type UserRankParams struct {
	BaseParam
}

type UserRankResponse struct {
	Data []UserRank             `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
