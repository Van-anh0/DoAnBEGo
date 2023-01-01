package model

type Room struct {
	BaseModel
	CinemaId string `json:"cinema_id" gorm:"cinema_id;type:char(36);not null"`
	Name     string `json:"name" gorm:"name;type:varchar(250);not null"`
	Status   string `json:"status" gorm:"status;type:varchar(100);not null"` // active, inactive
}

func (Room) TableName() string {
	return "room"
}

type RoomRequest struct {
	ID       *string `json:"id"`
	Name     *string `json:"name"`
	CinemaId *string `json:"cinema_id"`
	Status   *string `json:"status"`
}

type RoomParams struct {
	BaseParam
}

type RoomResponse struct {
	Data []Room                 `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
