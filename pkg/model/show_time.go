package model

import (
	"time"
)

type Showtime struct {
	BaseModel
	StartTime      time.Time `json:"start_time" gorm:"start_time;not null"`
	EndTime        time.Time `json:"end_time" gorm:"end_time;not null"`
	MovieTheaterId string    `json:"movie_theater_id" gorm:"movie_theater_id;char(36);not null"`
	Day            time.Time `json:"day" gorm:"day;not null;type:date"`
	SkuId          string    `json:"sku_id" gorm:"sku_id;char(36);not null"`
}

func (Showtime) TableName() string {
	return "showtime"
}

type ShowtimeRequest struct {
	ID             *string    `json:"id"`
	StartTime      *time.Time `json:"start_time"`
	EndTime        *time.Time `json:"end_time"`
	RoomID         *string    `json:"room_id"`
	MovieTheaterId *string    `json:"movie_theater_id"`
	Day            *time.Time `json:"day"`
	SkuId          *string    `json:"sku_id"`
}

type ShowtimeParams struct {
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
