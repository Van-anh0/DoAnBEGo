package model

type User struct {
	BaseModel
}

func (User) TableName() string {
	return "users"
}

type UserRequest struct {
}

type UserParam struct {
}
