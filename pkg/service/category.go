package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type CategoryService struct {
	repo repo.PGInterface
}

type CategoryInterface interface {
	Create(ctx context.Context, ob model.CategoryRequest) (rs *model.Category, err error)
	Update(ctx context.Context, ob model.CategoryRequest) (rs *model.Category, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Category, err error)
	GetList(ctx context.Context, req model.CategoryParams) (rs *model.CategoryResponse, err error)
}

func NewCategoryService(repo repo.PGInterface) CategoryInterface {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(ctx context.Context, req model.CategoryRequest) (rs *model.Category, err error) {

	ob := &model.Category{}
	common.Sync(req, ob)

	if err := s.repo.CreateCategory(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CategoryService) Update(ctx context.Context, req model.CategoryRequest) (rs *model.Category, err error) {
	ob, err := s.repo.GetOneCategory(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateCategory(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CategoryService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteCategory(ctx, id)
}

func (s *CategoryService) GetOne(ctx context.Context, id string) (rs *model.Category, err error) {

	ob, err := s.repo.GetOneCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CategoryService) GetList(ctx context.Context, req model.CategoryParams) (rs *model.CategoryResponse, err error) {

	ob, err := s.repo.GetListCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
