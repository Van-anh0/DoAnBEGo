package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
	"strings"
)

func (r *RepoPG) CreateCinema(ctx context.Context, ob *model.Cinema) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateCinema(ctx context.Context, ob *model.Cinema) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Updates(&ob).Error
}

func (r *RepoPG) DeleteCinema(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.Cinema{}).Error
}

func (r *RepoPG) GetOneCinema(ctx context.Context, id string) (*model.Cinema, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.Cinema{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

func (r *RepoPG) GetListCinema(ctx context.Context, req model.CinemaParams) (*model.CinemaResponse, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.CinemaResponse{}
	var err error
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	total := new(struct {
		Count int `json:"count"`
	})

	tx = tx.Model(&model.Cinema{}).Select("movie_theater.*")

	if req.MovieId != "" || req.Day != "" {
		tx = tx.Joins("JOIN showtime ON showtime.movie_theater_id = movie_theater.id")
		if req.MovieId != "" {
			tx = tx.Where("showtime.movie_id = ?", req.MovieId)
		}

		if req.Day != "" {
			tx.Where("showtime.day = ?", req.Day)
		}
	}

	if req.Search != "" {
		tx = tx.Where("unaccent(name) ilike %?%", req.Search)
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

	if req.MovieId != "" || req.Day != "" {
		tx = tx.Group("movie_theater.id")
	}

	if err := tx.Find(&rs.Data).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	if rs.Meta, err = r.GetPaginationInfo("", tx, total.Count, page, pageSize); err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}
