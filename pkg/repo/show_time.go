package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
)

func (r *RepoPG) CreateShowtime(ctx context.Context, ob *model.Showtime) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateShowtime(ctx context.Context, ob *model.Showtime) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Save(ob).Error
}

func (r *RepoPG) DeleteShowtime(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.Showtime{}).Error
}

func (r *RepoPG) GetOneShowtime(ctx context.Context, id string) (*model.Showtime, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.Showtime{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return nil, nil
}
