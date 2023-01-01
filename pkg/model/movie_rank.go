package model

type MovieRank struct {
	BaseModel
	UserId  string  `json:"user_id" gorm:"type:char(36);not null"`
	MovieId string  `json:"movie_id" gorm:"type:char(36);not null"`
	Point   float64 `json:"point" gorm:"type:float;default:0"` // [1-10]
}

func (MovieRank) TableName() string {
	return "movie_rank"
}

type MovieRankRequest struct {
	ID      *string  `json:"id"`
	UserId  *string  `json:"user_id"`
	MovieId *string  `json:"movie_id"`
	Point   *float64 `json:"point"` // [1-10]
}

type MovieRankParams struct {
	BaseParam
}

type MovieRankResponse struct {
	Data []MovieRank            `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
