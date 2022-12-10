package model

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"name;type:varchar(100)"`
	Avatar   string `json:"avatar" gorm:"avatar"`
	Email    string `json:"email" gorm:"email;type:varchar(250);not null;unique"`
	Password string `json:"password" gorm:"password;type:varchar(250);not null"`
	Gender   string `json:"gender" gorm:"column:gender"`
	Phone    string `json:"phone" gorm:"phone;type:varchar(20)"`
	Role     int    `json:"role" gorm:"role"`
	Address  string `json:"address" gorm:"address;type:varchar(250)"`
	IdCard   string `json:"id_card" gorm:"id_card;type:varchar(20)"`
	IsActive bool   `json:"is_active" gorm:"is_active"`
}

func (User) TableName() string {
	return "users"
}

type UserRequest struct {
}

type UserResponse struct {
}
