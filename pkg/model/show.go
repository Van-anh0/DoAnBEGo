package model

import (
	"time"
)

// thoi gian chieu phim
type Showtime struct {
	BaseModel
	Showtime  time.Time `json:"showtime" gorm:"showtime;not null"`                  // thoi gian chieu phim
	CinemaId  string    `json:"cinema_id" gorm:"cinema_id;char(36);not null"`       // id cua rap
	MovieId   string    `json:"movie_id" gorm:"movie_id;char(36);not null"`         // id cua phim
	RoomId    string    `json:"room_id" gorm:"room_id;char(36);not null"`           // id cua phong
	RoomName  string    `json:"room_name" gorm:"room_name;varchar(255);not null"`   // ten phong
	MovieName string    `json:"movie_name" gorm:"movie_name;varchar(255);not null"` // ten phim
}

func (Showtime) TableName() string {
	return "showtime"
}

type ShowtimeRequest struct {
	ID           *string     `json:"id"`
	Showtime     *time.Time  `json:"showtime"`
	CinemaId     *string     `json:"cinema_id"`
	MovieId      *string     `json:"movie_id"`
	RoomId       *string     `json:"room_id"`
	RoomName     *string     `json:"room_name"`
	MovieName    *string     `json:"movie_name"`
	ListShowtime []time.Time `json:"list_showtime"`
}

type ShowParams struct {
	BaseParam
	ListRoomId string `json:"list_room_id" form:"list_room_id"`
	MovieId    string `json:"movie_id" form:"movie_id"`
	CinemaId   string `json:"cinema_id" form:"cinema_id"`
}

type ShowtimeResponse struct {
	Data []Showtime             `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type ShowtimeGroupResponse struct {
	Data map[string][]Showtime  `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type ShowtimeGroupNestedResponse struct {
	Data map[string]map[string][]Showtime `json:"data"`
	Meta map[string]interface{}           `json:"meta"`
}
