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
	return tx.Where("id = ?", ob.ID).Save(ob).Error
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

	return nil, nil
}
