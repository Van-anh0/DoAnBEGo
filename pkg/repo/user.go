package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
	"github.com/praslar/cloud0/logger"
)

func (r *RepoPG) CreateUser(ctx context.Context, ob *model.User) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateUser(ctx context.Context, ob *model.User) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.WithContext(ctx).Where("id = ?", ob.ID).Updates(&ob).Error; err != nil {
		log.WithError(err)
		return err
	}
	return nil
}

func (r *RepoPG) DeleteUser(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.User{}).Error
}

func (r *RepoPG) GetOneUser(ctx context.Context, id string) (*model.User, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.User{}
	if err := tx.Where("id = ?", id).Take(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

func (r *RepoPG) GetListUser(ctx context.Context, req model.UserParams) (*model.UserResponse, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.UserResponse{}
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
