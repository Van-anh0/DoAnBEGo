package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type PromotionService struct {
	repo repo.PGInterface
}

type PromotionInterface interface {
	Create(ctx context.Context, ob model.PromotionRequest) (rs *model.Promotion, err error)
	Update(ctx context.Context, ob model.PromotionRequest) (rs *model.Promotion, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Promotion, err error)
	GetList(ctx context.Context, req model.PromotionParams) (rs *model.PromotionResponse, err error)
}

func NewPromotionService(repo repo.PGInterface) PromotionInterface {
	return &PromotionService{repo: repo}
}

func (s *PromotionService) Create(ctx context.Context, req model.PromotionRequest) (rs *model.Promotion, err error) {

	ob := &model.Promotion{}
	common.Sync(req, ob)

	if err := s.repo.CreatePromotion(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *PromotionService) Update(ctx context.Context, req model.PromotionRequest) (rs *model.Promotion, err error) {
	ob, err := s.repo.GetOnePromotion(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdatePromotion(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *PromotionService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeletePromotion(ctx, id)
}

func (s *PromotionService) GetOne(ctx context.Context, id string) (rs *model.Promotion, err error) {

	ob, err := s.repo.GetOnePromotion(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *PromotionService) GetList(ctx context.Context, req model.PromotionParams) (rs *model.PromotionResponse, err error) {

	ob, err := s.repo.GetListPromotion(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
