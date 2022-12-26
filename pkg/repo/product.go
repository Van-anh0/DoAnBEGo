package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
	"gorm.io/gorm"
	"strings"
)

func (r *RepoPG) CreateProduct(ctx context.Context, ob *model.Product) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateProduct(ctx context.Context, ob *model.Product) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Updates(&ob).Error
}

func (r *RepoPG) DeleteProduct(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.Product{}).Error
}

func (r *RepoPG) GetOneProduct(ctx context.Context, id string) (*model.Product, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.Product{}
	if err := tx.Where("id = ?", id).Preload("Sku", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Showtime")
	}).Take(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

func (r *RepoPG) GetListProduct(ctx context.Context, req model.ProductParams) (*model.ListProductResponse, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.ListProductResponse{}
	var err error
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	total := new(struct {
		Count int `json:"count"`
	})

	tx = tx.Select("product.*")

	if req.Day != "" || req.MovieTheaterId != "" {
		tx = tx.Joins("Join sku s on s.product_id = product.id").Joins("Join showtime st on st.sku_id = s.id")
		if req.Day != "" {
			tx = tx.Where("st.day = ?", req.Day)
		}

		if req.MovieTheaterId != "" {
			tx = tx.Where("st.movie_theater_id = ?", req.MovieTheaterId)
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

	if req.Day != "" || req.MovieTheaterId != "" {
		tx = tx.Group("product.id")
	}

	// find product and preload sku
	if err := tx.Preload("Sku", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Showtime")
	}).Find(&rs.Data).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	//2022-12-18T00:00:00+07:00
	//postman: 2022-12-18T00:00:00 07:00
	if rs.Meta, err = r.GetPaginationInfo("", tx, total.Count, page, pageSize); err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}
