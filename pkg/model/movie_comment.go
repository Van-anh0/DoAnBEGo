package model

type MovieComment struct {
	BaseModel
	MovieId string `json:"movie_id" gorm:"type:char(36);not null"`
	UserId  string `json:"user_id" gorm:"type:char(36);not null"`
	Comment string `json:"comment" gorm:"type:text;not null"`
}

func (MovieComment) TableName() string {
	return "movie_comment"
}

type MovieCommentRequest struct {
	ID      *string `json:"id"`
	MovieId *string `json:"movie_id"`
	UserId  *string `json:"user_id"`
	Comment *string `json:"comment"`
}

type MovieCommentParams struct {
	BaseParam
}

type MovieCommentResponse struct {
	Data []MovieComment         `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
