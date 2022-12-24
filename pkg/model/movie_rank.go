package model

type ProductRank struct {
	BaseModel
	UserId    string  `json:"user_id" gorm:"type:char(36);not null"`
	ProductId string  `json:"product_id" gorm:"type:char(36);not null"`
	Point     float64 `json:"point" gorm:"type:float;default:0"`
}

func (ProductRank) TableName() string {
	return "product_rank"
}

type ProductRankRequest struct {
	ID *string `json:"id"`
}

type ProductRankParams struct {
	BaseParam
}

type ProductRankResponse struct {
	Data []ProductRank          `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
