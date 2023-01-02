package model

import (
	"time"
)

// thoi gian chieu phim
type Showtime struct {
	BaseModel
	Showtime time.Time `json:"showtime" gorm:"showtime;not null"`            // thoi gian chieu phim
	CinemaId string    `json:"cinema_id" gorm:"cinema_id;char(36);not null"` // id cua rap
	MovieId  string    `json:"movie_id" gorm:"movie_id;char(36);not null"`   // id cua phim
}

func (Showtime) TableName() string {
	return "showtime"
}

type ShowtimeRequest struct {
	ID       *string    `json:"id"`
	Showtime *time.Time `json:"showtime"`
	CinemaId *string    `json:"cinema_id"`
	MovieId  *string    `json:"movie_id"`
}

type ShowParams struct {
	BaseParam
}

type ShowtimeResponse struct {
	Data []Showtime             `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type ShowtimeGroupResponse struct {
	Data map[string][]Showtime  `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
