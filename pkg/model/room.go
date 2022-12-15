package model

type Room struct {
	BaseModel
	Name           string `json:"name" gorm:"name;type:varchar(250);not null"`
	Type           string `json:"type" gorm:"type;type:varchar(250);not null"`
	MovieTheaterID string `json:"movie_theater_id" gorm:"movie_theater_id;type:char(36);not null"`
	Status         string `json:"status" gorm:"status;type:varchar(100);not null"`
}

func (Room) TableName() string {
	return "room"
}

type RoomRequest struct {
	ID             *string `json:"id"`
	Name           *string `json:"name"`
	Type           *string `json:"type"`
	MovieTheaterID *string `json:"movie_theater_id"`
	Status         *string `json:"status"`
}

type RoomParams struct {
	BaseParam
}

type RoomResponse struct {
	Data []Room                 `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
