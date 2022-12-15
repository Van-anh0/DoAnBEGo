package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
)

func (r *RepoPG) CreateMetadata(ctx context.Context, ob *model.Metadata) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateMetadata(ctx context.Context, ob *model.Metadata) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Save(ob).Error
}

func (r *RepoPG) DeleteMetadata(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.Metadata{}).Error
}

func (r *RepoPG) GetOneMetadata(ctx context.Context, id string) (*model.Metadata, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.Metadata{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return nil, nil
}
