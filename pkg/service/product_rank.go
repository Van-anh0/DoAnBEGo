package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type ProductRankService struct {
	repo repo.PGInterface
}

type ProductRankInterface interface {
	Create(ctx context.Context, ob model.ProductRankRequest) (rs *model.ProductRank, err error)
	Update(ctx context.Context, ob model.ProductRankRequest) (rs *model.ProductRank, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.ProductRank, err error)
	GetList(ctx context.Context, req model.ProductRankParams) (rs *model.ProductRankResponse, err error)
}

func NewProductRankService(repo repo.PGInterface) ProductRankInterface {
	return &ProductRankService{repo: repo}
}

func (s *ProductRankService) Create(ctx context.Context, req model.ProductRankRequest) (rs *model.ProductRank, err error) {

	ob := &model.ProductRank{}
	common.Sync(req, ob)

	if err := s.repo.CreateProductRank(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ProductRankService) Update(ctx context.Context, req model.ProductRankRequest) (rs *model.ProductRank, err error) {
	ob, err := s.repo.GetOneProductRank(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateProductRank(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ProductRankService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteProductRank(ctx, id)
}

func (s *ProductRankService) GetOne(ctx context.Context, id string) (rs *model.ProductRank, err error) {

	ob, err := s.repo.GetOneProductRank(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ProductRankService) GetList(ctx context.Context, req model.ProductRankParams) (rs *model.ProductRankResponse, err error) {

	ob, err := s.repo.GetListProductRank(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
