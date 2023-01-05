package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
	"strings"
)

func (r *RepoPG) CreateRoom(ctx context.Context, ob *model.Room) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateRoom(ctx context.Context, ob *model.Room) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Updates(&ob).Error
}

func (r *RepoPG) DeleteRoom(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.Room{}).Error
}

func (r *RepoPG) GetOneRoom(ctx context.Context, id string) (*model.Room, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.Room{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

func (r *RepoPG) GetListRoom(ctx context.Context, req model.RoomParams) (*model.RoomResponse, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.RoomResponse{}
	var err error
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	total := new(struct {
		Count int `json:"count"`
	})

	tx = tx.Table("room")

	if req.Search != "" {
		tx = tx.Where("unaccent(name) ilike %?%", req.Search)
	}

	if req.CinemaId != "" {
		tx = tx.Where("room.cinema_id = ?", req.CinemaId)
	}

	if req.MovieId != "" {
		tx = tx.Joins("inner join showtime on showtime.room_id = room.id").Where("showtime.movie_id = ?", req.MovieId)
	}

	if req.Filter != "" {
		filter := strings.Split(req.Filter, ",")
		for i := 0; i < len(filter); i += 2 {
			if i+1 < len(filter) {
				tx = tx.Where(filter[i]+" = ?", filter[i+1])
			}
		}
	}

	switch req.Sort {
	case utils.SORT_CREATED_AT_OLDEST:
		tx = tx.Order("created_at")
	default:
		tx = tx.Order("created_at desc")
	}

	if req.MovieId != "" {
		tx = tx.Group("room.id")
	}

	if err := tx.Find(&rs.Data).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	if rs.Meta, err = r.GetPaginationInfo("", tx, total.Count, page, pageSize); err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}
