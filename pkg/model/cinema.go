package model

// rap chieu phim
type Cinema struct {
	BaseModel
	Name    string `json:"name" gorm:"name;type:varchar(250);not null"`
	Address string `json:"address" gorm:"address;type:varchar(250);not null"`
	Phone   string `json:"phone" gorm:"phone;type:varchar(20);not null"`
	Status  string `json:"status" gorm:"status;type:varchar(100);not null"` // active, inactive
}

func (Cinema) TableName() string {
	return "cinema"
}

type CinemaRequest struct {
	ID      *string `json:"id"`
	Name    *string `json:"name"`
	Address *string `json:"address"`
	Phone   *string `json:"phone"`
	Status  *string `json:"status"`
}

type CinemaParams struct {
	BaseParam
	MovieId string `json:"movie_id" form:"movie_id"`
	Day     string `json:"day" form:"day"`
}

type CinemaResponse struct {
	Data []Cinema               `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
