package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
)

func (r *RepoPG) CreateRoom(ctx context.Context, ob *model.Room) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateRoom(ctx context.Context, ob *model.Room) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Save(ob).Error
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

	return nil, nil
}
