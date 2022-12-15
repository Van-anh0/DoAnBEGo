package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
)

func (r *RepoPG) CreateMovie(ctx context.Context, ob *model.Movie) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateMovie(ctx context.Context, ob *model.Movie) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Updates(&ob).Error
}

func (r *RepoPG) DeleteMovie(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.Movie{}).Error
}

func (r *RepoPG) GetOneMovie(ctx context.Context, id string) (*model.Movie, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.Movie{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

func (r *RepoPG) GetListMovie(ctx context.Context, req model.MovieParams) (*model.MovieResponse, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.MovieResponse{}
	var err error
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	total := new(struct {
		Count int `json:"count"`
	})

	if req.Search != "" {
		tx = tx.Where("unaccent(name) ilike %?%", req.Search)
	}

	if len(req.Filter) > 0 {
		for i := 0; i < len(req.Filter); i++ {
			tx = tx.Where("? = ?", req.Filter[i].Key, req.Filter[i].Value)
		}
	}

	switch req.Sort {
	case utils.SORT_CREATED_AT_OLDEST:
		tx = tx.Order("created_at")
	default:
		tx = tx.Order("created_at desc")
	}
	if err := tx.Find(&rs.Data).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	if rs.Meta, err = r.GetPaginationInfo("", tx, total.Count, page, pageSize); err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}
