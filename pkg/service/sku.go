package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type SkuService struct {
	repo repo.PGInterface
}

type SkuInterface interface {
	Create(ctx context.Context, ob model.SkuRequest) (rs *model.Sku, err error)
	Update(ctx context.Context, ob model.SkuRequest) (rs *model.Sku, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Sku, err error)
	GetList(ctx context.Context, req model.SkuParams) (rs *model.SkuResponse, err error)
}

func NewSkuService(repo repo.PGInterface) SkuInterface {
	return &SkuService{repo: repo}
}

func (s *SkuService) Create(ctx context.Context, req model.SkuRequest) (rs *model.Sku, err error) {

	ob := &model.Sku{}
	common.Sync(req, ob)

	if err := s.repo.CreateSku(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *SkuService) Update(ctx context.Context, req model.SkuRequest) (rs *model.Sku, err error) {
	ob, err := s.repo.GetOneSku(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateSku(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *SkuService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteSku(ctx, id)
}

func (s *SkuService) GetOne(ctx context.Context, id string) (rs *model.Sku, err error) {

	ob, err := s.repo.GetOneSku(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *SkuService) GetList(ctx context.Context, req model.SkuParams) (rs *model.SkuResponse, err error) {

	ob, err := s.repo.GetListSku(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
