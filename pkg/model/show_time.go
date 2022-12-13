package model

import (
	"time"
)

type Showtime struct {
	BaseModel
	StartTime      time.Time `json:"start_time" gorm:"start_time;not null"`
	EndTime        time.Time `json:"end_time" gorm:"end_time;not null"`
	RoomID         string    `json:"room_id" gorm:"room_id;char(36);not null"`
	MovieTheaterID string    `json:"movie_theater_id" gorm:"movie_theater_id;char(36);not null"`
}

func (Showtime) TableName() string {
	return "show_time"
}

type ShowtimeRequest struct {
}

type ShowtimeResponse struct {
}
