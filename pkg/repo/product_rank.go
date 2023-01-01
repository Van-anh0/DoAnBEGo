package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
	"strings"
)

func (r *RepoPG) CreateMovieRank(ctx context.Context, ob *model.MovieRank) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateMovieRank(ctx context.Context, ob *model.MovieRank) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Updates(&ob).Error
}

func (r *RepoPG) DeleteMovieRank(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.MovieRank{}).Error
}

func (r *RepoPG) GetOneMovieRank(ctx context.Context, id string) (*model.MovieRank, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.MovieRank{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

func (r *RepoPG) GetListMovieRank(ctx context.Context, req model.MovieRankParams) (*model.MovieRankResponse, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.MovieRankResponse{}
	var err error
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	total := new(struct {
		Count int `json:"count"`
	})

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
	if err := tx.Find(&rs.Data).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	if rs.Meta, err = r.GetPaginationInfo("", tx, total.Count, page, pageSize); err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}
