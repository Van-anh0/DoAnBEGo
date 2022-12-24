package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type CategoryHasProductService struct {
	repo repo.PGInterface
}

type CategoryHasProductInterface interface {
	Create(ctx context.Context, ob model.CategoryHasProductRequest) (rs *model.CategoryHasProduct, err error)
	Update(ctx context.Context, ob model.CategoryHasProductRequest) (rs *model.CategoryHasProduct, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.CategoryHasProduct, err error)
	GetList(ctx context.Context, req model.CategoryHasProductParams) (rs *model.CategoryHasProductResponse, err error)
}

func NewCategoryHasProductService(repo repo.PGInterface) CategoryHasProductInterface {
	return &CategoryHasProductService{repo: repo}
}

func (s *CategoryHasProductService) Create(ctx context.Context, req model.CategoryHasProductRequest) (rs *model.CategoryHasProduct, err error) {

	ob := &model.CategoryHasProduct{}
	common.Sync(req, ob)

	if err := s.repo.CreateCategoryHasProduct(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CategoryHasProductService) Update(ctx context.Context, req model.CategoryHasProductRequest) (rs *model.CategoryHasProduct, err error) {
	ob, err := s.repo.GetOneCategoryHasProduct(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateCategoryHasProduct(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CategoryHasProductService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteCategoryHasProduct(ctx, id)
}

func (s *CategoryHasProductService) GetOne(ctx context.Context, id string) (rs *model.CategoryHasProduct, err error) {

	ob, err := s.repo.GetOneCategoryHasProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CategoryHasProductService) GetList(ctx context.Context, req model.CategoryHasProductParams) (rs *model.CategoryHasProductResponse, err error) {

	ob, err := s.repo.GetListCategoryHasProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
