package model

type MovieTheater struct {
	BaseModel
	Name    string `json:"name" gorm:"name;type:varchar(250);not null"`
	Address string `json:"address" gorm:"address;type:varchar(250);not null"`
	Phone   string `json:"phone" gorm:"phone;type:varchar(20);not null"`
	Status  string `json:"status" gorm:"status;type:varchar(100);not null"`
}

func (MovieTheater) TableName() string {
	return "movie_theater"
}

type MovieTheaterRequest struct {
}

type MovieTheaterResponse struct {
}
