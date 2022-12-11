package model

import (
	"github.com/google/uuid"
	"time"
)

type Showtime struct {
	BaseModel
	StartTime      time.Time `json:"start_time" gorm:"start_time;not null"`
	EndTime        time.Time `json:"end_time" gorm:"end_time;not null"`
	RoomID         uuid.UUID `json:"room_id" gorm:"room_id;type:uuid;not null"`
	MovieTheaterID uuid.UUID `json:"movie_theater_id" gorm:"movie_theater_id;type:uuid;not null"`
}

func (Showtime) TableName() string {
	return "show_time"
}

type ShowtimeRequest struct {
}

type ShowtimeResponse struct {
}
