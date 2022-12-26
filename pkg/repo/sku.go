package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
	"strings"
)

func (r *RepoPG) CreateSku(ctx context.Context, ob *model.Sku) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateSku(ctx context.Context, ob *model.Sku) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Updates(&ob).Error
}

func (r *RepoPG) DeleteSku(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.Sku{}).Error
}

func (r *RepoPG) GetOneSku(ctx context.Context, id string) (*model.Sku, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.Sku{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

func (r *RepoPG) GetListSku(ctx context.Context, req model.SkuParams) (*model.SkuResponse, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.SkuResponse{}
	var err error
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	total := new(struct {
		Count int `json:"count"`
	})

	// filter by day or movie_theater_id
	if req.Day != "" || req.MovieTheaterId != "" {
		// join table showtime
		tx = tx.Select("sku.*").Joins("left join showtime on showtime.sku_id = sku.id")
		if req.Day != "" {
			tx = tx.Where("showtime.day = ?", req.Day)
		}
		if req.MovieTheaterId != "" {
			tx = tx.Where("showtime.movie_theater_id = ?", req.MovieTheaterId)
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

	// if req.Day or req.MovieTheaterId then group by sku.id
	if req.Day != "" || req.MovieTheaterId != "" {
		tx = tx.Group("sku.id")
	}

	if err := tx.Preload("Showtime").Find(&rs.Data).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	if rs.Meta, err = r.GetPaginationInfo("", tx, total.Count, page, pageSize); err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}
