package model

type Comment struct {
	BaseModel
	ProductId string `json:"product_id" gorm:"type:char(36);not null"`
	UserId    string `json:"user_id" gorm:"type:char(36);not null"`
	Comment   string `json:"comment" gorm:"type:text;not null"`
}

func (Comment) TableName() string {
	return "comment"
}

type CommentRequest struct {
	ID        *string `json:"id"`
	ProductId *string `json:"product_id"`
	UserId    *string `json:"user_id"`
	Comment   *string `json:"comment"`
}

type CommentParams struct {
	BaseParam
}

type CommentResponse struct {
	Data []Comment              `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
