package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
)

func (r *RepoPG) CreateOrder(ctx context.Context, ob *model.Order) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateOrder(ctx context.Context, ob *model.Order) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Save(ob).Error
}

func (r *RepoPG) DeleteOrder(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.Order{}).Error
}

func (r *RepoPG) GetOneOrder(ctx context.Context, id string) (*model.Order, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.Order{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return nil, nil
}
